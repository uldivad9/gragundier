package main

var sam = Clone {
	Name: "Sam",
	Gender: "Male",
	HP: 50,
	MP: 50,
	Dexterity: 10,
	Strength: 20,
	Intelligence: 10,
	Moves: 0,
	Abilities: []Ability{BasicAttack},
	Passives: []Passive{HoistShield},
}

var diane = Clone {
	Namae: "Diane",
	Gender: "Female",
	HP: 40,
	MP: 40,
	Dexterity: 20,
	Strength: 10,
	Intelligence: 10,
	Moves: 0,
	Abilities: []Ability{BasicAttack},
	Passives: []Passive{},
}

