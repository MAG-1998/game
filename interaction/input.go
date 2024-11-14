package interaction

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)
var reader = bufio.NewReader(os.Stdin)


func GetUserChoice(Specialattackisavailable bool) string{
	for {
	Playerchoice, _ := GetPlayerInput()
		if Playerchoice == "1"{
			return "пинаем гада"
		} else if Playerchoice == "2"{
			return "ням ням"
		} else if Playerchoice == "3" && Specialattackisavailable {
			return "хеликоптер хеликоптер"
		} else if Playerchoice=="3" && Specialattackisavailable == false{
			fmt.Println("недостаточно сигма")
		} else {
			fmt.Println("норм вещи пиши э!")
		}
		}
}	
	
	

func GetPlayerInput() (string,error) {
	fmt.Println("Выбор за тобой юный гопник: ")

	userInput, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}

	userInput=strings.TrimSpace(userInput)
	return userInput, nil
}