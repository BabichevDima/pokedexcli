package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)
type configURL struct{
    Next		string
    Previous	string
}

type cliCommand struct {
	name        string
	description string
	callback    func(*configURL) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Displays the names of 20 location areas in the Pokemon world",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Get the previous page of locations",
			callback:    commandMapB,
		},
	}
}

func runREPL() {
	scanner := bufio.NewScanner(os.Stdin)
	urlParam := configURL{
		Next: "https://pokeapi.co/api/v2/location-area/?offset=0&limit=20",
		Previous: "",
	}

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		words := cleanInput(scanner.Text())

		if len(words) == 0 {
			continue
		}

		commandName := words[0]
		command, exists := getCommands()[commandName]

		if exists {
			err := command.callback(&urlParam)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}
	}
}

func cleanInput(text string) []string{
	lower := strings.ToLower(strings.TrimSpace(text))
	words := strings.Fields(lower)
	return words 
}