package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

// Level holds the tile information for a complete dungeon level.
type Level struct {
	// Tiles are ordered row-by-row, left-to-right
	Tiles []MapTile
	Rooms []Rect
}

// NewLevel creates a new game level in a dungeon.
func NewLevel() Level {
	l := Level{}
	rooms := make([]Rect, 0)
	l.Rooms = rooms
	l.GenerateLevelTiles()
	return l
}

type MapTile struct {
	PixelX  int
	PixelY  int
	Blocked bool
	Image   *ebiten.Image
}

// GetIndexFromXY gets the index of the map array from a given X,Y TILE coordinate.
func (level *Level) GetIndexFromXY(x int, y int) int {
	gd := NewGameData()
	return (y * gd.ScreenWidth) + x
}

// everything is a wall initially
func (level *Level) CreateTiles() []MapTile {
	gd := NewGameData()
	tiles := make([]MapTile, gd.ScreenHeight*gd.ScreenWidth)
	for x := 0; x < gd.ScreenWidth; x++ {
		for y := 0; y < gd.ScreenHeight; y++ {
			index := level.GetIndexFromXY(x, y)
			wall, _, err := ebitenutil.NewImageFromFile("assets/wall.png")
			if err != nil {
				log.Fatal(err)
			}
			tile := MapTile{
				PixelX:  x * gd.TileWidth,
				PixelY:  y * gd.TileHeight,
				Blocked: true,
				Image:   wall,
			}
			tiles[index] = tile
		}
	}
	return tiles
}

// setting map for non-blocked rooms within rectangular rooms
func (level *Level) createRoom(room Rect) {
	for y := room.Y1 + 1; y < room.Y2; y++ {
		for x := room.X1 + 1; x < room.X2; x++ {
			index := level.GetIndexFromXY(x, y)
			level.Tiles[index].Blocked = false
			floor, _, err := ebitenutil.NewImageFromFile("assets/floor.png")
			if err != nil {
				log.Fatal(err)
			}
			level.Tiles[index].Image = floor

		}
	}
}

// creating rooms
func (level *Level) GenerateLevelTiles() {
	MIN_RECT_SIZE := 6
	MAX_RECT_SIZE := 10
	MAX_ROOMS := 30
	contains_rooms := false

	gd := NewGameData()
	tiles := level.CreateTiles()
	level.Tiles = tiles

	for i := 0; i < MAX_ROOMS; i++ {
		// randomly generating a room between min and max size
		w := GetRandomBetweenTwo(MIN_RECT_SIZE, MAX_RECT_SIZE)
		h := GetRandomBetweenTwo(MIN_RECT_SIZE, MAX_RECT_SIZE)
		// choosing a starting top left of the room
		x := GetDiceRoll(gd.ScreenWidth - w - 1)
		y := GetDiceRoll(gd.ScreenHeight - h - 1)

		new_room := NewRect(x, y, w, h)
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
func (level *Level) CreatePathForRoom(new_room Rect) {
	newX, newY := new_room.Center()
	prevX, prevY := level.Rooms[GetDiceRoll(len(level.Rooms))-1].Center()

	flipRes1 := GetDiceRoll(2)

	if flipRes1 == 1 {
		indexes := level.GetShortestPath(newX, newY, prevX, prevY)
		level.CreateTunnelFromIndexes(indexes)
	} else {
		flipRes2 := GetDiceRoll(2)
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
	gd := NewGameData()
	for i := min(x1, x2); i < max(x1, x2)+1; i++ {
		index := level.GetIndexFromXY(i, y)
		if index > 0 && index < gd.ScreenWidth*gd.ScreenHeight {
			level.Tiles[index].Blocked = false
			floor, _, err := ebitenutil.NewImageFromFile("assets/floor.png")
			if err != nil {
				log.Fatal(err)
			}
			level.Tiles[index].Image = floor
		}
	}
}

// tunnel between two pixel points (x,y1) and (x,y2)
func (level *Level) CreateVerticalTunnel(y1 int, y2 int, x int) {
	gd := NewGameData()
	for i := min(y1, y2); i < max(y1, y2)+1; i++ {
		index := level.GetIndexFromXY(x, i)
		if index > 0 && index < gd.ScreenHeight*gd.ScreenWidth {
			level.Tiles[index].Blocked = false
			floor, _, err := ebitenutil.NewImageFromFile("assets/floor.png")
			if err != nil {
				log.Fatal(err)
			}
			level.Tiles[index].Image = floor
		}
	}
}

func (level *Level) CreateTunnelFromIndexes(indexes []int) {
	gd := NewGameData()
	for _, index := range indexes {
		if index > 0 && index < gd.ScreenWidth*gd.ScreenHeight {
			level.Tiles[index].Blocked = false
			floor, _, err := ebitenutil.NewImageFromFile("assets/floor.png")
			if err != nil {
				log.Fatal(err)
			}
			level.Tiles[index].Image = floor
		}
	}
}
