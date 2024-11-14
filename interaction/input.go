package interaction

import (
	"bufio"
	"fmt"
	"os"
	"strings"

)
var reader = bufio.NewReader(os.Stdin)


func GetUserChoice(Specialattackisavailable bool) {
	
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