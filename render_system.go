package main

import "github.com/hajimehoshi/ebiten/v2"

func ProcessRenderables(g *Game, level Level, screen *ebiten.Image) {
	for _, res := range g.World.Query(g.WorldTags["renderables"]) {
		pos := res.Components[position].(*Position)
		img := res.Components[rendarable].(*Renderable).Image

		index := level.GetIndexFromXY(pos.X, pos.Y)
		tile := level.Tiles[index]
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(float64(tile.PixelX), float64(tile.PixelY))
		screen.DrawImage(img, op)
	}
}
