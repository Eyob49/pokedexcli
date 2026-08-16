package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
    commands := getCommands()
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		if !scanner.Scan() {
			if err := scanner.Err(); err != nil {
				fmt.Fprintln(os.Stderr, "Error reading input:", err)
			}
			return
		}
		line := scanner.Text()
		cleanedInput := cleanInput(line)
		if len(cleanedInput) == 0 {
			continue
		}

		commandName := cleanedInput[0]
		command, exists := commands[commandName]
		if !exists{
			fmt.Println("Unknown command")
			continue
		}
		err := command.callback()
		if err != nil{
			fmt.Println("Err:", err)
		}
	}
}
