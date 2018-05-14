package main

import "fmt"


type Ability struct{
	Cooldown int
}

type AbilityType interface{
	name() string
	description() string
	check(clone Clone) bool
	targets(enemy Squad) []Coord
	calculate(team Squad, enemy Squad, team_coord Coord, enemy_coord Coord)
	execute(team Squad, enemy Squad, team_coord Coord, enemy_coord Coord)
}

//Ability flowchart:
//AT ANY TIME YOU MUST BE ABLE TO CANCEL BEFORE CONFIRM
//1. Selected Ability returns available targets
//2. Ability executes on targets
//3. Effect calculation on each target with effect sequencing. eg. Explosive attack where primary target takes dmg first then the adjacent boards take damage simultaneously.  
//3.a this also includes passive calculations where counterattacks may happen and must be accounted for.

//List of selectable coordinates.  
//This doesn't save or calculate the aoes predictively for the user.
//It simply gives them options to choose.  We might add predictive aoes later.  

type BasicAttack struct {
	Ability
}

func (ba BasicAttack) name() string {
		return "Basic Attack"
	}

func (ba BasicAttack) description() string {
		return "Standard Attack based on Strength stat."
	}

func (ba BasicAttack) check(clone Clone) bool {
	if (ba.Cooldown == 0){
		return true
	} else {
		return false
	}
}

func (ba BasicAttack) targets(enemy Squad) []Coord {
	return_Coords := []Coord{}
	for y := 0; y<4; y++ {
		var column_target = 0
		var tick = 0
		for x := 0; x<4; x++ {
			if (enemy.Roster[x][y].Name != "   "){
				column_target = x 
				tick = 1
				//fmt.Println("Found a target", x)
			}
		}
		if tick == 1 {
			new_coord := Coord{column_target,y}
			return_Coords = append(return_Coords, new_coord)
		}
	}
	return return_Coords
}

func (ba BasicAttack) calculate(team Squad, enemy Squad, team_coord Coord, enemy_coord Coord) {
	var damage = team.Roster[team_coord.x][team_coord.y].Strength
	fmt.Println(damage)
}


func (ba BasicAttack) execute(team Squad, enemy Squad, team_coord Coord, enemy_coord Coord) {
	//calculate raw damage
	var damage = team.Roster[team_coord.x][team_coord.y].Strength
	//
	damage = damage_passive_calculate_target(enemy.Roster[enemy_coord.x][enemy_coord.y],damage)
	enemy.Roster[enemy_coord.x][enemy_coord.y].HP -= damage
	ba.Cooldown = 0
}
