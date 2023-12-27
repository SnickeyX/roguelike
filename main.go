package main

// Following ECS as the main design pattern
// Code possible from tutorials by https://github.com/RAshkettle (FatOldYeti) & RogueBasin.com

import (
	"log"

	"github.com/SnickeyX/roguelike/state"
	"github.com/SnickeyX/roguelike/utils"
	"github.com/SnickeyX/roguelike/world"
	"github.com/bytearena/ecs"
	"github.com/hajimehoshi/ebiten/v2"
)

// Game struct to hold all global data
type Game struct {
	Map         world.GameMap
	World       *ecs.Manager
	WorldTags   map[string]ecs.Tag
	Turn        state.TurnState
	TurnCounter int
}

func NewGame() *Game {
	g := &Game{}
	g.Map = world.NewGameMap()
	world, tags := world.InitializeWorld(g.Map.CurrentLevel)
	g.World = world
	g.WorldTags = tags
	g.Turn = state.PlayerTurn
	g.TurnCounter = 0
	return g
}

// update frame at each tic, 60hz by defualt
func (g *Game) Update() error {
	g.TurnCounter++
	if g.Turn == state.PlayerTurn && g.TurnCounter > 10 {
		TryMovePlayer(g)
	}
	if g.Turn == state.MonsterTurn {
		UpdateMonster(g)
	}
	return nil
}

// to draw every draw cycle
func (g *Game) Draw(screen *ebiten.Image) {
	lvl := g.Map.CurrentLevel
	lvl.DrawLevel(screen)
	ProcessRenderables(g, lvl, screen)
}

func (g *Game) Layout(w, h int) (int, int) {
	gd := utils.NewGameData()
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
