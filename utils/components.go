package utils

import "github.com/hajimehoshi/ebiten/v2"

// ECS components

type Player struct{}

type Monster struct {
	Name string
}

type Position struct {
	X int
	Y int
}

type Renderable struct {
	Image *ebiten.Image
}

type Movable struct{}
