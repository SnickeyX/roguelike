package world

// implementing the ECS system

import (
	"log"

	"github.com/SnickeyX/roguelike/utils"
	"github.com/bytearena/ecs"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

// captialised for Go to allow exporting
var Position *ecs.Component
var Rendarable *ecs.Component
var Monster *ecs.Component

func InitializeWorld(startingLevel Level) (*ecs.Manager, map[string]ecs.Tag) {
	tags := make(map[string]ecs.Tag)
	manager := ecs.NewManager()

	//More stuff will go here
	player := manager.NewComponent()
	Monster = manager.NewComponent()
	Position = manager.NewComponent()
	Rendarable = manager.NewComponent()
	movable := manager.NewComponent()

	// to ensure player always starts in a room and not a wall
	startingRoom := startingLevel.Rooms[0]
	x, y := startingRoom.Center()
	// init player location
	startingLevel.PlayerLoc[0], startingLevel.PlayerLoc[1] = x, y

	playerImg, _, err := ebitenutil.NewImageFromFile("assets/knight_idle_1.png")
	if err != nil {
		log.Fatal(err)
	}

	skellyImg, _, err := ebitenutil.NewImageFromFile("assets/skele_idle_1.png")
	if err != nil {
		log.Fatal(err)
	}

	manager.NewEntity().
		AddComponent(player, &utils.Player{}).
		AddComponent(Rendarable,
			&utils.Renderable{Image: playerImg}).
		AddComponent(movable, utils.Movable{}).
		AddComponent(Position, &utils.Position{
			// middle of the screen
			X: x,
			Y: y,
		})

	for _, room := range startingLevel.Rooms {
		if room.X1 != startingRoom.X1 {
			cX, cY := room.Center()
			manager.NewEntity().
				AddComponent(Monster, &utils.Monster{Name: "Skeleton"}).
				AddComponent(Rendarable,
					&utils.Renderable{Image: skellyImg}).
				AddComponent(Position, &utils.Position{
					// middle of the screen
					X: cX,
					Y: cY,
				})
		}
	}

	monsters := ecs.BuildTag(Monster, Position)
	tags["monsters"] = monsters

	players := ecs.BuildTag(player, Position)
	tags["players"] = players

	renderables := ecs.BuildTag(Rendarable, Position)
	tags["renderables"] = renderables

	return manager, tags
}
