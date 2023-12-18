package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

// Game struct to hold all global data
type Game struct {}

func NewGame() *Game {
	g := &Game{}
	return g 
}

// update frame at each tic, 60hz by defualt
func (g *Game) Update() error {
	return nil 
}

// to draw every draw cycle
func (g *Game) Draw(screen *ebiten.Image) {	
}

func (g *Game) Layout(w, h int) (int,int) {return 1280, 800}

func main() {
	g := NewGame()
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	ebiten.SetWindowTitle("Dungeon Crawler^(tm)")
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}