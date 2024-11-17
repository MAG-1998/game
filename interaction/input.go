package interaction

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func GetUserChoice(Specialattackisavailable bool) string {
	for {
		Playerchoice, _ := GetPlayerInput()
		if Playerchoice == "1" {
			return "1"
		} else if Playerchoice == "2" {
			return "2"
		} else if Playerchoice == "3" && Specialattackisavailable {
			return "3"
		} else if Playerchoice == "3" && !Specialattackisavailable {
			fmt.Print("\nнедостаточно сигма\n")

		} else {
			fmt.Print("\nнорм вещи пиши э!\n")
		}
	}
}

func GetPlayerInput() (string, error) {
	fmt.Print("Выбор за тобой юный гопник: ")

	userInput, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}

	userInput = strings.TrimSpace(userInput)
	return userInput, nil
}
