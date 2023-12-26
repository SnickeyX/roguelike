package main

import (
	"github.com/SnickeyX/roguelike/utils"
	"github.com/SnickeyX/roguelike/world"
	"github.com/hajimehoshi/ebiten/v2"
)

func ProcessRenderables(g *Game, level world.Level, screen *ebiten.Image) {
	for _, res := range g.World.Query(g.WorldTags["renderables"]) {
		pos := res.Components[world.Position].(*utils.Position)
		img := res.Components[world.Rendarable].(*utils.Renderable).Image

		if level.IsVizToPlayer(pos.X, pos.Y) {
			index := level.GetIndexFromXY(pos.X, pos.Y)
			tile := level.Tiles[index]
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64(tile.PixelX), float64(tile.PixelY))
			screen.DrawImage(img, op)
		}
	}
}
