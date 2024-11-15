package interaction

import (
	"fmt"
	"os"
	"strings"
	"github.com/MAG-1998/game/interaction/utils"
)

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

func Showavailableactions(Specialattackisavailable bool){
	fmt.Println(utils.GetRandomPhrase())

	fmt.Println("(1)", utils.GetRandomPhraseforattack())
	fmt.Println("(2)",utils.GetRandomPhraseforheal())

	if Specialattackisavailable {
		fmt.Println("(3) Пинок с разворота")
	}
}

