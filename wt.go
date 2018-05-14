package main 

import "fmt"
import "bufio"
import "os"
import st "strings"

type Clone struct {
		Name string
		Gender string
		HP int
		MP int
		XP int
		Dexterity int
		Strength int
		Intelligence int
		Moves int
		Abilities []AbilityType
		Passives []PassiveType
	}

type Coord struct {
	x int
	y int
}

type Squad struct {
	Roster [][]Clone
}

var empty = Clone{
	Name: "   "}


func main() {
	basicAttack := BasicAttack{}
	zrardi.Abilities = append(zrardi.Abilities, basicAttack)
	sam.Abilities = append(sam.Abilities, basicAttack)
	diane.Abilities = append(diane.Abilities, basicAttack)

	hoistShield := HoistShield{}
	sam.Passives = append(sam.Passives, hoistShield)


	enemies := [][]Clone{
		{empty, empty, empty, empty},
		{zrardi, empty, empty, empty},
		{empty, zrardi, empty, empty},
		{empty, empty, zrardi, zrardi},
	}

	allies := [][]Clone{
		{sam, empty, empty, empty},
		{empty, empty, sam, sam},
		{empty, sam, empty, empty},
		{empty, empty, empty, empty},
	}
	team := Squad{Roster: allies}
	enemy := Squad{Roster: enemies}

	//team_print(enemy)
	//swap(enemy,Coord{0,3},Coord{3,3})
	//team_print(team)

	//enemy.Roster[1][0].Abilities = append(enemy.Roster[1][0].Abilities,basicAttack)
	//enemy.Roster[1][0].Abilities[0].execute(enemy,team,Coord{1,0},Coord{0,0})

	//fmt.Println(team.Roster[0][0].Name)
	//fmt.Println(team.Roster[0][0].HP)

	for (!team_death(team) || !team_death(enemy)) {
		team_print(enemy)
		fmt.Println()
		team_print(team)
		fmt.Printf("Enter move: ")
		reader := bufio.NewReader(os.Stdin)
		text, _ := reader.ReadString('\n')
		text = st.TrimRight(text,"\n")
		text = st.TrimRight(text,"\n")
		//fmt.Println(text[3])
		//fmt.Println(st.Compare(text,"move"))

		if (st.Contains(text,"move")){
			fmt.Println("Moving...")
		} else if (st.Contains(text,"attack")){
			fmt.Println("Attacking...")
		}


		for index_x, clone_column := range enemy.Roster {
			fmt.Println("DEBUG: X loop working ok...")
			for index_y, clone := range clone_column {
				fmt.Println("DEBUG: Y loop working ok... ",index_y)
				if !checkempty(enemy, Coord{index_x,index_y}) {
					fmt.Println("DEBUG: checkempty working as intended...")
					clone.Abilities[0].execute(enemy,team,Coord{index_x,index_y},melee_choose(clone.Abilities[0].targets(team)))
				}
			}
		}
	}
}


//#PRINT
func team_print(team Squad){
	for y := 0; y<4; y++ {
		for x:=0; x<4; x++ {
			fmt.Printf("|%-6v|", team.Roster[y][x].Name)
		}
		fmt.Printf("\n")
		//For formatting another space
		//Learning C was actually pretty useful here.
		fmt.Println("")
	}
}


//#SWAP
func swap(team Squad, t1 Coord, t2 Coord){
	interim := team.Roster[t1.y][t1.x]
	team.Roster[t1.y][t1.x] = team.Roster[t2.y][t2.x]
	team.Roster[t2.y][t2.x] = interim
}

func flipped_copy_of(team Squad) Squad {
	copyteam := team 
	//3 rows are flipped with 4,3,2 of the first columns swapped in each one.
	for y := 0; y<3; y++ {
		//diminish number of columns per row
		for x:=3; x>0; x++ {
			swap(copyteam,Coord{x,y},Coord{3-x,3-y})
		}
	}
	return copyteam
}

func clone_death(team Squad, t1 Coord){
	team.Roster[t1.y][t1.x] = empty
}


func checkempty(team Squad, t1 Coord) bool{
	if team.Roster[t1.y][t1.x].Name == "   " {
		return true
	} else {
		return false
	}
}	

func team_death(team Squad) bool{
	res := true
	for y:= 0; y<4; y++ {
		for x := 0; x<4; x++ {
			if !checkempty(team,Coord{y,x}) {
				res = false
			}
		}
	}
	return res
}


//summon only works if the object square is in fact empty.
func summon(team Squad, t1 Coord, object Clone) {

	if (checkempty(team, t1)) {
		team.Roster[t1.x][t1.y] = object
	}
}