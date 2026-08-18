package main

import (
	"fmt"
	"os"
	"strings"
	"bufio"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

type config struct {
	commands map[string]cliCommand
	next *string
	previous *string
}

func commandExit(cfg *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name: "exit",
			description: "Exit the Pokedex",
			callback: commandExit,
		},
		"help": {
			name: "help",
			description: "Displays a help message",
			callback: commandHelp,
		},
		"map": {
			name:        "map",
			description: "Displays the next 20 location areas",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description:  "Displays the previous 20 location areas",
			callback:     commandMapb,
		},
	}
}

func commandHelp(cfg *config) error{
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()

	for _, cmd := range cfg.commands {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}

	return nil
}



func cleanInput(text string) []string {
	text = strings.ToLower(text)
	return strings.Fields(text)
}

func startRepl(cfg *config) {
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
		command, exists := cfg.commands[commandName]
		if !exists{
			fmt.Println("Unknown command")
			continue
		}
		err := command.callback(cfg)
		if err != nil{
			fmt.Println("Err:", err)
		}
	}
}
