package main

import (
	"fmt"

	"github.com/SnickeyX/roguelike/state"
	"github.com/SnickeyX/roguelike/utils"
	"github.com/SnickeyX/roguelike/world"
	"github.com/bytearena/ecs"
)

func AttackSystem(g *Game, attackerQ *ecs.QueryResult, defenderQ *ecs.QueryResult) {
	attackerName := attackerQ.Components[world.Name].(*utils.Name).Label
	attackerWeapon := attackerQ.Components[world.MeleeWeapon].(*utils.MeleeWeapon)
	defenderName := defenderQ.Components[world.Name].(*utils.Name).Label
	defenderArmor := defenderQ.Components[world.Armor].(*utils.Armor)

	d12 := utils.GetDiceRoll(12)

	// hit
	if d12+attackerWeapon.BonusDmg > defenderArmor.Class {
		// variable total damage
		total_dmg := utils.GetRandomBetweenTwo(attackerWeapon.MinDmg, attackerWeapon.MaxDmg) - defenderArmor.Def
		if total_dmg < 0 {
			total_dmg = 0
		}
		defenderHealth := defenderQ.Components[world.Health].(*utils.Health)
		defenderHealth.CurrHP -= total_dmg
		LogMessage += fmt.Sprintf("%s swings %s at %s and hits for %d health.\n",
			attackerName, attackerWeapon.Name, defenderName, total_dmg)

		if defenderHealth.CurrHP <= 0 {
			LogMessage += fmt.Sprintf("%s has died!\n", defenderName)
			if defenderName == "Snickey" {
				LogMessage += "Snickey died .... Game Over!\n"
				g.Turn = state.GameOver
			}
			g.World.DisposeEntity(defenderQ.Entity)
			// unblock tile after monster dies
			pos := defenderQ.Components[world.Position].(*utils.Position)
			index := g.Map.CurrentLevel.GetIndexFromXY(pos.X, pos.Y)
			g.Map.CurrentLevel.Tiles[index].Blocked = false
		}
	} else {
		LogMessage += fmt.Sprintf("%s swings %s at %s and misses!.\n",
			attackerName, attackerWeapon.Name, defenderName)
	}

}
