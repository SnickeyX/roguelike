package main

import (
	"github.com/SnickeyX/roguelike/state"
	"github.com/SnickeyX/roguelike/utils"
	"github.com/SnickeyX/roguelike/world"
	"github.com/hajimehoshi/ebiten/v2"
)

func TryMovePlayer(g *Game) {
	turnTaken := false
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
	// skip turn key
	if ebiten.IsKeyPressed(ebiten.KeyQ) {
		turnTaken = true
	}

	level := g.Map.CurrentLevel

	// probably should not move all players but oh well
	for _, result := range g.World.Query(players) {
		pos := result.Components[world.Position].(*utils.Position)
		new_x := (pos.X + x) % utils.GameConstants.ScreenWidth
		new_y := (pos.Y + y)

		// out-of-bounds prot
		if new_y <= 0 {
			new_y = utils.GameConstants.ScreenHeight - 1
		}
		if new_y > utils.GameConstants.ScreenHeight-1 {
			new_y = 1
		}

		index := g.Map.CurrentLevel.GetIndexFromXY(new_x, new_y)
		tile := level.Tiles[index]

		if !tile.Blocked {
			level.Tiles[level.GetIndexFromXY(pos.X, pos.Y)].Blocked = false

			level.PlayerLoc[0], level.PlayerLoc[1] = new_x, new_y
			pos.X = new_x
			pos.Y = new_y
			level.Tiles[index].Blocked = true
		}
	}

	if x != 0 || y != 0 || turnTaken {
		g.Turn = state.GetNextState(g.Turn)
		g.TurnCounter = 0
	}
}
