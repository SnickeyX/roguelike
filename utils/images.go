package utils

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var FloorImg *ebiten.Image
var WallImg *ebiten.Image
var PlayerImg *ebiten.Image
var SkeleIdleImg *ebiten.Image
var SkeleBuffImg *ebiten.Image

// loading all assets if not loaded
func LoadAllAssets() {
	if FloorImg == nil {
		floor, _, err := ebitenutil.NewImageFromFile("assets/floor.png")
		if err != nil {
			log.Fatal(err)
		}
		FloorImg = floor
	}

	if WallImg == nil {
		wall, _, err := ebitenutil.NewImageFromFile("assets/wall.png")
		if err != nil {
			log.Fatal(err)
		}
		WallImg = wall
	}

	if PlayerImg == nil {
		player, _, err := ebitenutil.NewImageFromFile("assets/knight_idle_1.png")
		if err != nil {
			log.Fatal(err)
		}
		PlayerImg = player
	}

	if SkeleIdleImg == nil {
		skele_idle, _, err := ebitenutil.NewImageFromFile("assets/skele_idle_1.png")
		if err != nil {
			log.Fatal(err)
		}
		SkeleIdleImg = skele_idle
	}

	if SkeleBuffImg == nil {
		skele_buff, _, err := ebitenutil.NewImageFromFile("assets/skele_buff_1.png")
		if err != nil {
			log.Fatal(err)
		}
		SkeleBuffImg = skele_buff
	}
}
