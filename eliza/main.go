package main

import (
	"bufio"
	"fmt"
	"myapp/doctor"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	whatToSay := doctor.Intro()

	fmt.Println(whatToSay)

	var userInput string
	for userInput != "quit" {
		fmt.Print("-> ")
		userInput, _ = reader.ReadString('\n')
		userInput = strings.TrimSpace(userInput)

		response := doctor.Response(userInput)
		fmt.Println(response)
	}
}
