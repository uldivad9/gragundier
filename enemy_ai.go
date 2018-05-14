package main

//import "math"
import "math/rand"

//Let the enemy just mindlessly attack for now.  If I actually care about AI I'll do something else.

func melee_choose(coord_list []Coord) Coord{
	return coord_list[rand.Intn(len(coord_list))]
}