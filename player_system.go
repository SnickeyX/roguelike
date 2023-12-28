package main

import (
	"github.com/SnickeyX/roguelike/state"
	"github.com/SnickeyX/roguelike/utils"
	"github.com/SnickeyX/roguelike/world"
	"github.com/hajimehoshi/ebiten/v2"
)

func TakePlayerAction(g *Game) {
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

	for _, playerQ := range g.World.Query(players) {
		pos := playerQ.Components[world.Position].(*utils.Position)
		new_x := (pos.X + x) % utils.GameConstants.ScreenWidth
		new_y := (pos.Y + y)

		index := g.Map.CurrentLevel.GetIndexFromXY(new_x, new_y)
		tile := level.Tiles[index]

		if !tile.Blocked {
			level.Tiles[level.GetIndexFromXY(pos.X, pos.Y)].Blocked = false

			level.PlayerLoc[0], level.PlayerLoc[1] = new_x, new_y
			pos.X = new_x
			pos.Y = new_y
			level.Tiles[index].Blocked = true
		} else {
			// player attacks all monsters wihtin its AOE, up to a maximum of num_hits
			if tile.TileType != world.WALL {
				p_weapon := playerQ.Components[world.MeleeWeapon].(*utils.MeleeWeapon)
				aoe := p_weapon.Aoe
				num_hits := p_weapon.NumHits
				curr_hits := 0
				monsters := g.WorldTags["monsters"]
				for _, monQ := range g.World.Query(monsters) {
					if curr_hits == num_hits {
						break
					}
					mon_p := monQ.Components[world.Position].(*utils.Position)
					d := utils.ChebyshevDist(pos.X, pos.Y, mon_p.X, mon_p.Y)
					if d <= float64(aoe) {
						AttackSystem(g, playerQ, monQ)
						curr_hits++
					}
				}
			}
		}
	}

	if x != 0 || y != 0 || turnTaken {
		g.Turn = state.GetNextState(g.Turn)
		g.TurnCounter = 0
	}
}
