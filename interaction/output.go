package interaction

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
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
		fmt.Println("УДАЧИ САЛАГИ")
		fmt.Println("Ты нашел гопника выбери свое действие:")
	} else {
		fmt.Print("раз такие дела приди когда научишься писать привет\n")
		os.Exit(0)
	}
}

func getRandomPhrase() string {
	phrases := []string{"Че будем?", "Делай! Делай!", "Ехалооо!", "Снова в бой, гопник!", "Ботаник снова на ногах", "Выбери че надо"}

	rand.Seed(time.Now().UnixNano())

	randomIndex := rand.Intn(len(phrases))
	return phrases[randomIndex]
}

func Showavailableactions(Specialattackisavailable bool) {
	fmt.Println(getRandomPhrase())
	fmt.Println("--------------")
	fmt.Println("(1) Гасим лоха")
	fmt.Println("(2) Семки")

	if Specialattackisavailable {
		fmt.Println("(3) Пинок с разворота")
	}
}
