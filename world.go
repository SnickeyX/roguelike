package main

// implementing the ECS system

import (
	"log"

	"github.com/bytearena/ecs"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var position *ecs.Component
var rendarable *ecs.Component

func InitializeWorld(startingLevel Level) (*ecs.Manager, map[string]ecs.Tag) {
	tags := make(map[string]ecs.Tag)
	manager := ecs.NewManager()

	//More stuff will go here
	player := manager.NewComponent()
	position = manager.NewComponent()
	rendarable = manager.NewComponent()
	movable := manager.NewComponent()

	// to ensure player always starts in a room and not a wall
	startingRoom := startingLevel.Rooms[0]
	x, y := startingRoom.Center()

	playerImg, _, err := ebitenutil.NewImageFromFile("assets/player.png")
	if err != nil {
		log.Fatal(err)
	}

	manager.NewEntity().
		AddComponent(player, Player{}).
		AddComponent(rendarable,
			&Renderable{Image: playerImg}).
		AddComponent(movable, Movable{}).
		AddComponent(position, &Position{
			// middle of the screen
			X: x,
			Y: y,
		})

	players := ecs.BuildTag(player, position)
	tags["players"] = players

	renderables := ecs.BuildTag(rendarable, position)
	tags["renderables"] = renderables

	return manager, tags
}
