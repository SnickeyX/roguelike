package world

import (
	"log"

	"github.com/SnickeyX/roguelike/utils"

	"github.com/hajimehoshi/ebiten/v2"
)

type TileType int

const (
	WALL TileType = iota
	FLOOR
)

// Level holds the tile information for a complete dungeon level.
type Level struct {
	// Tiles are ordered row-by-row, left-to-right
	Tiles     []*MapTile
	Rooms     []utils.Rect
	PlayerLoc []int
	FovDist   float64 // radius of player FOV
}

// NewLevel creates a new game level in a dungeon.
func NewLevel() Level {
	l := Level{}
	l.FovDist = 6
	l.PlayerLoc = make([]int, 2)
	l.Rooms = make([]utils.Rect, 0)
	l.GenerateLevelTiles()
	return l
}

type MapTile struct {
	PixelX     int
	PixelY     int
	Blocked    bool
	IsRevealed bool
	Image      *ebiten.Image
	TileType   TileType
}

// GetIndexFromXY gets the index of the map array from a given X,Y TILE coordinate.
func (level *Level) GetIndexFromXY(x, y int) int {
	return (y * utils.GameConstants.ScreenWidth) + x
}

func (lvl *Level) DrawLevel(screen *ebiten.Image) {
	for x := 0; x < utils.GameConstants.ScreenWidth; x++ {
		for y := 0; y < utils.GameConstants.ScreenHeight; y++ {
			index := lvl.GetIndexFromXY(x, y)
			isViz := lvl.IsVizToPlayer(x, y)
			tile := lvl.Tiles[index]
			if isViz {
				op := &ebiten.DrawImageOptions{}
				op.GeoM.Translate(float64(tile.PixelX), float64(tile.PixelY))
				screen.DrawImage(tile.Image, op)
				lvl.Tiles[index].IsRevealed = true

			} else if tile.IsRevealed {
				op := &ebiten.DrawImageOptions{}
				op.GeoM.Translate(float64(tile.PixelX), float64(tile.PixelY))
				op.ColorScale.ScaleAlpha(0.5)
				screen.DrawImage(tile.Image, op)
			}

		}
	}
}

func (level *Level) IsVizToPlayer(x, y int) bool {
	px, py := level.PlayerLoc[0], level.PlayerLoc[1]
	d := utils.EuclidianDist(x, y, px, py)
	return d < level.FovDist
}

// everything is a wall initially
func (level *Level) CreateTiles() []*MapTile {
	tiles := make([]*MapTile, utils.GameConstants.ScreenHeight*utils.GameConstants.ScreenWidth)
	for x := 0; x < utils.GameConstants.ScreenWidth; x++ {
		for y := 0; y < utils.GameConstants.ScreenHeight; y++ {
			index := level.GetIndexFromXY(x, y)
			tile := MapTile{
				PixelX:     x * utils.GameConstants.TileWidth,
				PixelY:     y * utils.GameConstants.TileHeight,
				Blocked:    true,
				IsRevealed: false,
				Image:      utils.WallImg,
				TileType:   WALL,
			}
			tiles[index] = &tile
		}
	}
	return tiles
}

// setting map for non-blocked rooms within rectangular rooms
func (level *Level) createRoom(room utils.Rect) {
	for y := room.Y1 + 1; y < room.Y2; y++ {
		for x := room.X1 + 1; x < room.X2; x++ {
			index := level.GetIndexFromXY(x, y)
			level.Tiles[index].Blocked = false
			level.Tiles[index].TileType = FLOOR
			level.Tiles[index].Image = utils.FloorImg

		}
	}
}

// creating rooms
func (level *Level) GenerateLevelTiles() {
	MIN_RECT_SIZE := 6
	MAX_RECT_SIZE := 10
	MAX_ROOMS := 30
	contains_rooms := false

	tiles := level.CreateTiles()
	level.Tiles = tiles

	for i := 0; i < MAX_ROOMS; i++ {
		// randomly generating a room between min and max size
		w := utils.GetRandomBetweenTwo(MIN_RECT_SIZE, MAX_RECT_SIZE)
		h := utils.GetRandomBetweenTwo(MIN_RECT_SIZE, MAX_RECT_SIZE)
		// choosing a starting top left of the room
		x := utils.GetDiceRoll(utils.GameConstants.ScreenWidth - w - 1)
		y := utils.GetDiceRoll(utils.GameConstants.ScreenHeight - h - 1)

		new_room := utils.NewRect(x, y, w, h)
		canAdd := true
		// ensuring new_room does not intersect with any existing rooms
		for _, otherR := range level.Rooms {
			if new_room.Intersect(otherR) {
				canAdd = false
				break
			}
		}
		// add new room if it does not intersect with an existing room
		if canAdd {
			level.createRoom(new_room)
			if contains_rooms {
				level.CreatePathForRoom(new_room)
			}
			level.Rooms = append(level.Rooms, new_room)
			contains_rooms = true
		}
	}
}

// Create path between rooms by either:
//  1. Running Dijkstra between two points from a graph with randomized weight (with Pr = 50%)
//  2. Carve straight tunnels that are either:
//     a) aligned horizontally between rooms or b) aligned vertically between rooms (with Pr = 25% each)
func (level *Level) CreatePathForRoom(new_room utils.Rect) {
	newX, newY := new_room.Center()
	prevX, prevY := level.Rooms[utils.GetDiceRoll(len(level.Rooms))-1].Center()

	flipRes1 := utils.GetDiceRoll(2)

	if flipRes1 == 1 {
		// get random path without any care for blocked tiles
		indexes := level.GetPath(newX, newY, prevX, prevY, true, false)
		if indexes == nil {
			log.Fatal("Unable to carve level!")
		}
		level.CreateTunnelFromIndexes(indexes)
	} else {
		flipRes2 := utils.GetDiceRoll(2)
		if flipRes2 == 1 {
			level.CreateHorizontalTunnel(newX, prevX, newY)
			level.CreateVerticalTunnel(newY, prevY, prevX)
		} else {
			level.CreateHorizontalTunnel(newX, prevX, prevY)
			level.CreateVerticalTunnel(newY, prevY, newX)
		}
	}
}

// tunnel between two pixel points (x1,y) and (x2,y)
func (level *Level) CreateHorizontalTunnel(x1 int, x2 int, y int) {
	for i := min(x1, x2); i < max(x1, x2)+1; i++ {
		index := level.GetIndexFromXY(i, y)
		if index > 0 && index < utils.GameConstants.ScreenWidth*utils.GameConstants.ScreenHeight {
			level.Tiles[index].Blocked = false
			level.Tiles[index].TileType = FLOOR
			level.Tiles[index].Image = utils.FloorImg
		}
	}
}

// tunnel between two pixel points (x,y1) and (x,y2)
func (level *Level) CreateVerticalTunnel(y1 int, y2 int, x int) {
	for i := min(y1, y2); i < max(y1, y2)+1; i++ {
		index := level.GetIndexFromXY(x, i)
		if index > 0 && index < utils.GameConstants.ScreenHeight*utils.GameConstants.ScreenWidth {
			level.Tiles[index].Blocked = false
			level.Tiles[index].TileType = FLOOR
			level.Tiles[index].Image = utils.FloorImg
		}
	}
}

func (level *Level) CreateTunnelFromIndexes(indexes []int) {
	for _, index := range indexes {
		if index > 0 && index < utils.GameConstants.ScreenWidth*utils.GameConstants.ScreenHeight {
			level.Tiles[index].Blocked = false
			level.Tiles[index].TileType = FLOOR
			level.Tiles[index].Image = utils.FloorImg
		}
	}
}
