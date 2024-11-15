package actions

import (
	"fmt"
	"math/rand"
	"time"
)

var randsource = rand.NewSource(time.Now().UnixNano())

var rangenerator = rand.New(randsource)
var Currentmonsterhealth = MONSTER_HEALTH
var Currentplayerhealth = PLAYER_HEALTH

func AttackMonster(isspecialattack bool) {
	minAttackValue := PLAYER_ATTACK_MIN_DMG
	maxAttackValue := PLAYER_ATTACK_MAX_DMG

	if isspecialattack {
		minAttackValue = PLAYER_SPECIAL_ATTACK_MIN_DMG
		maxAttackValue = PLAYER_SPECIAL_ATTACK_MAX_DMG
	}
	dmgvalue := generaterandombetween(minAttackValue, maxAttackValue)
	if Currentmonsterhealth <= dmgvalue {
		Currentmonsterhealth = 0
	} else {
		Currentmonsterhealth -= dmgvalue // currentmonsterhealth = currentmonsterhealt - dmgvalue
	}
}

func HealPlayer() {

	healvalue := generaterandombetween(PLAYER_HEAL_MIN, PLAYER_HEAL_MAX)
	healthdiff := PLAYER_HEALTH - Currentplayerhealth
	if healthdiff >= healvalue {
		Currentplayerhealth += healvalue
	} else {
		Currentplayerhealth = PLAYER_HEALTH
	}
}

func AttackPlayer(isspecialattack bool) {
	minAttackValue := MONSTER_ATTACK_MIN_DMG
	maxAttackValue := MONSTER_ATTACK_MAX_DMG

	if isspecialattack {
		minAttackValue = MONSTER_SPECIAL_ATTACK_MIN_DMG
		maxAttackValue = MONSTER_SPECIAL_ATTACK_MAX_DMG
		for i := 0; i < 800; i++ {
			fmt.Print("---Ботаник оказался не так уж и прост---получай пизды гопник---")

		}
		if (Currentmonsterhealth + 30) <= MONSTER_HEALTH {
			Currentmonsterhealth += 30
		} else {
			Currentmonsterhealth = MONSTER_HEALTH
		}
	}
	dmgvalue := generaterandombetween(minAttackValue, maxAttackValue)

	if Currentplayerhealth <= dmgvalue {
		Currentplayerhealth = LOSE_POINT
	} else {
		Currentplayerhealth -= dmgvalue
	}
}
func Gethealthamounts() (int, int) {
	return Currentplayerhealth, Currentmonsterhealth
}

func generaterandombetween(min int, max int) int {
	return rangenerator.Intn(max-min) + min
}
