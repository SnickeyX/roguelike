package main

import (
	"log"

	"github.com/SnickeyX/roguelike/state"
	"github.com/SnickeyX/roguelike/utils"
	"github.com/SnickeyX/roguelike/world"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const MONSTER_FOV = 4

func UpdateMonster(g *Game) {
	gd := utils.NewGameData()
	l := g.Map.CurrentLevel
	// NOTE: won't work for multiple players, need to query players component to handle that
	// for single-player atm, this works
	pX, pY := l.PlayerLoc[0], l.PlayerLoc[1]

	monsters := g.WorldTags["monsters"]

	// pre-loading images: make this better!
	buffSkellyImg, _, err := ebitenutil.NewImageFromFile("assets/skele_buff_1.png")
	if err != nil {
		log.Fatal(err)
	}
	skellyImg, _, err := ebitenutil.NewImageFromFile("assets/skele_idle_1.png")
	if err != nil {
		log.Fatal(err)
	}

	for _, res := range g.World.Query(monsters) {
		pos := res.Components[world.Position].(*utils.Position)
		ren := res.Components[world.Rendarable].(*utils.Renderable)
		mon := res.Components[world.Monster].(*utils.Monster)
		// dist between monster and player
		d := utils.EuclidianDist(pX, pY, pos.X, pos.Y)

		if d < MONSTER_FOV {
			ren.Image = buffSkellyImg
			// get random path without any care for blocked tiles
			indexes := l.GetPath(pX, pY, pos.X, pos.Y, false, true)
			if indexes == nil || len(indexes) < 2 {
				log.Printf("%s cant move to player", mon.Name)
			} else {
				next_tile := l.Tiles[indexes[len(indexes)-2]]
				if !next_tile.Blocked {
					l.Tiles[l.GetIndexFromXY(pos.X, pos.Y)].Blocked = false
					pos.X = next_tile.PixelX / gd.TileWidth
					pos.Y = next_tile.PixelY / gd.TileHeight
					next_tile.Blocked = true
				}
			}
		} else {
			ren.Image = skellyImg
		}

	}
	g.Turn = state.PlayerTurn
}
