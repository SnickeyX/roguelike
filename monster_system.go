package main

import (
	"log"

	"github.com/SnickeyX/roguelike/state"
	"github.com/SnickeyX/roguelike/utils"
	"github.com/SnickeyX/roguelike/world"
)

const MONSTER_FOV = 8

func TakeMonsterAction(g *Game) {
	l := g.Map.CurrentLevel

	players := g.WorldTags["players"]
	// NOTE: assuming single-player for the moment
	player := g.World.Query(players)[0]
	pos := player.Components[world.Position].(*utils.Position)
	pX, pY := pos.X, pos.Y

	monsters := g.WorldTags["monsters"]
	for _, res := range g.World.Query(monsters) {
		pos := res.Components[world.Position].(*utils.Position)
		ren := res.Components[world.Rendarable].(*utils.Renderable)
		name := res.Components[world.Name].(*utils.Name)
		aoe := res.Components[world.MeleeWeapon].(*utils.MeleeWeapon).Aoe
		// dist between monster and player
		d_e := utils.EuclidianDist(pX, pY, pos.X, pos.Y)
		d_c := utils.ChebyshevDist(pX, pY, pos.X, pos.Y)

		// attack player if they are within monster's aoe
		if d_c <= float64(aoe) {
			AttackSystem(g, res, player)
		} else if d_e < MONSTER_FOV {
			ren.Image = utils.SkeleBuffImg
			// get random path without any care for blocked tiles
			indexes := l.GetPath(pX, pY, pos.X, pos.Y, false, true)
			if indexes == nil || len(indexes) < 2 {
				log.Printf("%s cant move to player", name.Label)
			} else {
				next_tile := l.Tiles[indexes[len(indexes)-2]]
				if !next_tile.Blocked {
					l.Tiles[l.GetIndexFromXY(pos.X, pos.Y)].Blocked = false
					pos.X = next_tile.PixelX / utils.GameConstants.TileWidth
					pos.Y = next_tile.PixelY / utils.GameConstants.TileHeight
					next_tile.Blocked = true
				}
			}
		} else {
			ren.Image = utils.SkeleIdleImg
		}

	}
	g.Turn = state.GetNextState(g.Turn)
	g.TurnCounter = 0
}
