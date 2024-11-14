package main

import (
	"bufio"
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
	isSpecialround := currentround%3 == 0
	
	interaction.Showavailableactions(isSpecialround)
}

func endgame() {}
