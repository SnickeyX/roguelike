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
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var LogMessage string
var logTics int = 0

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
	utils.LoadAllAssets()
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
		TakePlayerAction(g)
	}
	if g.Turn == state.MonsterTurn && g.TurnCounter > 10 {
		TakeMonsterAction(g)
	}
	return nil
}

// to draw every draw cycle
func (g *Game) Draw(screen *ebiten.Image) {
	lvl := g.Map.CurrentLevel
	lvl.DrawLevel(screen)
	ProcessRenderables(g, lvl, screen)
	logTics++
	ebitenutil.DebugPrintAt(screen, LogMessage,
		(utils.GameConstants.ScreenWidth*utils.GameConstants.TileWidth)/3,
		(utils.GameConstants.ScreenHeight*utils.GameConstants.TileHeight)-
			((utils.GameConstants.UiHeight*utils.GameConstants.TileHeight)/2))
	if logTics > 300 {
		logTics = 0
		// clear
		LogMessage = ""
	}
}

func (g *Game) Layout(w, h int) (int, int) {
	width := utils.GameConstants.ScreenWidth * utils.GameConstants.TileWidth
	height := utils.GameConstants.ScreenHeight * utils.GameConstants.TileHeight
	return width, height
}

func main() {
	g := NewGame()
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	ebiten.SetWindowTitle("Dungeon Crawler^(tm)")
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
