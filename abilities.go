package main

import "fmt"

type Ability struct{
	Cooldown int
}

type AbilityFunctions interface{
	name() string
	description() string
	check() bool
	targets() ([]int, []int)
	calculate()
	execute()
}


//import "fmt"
//Ability flowchart:
//AT ANY TIME YOU MUST BE ABLE TO CANCEL BEFORE CONFIRM
//1. Selected Ability returns available targets
//2. Ability executes on targets
//3. Effect calculation on each target with effect sequencing. eg. Explosive attack where primary target takes dmg first then the adjacent boards take damage simultaneously.  
//3.a this also includes passive calculations where counterattacks may happen and must be accounted for.

//func Ability(fp func() {

//}


//func Melee(fp func()) {
//	MeleeTargets(team Squad)
	//MeleeEffect()
//}


//List of selectable coordinates.  
//This doesn't save or calculate the aoes predictively for the user.
//It simply gives them options to choose.  We might add predictive aoes later.  

//Melee targets the "bottom" of each column in the array.  
//You're opponent will always be facing "above" I can go back and change it if for whatever reason, this isn't good enough.
//Famous last words for a coder. inb4 sphagetti code.
type BasicAttack struct {
	Ability
}

func (ba BasicAttack)  name() string {
	return "Basic Attack"
}

func (ba BasicAttack) description() string {
	return "Standard Attack based on Strength stat."
}

func (ba BasicAttack) check(clone Clone) bool {
	if (ba == BasicAttack || ba.Cooldown == 0){
		return true
	} else {
		return false
	}
}

func (ba BasicAttack) targets(enemy Squad) ([]int, []int) {
	var target_arr_x []int
	var target_arr_y []int
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
			target_arr_x = append(target_arr_x, y)
			target_arr_y = append(target_arr_y, column_target)
		}
	}
	return target_arr_x, target_arr_y
}

func (ba BasicAttack) calculate(team, Squad, enemy Squad, team_x int, team_y int, enemy_x int, enemy_y int) {
	var damage = team.Roster[team_x][team_y].Strength
	fmt.Println(damage)
}


func (ba BasicAttack) execute(team Squad, enemy Squad, team_x int, team_y int, enemy_x int, enemy_y int) {
	var damage = team.Roster[team_x][team_y].Strength
	enemy.Roster[enemy_x][enemy_y].HP -= damage
	ba.Cooldown = 0
}