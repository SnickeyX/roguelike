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

func InitializeWorld(startingLevel Level) (*ecs.Manager, map[string]ecs.Tag) {
	tags := make(map[string]ecs.Tag)
	manager := ecs.NewManager()

	//More stuff will go here
	player := manager.NewComponent()
	Position = manager.NewComponent()
	Rendarable = manager.NewComponent()
	movable := manager.NewComponent()

	// to ensure player always starts in a room and not a wall
	startingRoom := startingLevel.Rooms[0]
	x, y := startingRoom.Center()
	// init player location
	startingLevel.PlayerLoc[0], startingLevel.PlayerLoc[1] = x, y

	playerImg, _, err := ebitenutil.NewImageFromFile("assets/player.png")
	if err != nil {
		log.Fatal(err)
	}

	manager.NewEntity().
		AddComponent(player, utils.Player{}).
		AddComponent(Rendarable,
			&utils.Renderable{Image: playerImg}).
		AddComponent(movable, utils.Movable{}).
		AddComponent(Position, &utils.Position{
			// middle of the screen
			X: x,
			Y: y,
		})

	players := ecs.BuildTag(player, Position)
	tags["players"] = players

	renderables := ecs.BuildTag(Rendarable, Position)
	tags["renderables"] = renderables

	return manager, tags
}
