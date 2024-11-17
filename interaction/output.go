package interaction

import (
	"fmt"
	"os"
	"strings"

	"github.com/MAG-1998/game/actions"
	"github.com/MAG-1998/game/interaction/utils"
)

type Rounddata struct {
	Action           string
	PlayerAttackDmg  int
	Playerheal       int
	Monsterattackdmg int
	Monsterheal      int
	MonsterHealth    int
	PlayerHealth     int
}

func PrintGreeting() {
	fmt.Println("ХОЧЕШЬ ИГРАТЬ НАУЧИСЬ ЗДОРОВАТЬСЯ НАПИШИ 'ПРИВЕТ'")
	Privet, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("error")
	}
	Privet = strings.TrimSpace(Privet)
	if Privet == "ПРИВЕТ" || Privet == "Привет" || Privet == "привет" || Privet == "Privet" || Privet == "privet" || Privet == "PRIVET" {
		fmt.Println("--------------------Botanik Slayer----------------------")
		fmt.Println("НОВАЯ ИГРА НАЧИНАЕТСЯ...")
		fmt.Println("УДАЧИ САЛАГИ. ВПЕРЕДИ МНОГО ПРИКЛЮЧЕНИЙ!")
		fmt.Println("Ты нашел ботаника выбери свое действие:")
	} else {
		fmt.Print("раз такие дела приди когда научишься писать привет\n")
		os.Exit(0)
	}
}

func Showavailableactions(Specialattackisavailable bool) {
	fmt.Println(utils.GetRandomPhrase())

	fmt.Println("(1)", utils.GetRandomPhraseforattack())
	fmt.Println("(2)", utils.GetRandomPhraseforheal())

	if Specialattackisavailable {
		fmt.Println("(3) Пинок с разворота")
	}
}

func Declarewinner(winner string, loser string) {
	fmt.Println("___________________________________________________________________________________")
	fmt.Println("-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-Милиция приехала-+-+-+-+-+-+-+-+--+-+-+-+-+-+-+-+-\n-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-Драка закончилась-+-+-+-+-+-+-+-+--+-+-+-+-+-+-+-+-")
	fmt.Println("___________________________________________________________________________________")
	fmt.Println("Но одного отправили в больницу а другого в обезянник")
	fmt.Println("___________________________________________________________________________________")
	fmt.Printf("%vа в больницу а %vа в обезьянник", loser, winner)
}

func PrintRoundstatistics(rounddata *Rounddata) {
	if rounddata.Action == "1" {
		fmt.Printf("Гопник ударил Ботана на %v хп. \n", rounddata.PlayerAttackDmg)
	} else if rounddata.Action == "3" {
		fmt.Printf("Гопник захуярил Ботана с вертухи на %v хп. \n", rounddata.PlayerAttackDmg)
	} else {
		fmt.Printf("Гопник зарегенился на %v\n", rounddata.Playerheal)
	}
	
	if rounddata.Monsterattackdmg > actions.MONSTER_ATTACK_MAX_DMG {
		fmt.Printf("\nБотаник забуянился и пнул по яйцам урон колосальный на %v, и при этом прилив уверенности восстановил ему здоровье на %v\n", rounddata.Monsterattackdmg,rounddata.Monsterheal)
	} else {fmt.Printf("Ботан укусил в ответ на %v\n", rounddata.Monsterattackdmg)}
	fmt.Println("Здоровье Гопника: \n", rounddata.PlayerHealth)
	fmt.Println("Здоровье Ботана: \n", rounddata.MonsterHealth)
}
