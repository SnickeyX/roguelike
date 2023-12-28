package utils

import "github.com/hajimehoshi/ebiten/v2"

const ()

// ECS components

type Player struct{}

type MeleeWeapon struct {
	Name     string
	Aoe      int
	NumHits  int // number of monsters hit in a single turn
	Level    int
	MinDmg   int
	MaxDmg   int
	BonusDmg int // scale with level
}

type Armor struct {
	Name  string
	Def   int
	Class int // attack dmg must be more than this to hit
}

type Name struct {
	Label string
}

type Health struct {
	MaxHP  int
	CurrHP int
}

type Monster struct {
}

type Position struct {
	X int
	Y int
}

type Renderable struct {
	Image *ebiten.Image
}

type Movable struct{}
