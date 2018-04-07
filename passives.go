package main

type Passive struct {}

type PassiveFunctions interface {
	name()
	description()
	reaction()
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

//Passives
type HoistShield struct {
	DamagePassive
}

