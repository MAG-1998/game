package main

import (
	"bufio"
	"fmt"
	"os"
	"github.com/MAG-1998/game/actions"
	"github.com/MAG-1998/game/interaction"

)

var reader = bufio.NewReader(os.Stdin)
var currentround = 0

func main() {
	startgame()

	winner := "" //player || monster || draw

	for winner == "" {
		winner = executeround()
	}

	endgame()
}

func startgame() {
	interaction.PrintGreeting()
}

func executeround() string {
	fmt.Println("------------- Botanik HP:",actions.Currentmonsterhealth,"-----------------Gopnik HP:", actions.Currentplayerhealth)
	currentround++
	IsSpecialround := currentround%3 == 0

	interaction.Showavailableactions(IsSpecialround)
	userchoice := interaction.GetUserChoice(IsSpecialround)
	
	if userchoice == "1" {
		fmt.Println("Пинай гада!")
		actions.AttackMonster(IsSpecialround)
	} else if userchoice == "2" {
		fmt.Println("Ням Ням")
	} else if userchoice == "3" {
actions.AttackMonster(IsSpecialround)
fmt.Println("Хеликоптер---Хеликоптер")
	}

	return ""
}

func endgame() {}
