package main

import (
	"bufio"
	"fmt"
	"os"

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
	currentround++
	IsSpecialround := currentround%3 == 0

	interaction.Showavailableactions(IsSpecialround)
	userchoice := interaction.GetUserChoice(IsSpecialround)
	fmt.Println(userchoice)
	if userchoice == "\nпинаем гада\n" {

	} else if userchoice == "\nням ням\n" {

	} else if userchoice == "\nхеликоптер хеликоптер\n" {

	}

	return ""
}

func endgame() {}
