package main

//Battlescript:
//Steps are basically each unit gets a "move" token
//Some units might get more others might get less
//We use the token system to determine who can "swap" or "move"

//Actions are also determined with "action" tokens
//Each unit on average should get a average number of action tokens though more might be given with abilities.
//Some abilities cost more action tokens then others
//Stats might determine how many action tokens are given later
//For now, I'm going to just give a generic average number and balance from there

//Abilities occur in order of command.  Once all moves are committed, the other side does their moves back and forth until the win condition is met
//win_condition check happens after every action.

var dead_team = [][]Clone {
	{empty, empty, empty, empty},
	{empty, empty, empty, empty},
	{empty, empty, empty, empty},
	{empty, empty, empty, empty},
}

var level1 = [][]Clone {
	{zrardi, zrardi, zrardi, zrardi},
	{zrardi, empty, empty, empty},
	{empty, empty, empty, empty},
	{empty, empty, empty, empty},
}