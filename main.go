package main

import (
	"fmt"

	"github.com/MAG-1998/game/actions"
	"github.com/MAG-1998/game/interaction"
)

var currentround = 0

func main() {
	startgame()

	winner := "" //player || monster || draw
	loser := ""
	for winner == "" {
		winner, loser = executeround()
	}

	endgame(winner,loser)
}

func startgame() {
	interaction.PrintGreeting()
}

func executeround() (winner string,loser string) {
	currentround++
	var playerhealth int
	var monsterhealth int
	IsSpecialround := currentround%3 == 0
	is777 := currentround%13 == 0

	interaction.Showavailableactions(IsSpecialround)
	userchoice := interaction.GetUserChoice(IsSpecialround)

	if userchoice == "1" {
		fmt.Println("Пинай гада!")
		 actions.AttackMonster(IsSpecialround)
	} else if userchoice == "2" {
		fmt.Println("Ням Ням")
		 actions.HealPlayer()
	} else if userchoice == "3" {
		actions.AttackMonster(IsSpecialround)
		fmt.Println("Хеликоптер---Хеликоптер")
	}
	actions.AttackPlayer(is777)
	playerhealth, monsterhealth = actions.Gethealthamounts()
	fmt.Printf("================HP Гопника : %v============= HP Ботаника: %v================", playerhealth, monsterhealth)

	if playerhealth <=0 {
		return "Ботаник", "Гопник"
	} else if monsterhealth<=0 {
		return "Гопник", "Ботаник"
	}
	return "",""
}

func endgame(winner string, loser string) {
	interaction.Declarewinner(winner, loser)
}
