package main 

import "fmt"
import "bufio"
import "os"
import "strings"

type Clone struct {
		Name string
		Gender string
		HP int
		MP int
		Dexterity int
		Strength int
		Intelligence int
		Moves int
		Abilities []Ability
		Passives []Passive
	}

type Squad struct {
	Roster [][]Clone
}

var empty = Clone{
	Name: "   "}


func main() {

	enemies := [][]Clone{
		{empty, empty, empty, empty},
		{zrardi, empty, empty, empty},
		{empty, zrardi, empty, empty},
		{empty, empty, zrardi, zrardi},
	}

	team := [][]Clone{
		{sam, empty, empty, empty},
		{empty, empty, sam, sam},
		{empty, sam, empty, empty},
		{empty, empty, empty, empty},
	}
	wteam := Squad{Roster: team}
	enemy := Squad{Roster: enemies}

	team_print(enemy)
	swap(enemy,0,3,3,0)
	team_print(wteam)
	arrx, arry := MeleeTargets(enemy)
	fmt.Println(arrx)
	fmt.Println(arry)
	MeleeEffects(wteam, enemy, 0,0, 1,0)
	fmt.Println(enemy.Roster[1][0].Name)
	fmt.Println(enemy.Roster[1][0].HP)
	team_print(enemy)
	team_print(flipped_copy_of(enemy))
	
	while (enemy.Roster != [][]Clone || wteam.Roster != [][]Clone) {
		team_print(enemy)
		team_print(wteam)
		fmt.Println("Enter move: ")
		reader := bufio.NewReader(os.Stdin)
		text, _ := reader.ReadString('\n')
		if (text == 'm'){
			
		}
	}
}


//#PRINT
func team_print(team Squad){
	for x := 0; x<4; x++ {
		for y:=0; y<4; y++ {
			fmt.Printf(" %5v", team.Roster[x][y].Name)
		}
		fmt.Printf("\n")
		//For formatting another space
		//Learning C was actually pretty useful here.
		fmt.Println("")
	}
}


//#SWAP
func swap(team Squad, x1 int, y1 int, x2 int, y2 int){
	interim := team.Roster[x1][x2]
	team.Roster[x1][y1] = team.Roster[x2][y2]
	team.Roster[x2][y2] = interim
}

func flip_and_copy(team Squad){
	copyteam := team 
	//3 rows are flipped with 4,3,2 of the first columns swapped in each one.
	for (x := 0; x<3; x++) {
		//diminish number of columns per row
		for (y:=3; y>0; y++){
			swap(copyteam,x,y,3-x,3-y)
		}
	}
	return copyteam
}

func simple_death(team Squad, x1 int, y1 int){
	team.Roster[x1][y1] = empty
}

func checkempty(team Squad, x1 int , y1 int){
	if team.Roster[x1][y1].Name = "   " {
		return true
	}
	else {
		return false
	}
}


//summon only works if the object square is in fact empty.
func summon(team Squad, x1 int, y1 int, object Clone) {
	if (checkempty(team, x1, y1)) {
		team.Roster[x1][y1] = object
	}
}