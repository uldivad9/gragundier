package main

//import (
//	"reflect"
//)

type Passive struct {}

type PassiveType interface {
	name() string
	description() string
	reaction(clone Clone, damage int) int
}


type DamagePassive struct {
	Passive
}

type BuffPassive struct {
	Passive
}

type CCPassive struct {
	Passive
}

func damage_passive_calculate_target(clone Clone, damage int) int {
	var fully_calculated_damage = damage
	for _, passive := range clone.Passives {
			fully_calculated_damage = passive.reaction(clone,fully_calculated_damage)
	}
	return fully_calculated_damage
}

//Passives

//HoistShield
//Sam's Defensive Passive
type HoistShield struct {
	DamagePassive
}

func (hs HoistShield) name() string {
	return "Hoist Shield"
}

func (hs HoistShield) description() string {
	return "Upon Damage, decreases all incoming damage by half (rounded down) for one turn (CD: 3 turns)"
}

func (hs HoistShield) reaction(clone Clone, damage int) int{
	return int(damage/2)
}

//Frenzy
//Zrardi's Melee Buff Passive
type Frenzy struct {
	DamagePassive
}

func (fr Frenzy) name() string {
	return "Frenzy"
}

func (fr Frenzy) description() string {
	return "Upon Dealing Damage, Zrardi's Strength grows by 10 indefinitely"
}

func (fr Frenzy) reaction(clone Clone, damage int) int{
	clone.Strength += int(5)
	return damage
}

