package main

func main () {
	startgame()

	winner := "" //player || monster || draw

	for winner == "" {
		winner=executeround()
	}

	endgame()
}

func startgame() {}

func executeround() string{}

func endgame() {}