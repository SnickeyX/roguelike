package main

import (
	"github.com/SnickeyX/roguelike/state"
	"github.com/SnickeyX/roguelike/utils"
	"github.com/SnickeyX/roguelike/world"
	"github.com/hajimehoshi/ebiten/v2"
)

func TryMovePlayer(g *Game) {
	gd := utils.NewGameData()
	players := g.WorldTags["players"]
	x := 0
	y := 0
	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		y = -1
	}
	if ebiten.IsKeyPressed(ebiten.KeyDown) {
		y = 1
	}
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		x = -1
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		x = 1
	}

	level := g.Map.CurrentLevel

	// probably should not move all players but oh well
	for _, result := range g.World.Query(players) {
		pos := result.Components[world.Position].(*utils.Position)
		new_x := (pos.X + x) % gd.ScreenWidth
		new_y := (pos.Y + y)

		// out-of-bounds prot
		if new_y <= 0 {
			new_y = gd.ScreenHeight - 1
		}
		if new_y > gd.ScreenHeight-1 {
			new_y = 1
		}

		index := g.Map.CurrentLevel.GetIndexFromXY(new_x, new_y)
		tile := level.Tiles[index]

		if !tile.Blocked {
			level.PlayerLoc[0], level.PlayerLoc[1] = new_x, new_y
			pos.X = new_x
			pos.Y = new_y
		}
	}

	if x != 0 || y != 0 {
		g.Turn = state.GetNextState(g.Turn)
		g.TurnCounter = 0
	}
}
