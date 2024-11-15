package actions

import (
	"math/rand"
	"time"
)

var randsource = rand.NewSource(time.Now().UnixNano())
var rangenerator = rand.New(randsource)

func AttackMonster(isspecialattack bool) {

}

func HealPlayer() {

}

func AttackPlayer() {

}

func generaterandom(min int, max int) int {
	return rangenerator.Intn(max-min) + min
}
