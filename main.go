package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex >")
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
		fmt.Print("Your command was: ", cleanedInput[0], "\n")
	}
}
