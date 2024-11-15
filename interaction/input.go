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
			return "\nпинаем гада\n"
		} else if Playerchoice == "2"{
			return "\nням ням\n"
		} else if Playerchoice == "3" && Specialattackisavailable {
			return "\nхеликоптер хеликоптер\n"
		} else if Playerchoice=="3" && Specialattackisavailable == false{
			fmt.Print("\nнедостаточно сигма\n")
		} else {
			fmt.Print("\nнорм вещи пиши э!\n")
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