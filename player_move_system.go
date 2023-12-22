package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

func TryMovePlayer(g *Game) {
	gd := NewGameData()
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
		pos := result.Components[position].(*Position)
		new_x := (pos.X + x) % gd.ScreenWidth
		new_y := (pos.Y + y)
		// silly out-of-bounds prot
		if new_y <= 0 {
			new_y = gd.ScreenHeight - 1
		}
		if new_y > gd.ScreenHeight-1 {
			new_y = 1
		}

		index := g.Map.CurrentLevel.GetIndexFromXY(new_x, new_y)
		tile := level.Tiles[index]

		if !tile.Blocked {
			pos.X = new_x
			pos.Y = new_y
		}
	}
}
