package interaction

import (
	"fmt"
	"strings"
)

func PrintGreeting() {
	fmt.Println("ХОЧЕШЬ ИГРАТЬ НАЦЧИСЬ ЗДОРОВАТЬСЯ НАПИШИ 'ПРИВЕТ'")
	Privet, err := reader.ReadString('\n')
	if err != nil {fmt.Println("error")}
	Privet = strings.TrimSpace(Privet)
	if Privet != "ПРИВЕТ" || Privet != "Привет" || Privet != "привет" || Privet != "Privet"|| Privet != "privet"|| Privet != "PRIVET" {
		 fmt.Print("раз такие дела приди когда научишься писать привет"); return}
	fmt.Println("MONSTER SLAYER")
	fmt.Println(" НОВАЯ ИГРА НАЧИНАЕТСЯ...")
	fmt.Println("УДАЧИ САЛАГИ")
}
func Showavailableactions(Specialattackisavailable bool) {
 fmt.Println("Выбери че надо")
 fmt.Println("--------------")
 fmt.Println("(1) Гасим лоха")
 fmt.Println("(2) Семки")

 if  Specialattackisavailable {
	fmt.Println("(3) Пинок с разворота")
 }
}