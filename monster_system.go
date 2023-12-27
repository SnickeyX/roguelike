package main

import (
	"log"

	"github.com/SnickeyX/roguelike/state"
	"github.com/SnickeyX/roguelike/utils"
	"github.com/SnickeyX/roguelike/world"
)

const MONSTER_FOV = 5

func UpdateMonster(g *Game) {
	l := g.Map.CurrentLevel
	// NOTE: won't work for multiple players, need to query players component to handle that
	pX, pY := l.PlayerLoc[0], l.PlayerLoc[1]

	monsters := g.WorldTags["monsters"]

	for _, res := range g.World.Query(monsters) {
		pos := res.Components[world.Position].(*utils.Position)
		mon := res.Components[world.Monster].(*utils.Monster)
		// dist between monster and player
		d := utils.EuclidianDist(pX, pY, pos.X, pos.Y)

		if d < MONSTER_FOV {
			log.Printf("%s shivers to its bones.", mon.Name)
		}
	}
	g.Turn = state.PlayerTurn
}
