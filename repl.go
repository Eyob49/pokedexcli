package main

import (
	"fmt"
	"os"
	"strings"
	"bufio"
	"github.com/Eyob49/pokedexcli/internal/pokecache"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

type config struct {
	commands map[string]cliCommand
	next     *string
	previous *string
	cache    *pokecache.Cache
}

func commandExit(cfg *config, args ...string) error {
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
		"explore": {
			name:         "explore",
			description:  "Displays a list of all the Pokémon located there",
			callback:     commandExplore,
		},
	}
}

func commandHelp(cfg *config, args ...string) error{
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
		err := command.callback(cfg, cleanedInput[1:]...)
		if err != nil{
			fmt.Println("Err:", err)
		}
	}
}
