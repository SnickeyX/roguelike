package world

// implementing the ECS system

import (
	"fmt"

	"github.com/SnickeyX/roguelike/utils"
	"github.com/bytearena/ecs"
)

// captialised for Go to allow exporting
var Position *ecs.Component
var Rendarable *ecs.Component
var Monster *ecs.Component
var Health *ecs.Component
var MeleeWeapon *ecs.Component
var Armor *ecs.Component
var Name *ecs.Component

func InitializeWorld(startingLevel Level) (*ecs.Manager, map[string]ecs.Tag) {
	tags := make(map[string]ecs.Tag)
	manager := ecs.NewManager()

	player := manager.NewComponent()
	Monster = manager.NewComponent()
	Position = manager.NewComponent()
	Rendarable = manager.NewComponent()
	movable := manager.NewComponent()
	Health = manager.NewComponent()
	MeleeWeapon = manager.NewComponent()
	Armor = manager.NewComponent()
	Name = manager.NewComponent()

	// to ensure player always starts in a room and not a wall
	startingRoom := startingLevel.Rooms[0]
	x, y := startingRoom.Center()
	// init player location
	startingLevel.PlayerLoc[0], startingLevel.PlayerLoc[1] = x, y

	manager.NewEntity().
		AddComponent(player, &utils.Player{}).
		AddComponent(Name, &utils.Name{Label: "Snickey"}).
		AddComponent(Armor, &utils.Armor{Name: "chainmail", Def: 1, Class: 1}).
		AddComponent(Health, &utils.Health{MaxHP: 10, CurrHP: 10}).
		AddComponent(MeleeWeapon,
			&utils.MeleeWeapon{Name: "fists", Level: 1, MinDmg: 1,
				MaxDmg: 3, BonusDmg: 2, Aoe: 1, NumHits: 1}).
		AddComponent(Rendarable,
			&utils.Renderable{Image: utils.PlayerImg}).
		AddComponent(movable, utils.Movable{}).
		AddComponent(Position, &utils.Position{
			// middle of the screen
			X: x,
			Y: y,
		})

	skele_num := 0
	for _, room := range startingLevel.Rooms {
		if room.X1 != startingRoom.X1 {
			skele_num++
			cX, cY := room.Center()
			manager.NewEntity().
				AddComponent(Monster, &utils.Monster{}).
				AddComponent(Name, &utils.Name{Label: fmt.Sprintf("Skeleton%v", skele_num)}).
				AddComponent(Armor, &utils.Armor{Name: "bones", Def: 1, Class: 1}).
				AddComponent(Health, &utils.Health{MaxHP: 5, CurrHP: 5}).
				AddComponent(MeleeWeapon,
					&utils.MeleeWeapon{Name: "bone", Level: 1, MinDmg: 0,
						MaxDmg: 2, BonusDmg: 1, Aoe: 1, NumHits: 1}).
				AddComponent(Rendarable,
					&utils.Renderable{Image: utils.SkeleIdleImg}).
				AddComponent(Position, &utils.Position{
					// middle of the screen
					X: cX,
					Y: cY,
				})
		}
	}

	monsters := ecs.BuildTag(Monster, Name, Position, Rendarable, Health, MeleeWeapon, Armor)
	tags["monsters"] = monsters

	players := ecs.BuildTag(player, Name, Position, Health, MeleeWeapon, Armor)
	tags["players"] = players

	renderables := ecs.BuildTag(Rendarable, Position)
	tags["renderables"] = renderables

	return manager, tags
}
