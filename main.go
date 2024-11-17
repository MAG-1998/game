package main

import (
	"fmt"

	"github.com/MAG-1998/game/actions"
	"github.com/MAG-1998/game/interaction"
)

var currentround = 0
var gamerounds = []interaction.Rounddata{}

var PlayerAttackDmg int
var Playerhealvalue int
var Monsterattackdmg int
var Monsterheal int

func main() {
	startgame()

	winner := "" //player || monster || draw
	loser := ""
	for winner == "" {
		winner, loser = executeround()
	}

	endgame(winner, loser)
}

func startgame() {
	interaction.PrintGreeting()
}

func executeround() (winner string, loser string) {
	currentround++
	var playerhealth int
	var monsterhealth int
	IsSpecialround := currentround%actions.PLAYER_SPECIALATTACKROUND == 0
	is777 := currentround%actions.MONSTER_SPECIALATTACKROUND == 0

	interaction.Showavailableactions(IsSpecialround)
	userchoice := interaction.GetUserChoice(IsSpecialround)

	if userchoice == "1" {
		PlayerAttackDmg = actions.AttackMonster(IsSpecialround, is777)
	} else if userchoice == "2" {
		fmt.Println("Ням Ням")
		Playerhealvalue = actions.HealPlayer()
	} else if userchoice == "3" {
		PlayerAttackDmg = actions.AttackMonster(IsSpecialround, is777)
		fmt.Println("Хеликоптер---Хеликоптер")
	}
	Monsterattackdmg, Monsterheal = actions.AttackPlayer(is777)
	playerhealth, monsterhealth = actions.Gethealthamounts()
	rounddata := interaction.Rounddata{
		Action:           userchoice,
		PlayerHealth:     playerhealth,
		MonsterHealth:    monsterhealth,
		PlayerAttackDmg:  PlayerAttackDmg,
		Monsterattackdmg: Monsterattackdmg,
		Playerheal:       Playerhealvalue,
		Monsterheal:      Monsterheal,
	}
	interaction.PrintRoundstatistics(&rounddata)
	gamerounds = append(gamerounds, rounddata)

	if playerhealth <= 0 {
		return "Ботаник", "Гопник"
	} else if monsterhealth <= 0 {
		return "Гопник", "Ботаник"
	}
	return "", ""
}

func endgame(winner string, loser string) {
	interaction.Declarewinner(winner, loser)
	interaction.Writelogfile(&gamerounds)
}
