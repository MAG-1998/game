package actions

import (
	"math/rand"
	"time"
	
)

var randsource = rand.NewSource(time.Now().UnixNano())
var rangenerator = rand.New(randsource)
var Currentmonsterhealth = 100
var Currentplayerhealth = 100

func AttackMonster(isspecialattack bool) {
	minAttackValue:= 5
	maxAttackValue:= 10

	if isspecialattack {
		minAttackValue = 10
		maxAttackValue = 20
	}
	dmgvalue := generaterandombetween(minAttackValue, maxAttackValue)
	Currentmonsterhealth -= dmgvalue // currentmonsterhealth = currentmonsterhealt - dmgvalue

}

func HealPlayer(healoption string) {
	
	
	minhealvalue := 10
	maxhealvalue := 20
	healvalue := generaterandombetween(minhealvalue,maxhealvalue)
	healthdiff := 100 - Currentplayerhealth
	if healthdiff>=healvalue{
		Currentplayerhealth+=healvalue
	} else {
		Currentplayerhealth = 100
	}
	
}

func AttackPlayer() {

}

func generaterandombetween(min int, max int) int {
	return rangenerator.Intn(max-min) + min
}
