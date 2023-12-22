package main

// Following ECS as the main design pattern
// Code possible from tutorials by https://github.com/RAshkettle (FatOldYeti) & RogueBasin.com

import (
	"log"

	"github.com/bytearena/ecs"

	"github.com/hajimehoshi/ebiten/v2"
)

// Game struct to hold all global data
type Game struct {
	Map         GameMap
	World       *ecs.Manager
	WorldTags   map[string]ecs.Tag
	Turn        TurnState
	TurnCounter int
}

func NewGame() *Game {
	g := &Game{}
	g.Map = NewGameMap()
	world, tags := InitializeWorld(g.Map.CurrentLevel)
	g.World = world
	g.WorldTags = tags
	g.Turn = PlayerTurn
	g.TurnCounter = 0
	return g
}

// update frame at each tic, 60hz by defualt
func (g *Game) Update() error {
	g.TurnCounter++
	if g.Turn == PlayerTurn && g.TurnCounter > 10 {
		TryMovePlayer(g)
	}
	g.Turn = PlayerTurn
	return nil
}

func (lvl *Level) DrawLevel(screen *ebiten.Image) {
	gd := NewGameData()
	for x := 0; x < gd.ScreenWidth; x++ {
		for y := 0; y < gd.ScreenHeight; y++ {
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

// to draw every draw cycle
func (g *Game) Draw(screen *ebiten.Image) {
	lvl := g.Map.CurrentLevel
	lvl.DrawLevel(screen)
	ProcessRenderables(g, lvl, screen)
}

func (g *Game) Layout(w, h int) (int, int) {
	gd := NewGameData()
	return gd.ScreenWidth * gd.TileWidth, gd.ScreenHeight * gd.TileHeight
}

func main() {
	g := NewGame()
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	ebiten.SetWindowTitle("Dungeon Crawler^(tm)")
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
