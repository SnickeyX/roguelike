module github.com/SnickeyX/roguelike

go 1.21.5

require (
	github.com/SnickeyX/roguelike/state v0.0.0-00010101000000-000000000000
	github.com/SnickeyX/roguelike/utils v0.0.0-00010101000000-000000000000
	github.com/SnickeyX/roguelike/world v0.0.0-00010101000000-000000000000
	github.com/bytearena/ecs v1.0.0
	github.com/hajimehoshi/ebiten/v2 v2.6.3
)

require (
	github.com/dominikbraun/graph v0.23.0 // indirect
	github.com/ebitengine/purego v0.5.0 // indirect
	github.com/gammazero/deque v0.2.1 // indirect
	github.com/jezek/xgb v1.1.0 // indirect
	golang.org/x/exp/shiny v0.0.0-20230817173708-d852ddb80c63 // indirect
	golang.org/x/image v0.12.0 // indirect
	golang.org/x/mobile v0.0.0-20230922142353-e2f452493d57 // indirect
	golang.org/x/sync v0.3.0 // indirect
	golang.org/x/sys v0.12.0 // indirect
)

replace github.com/SnickeyX/roguelike/world => ./world

replace github.com/SnickeyX/roguelike/state => ./state

replace github.com/SnickeyX/roguelike/utils => ./utils
