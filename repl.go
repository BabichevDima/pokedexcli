package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"github.com/BabichevDima/pokedexcli/internal/pokeapi"
)
type config struct {
	pokeapiClient		pokeapi.Client
	nextLocationsURL	*string
	prevLocationsURL	*string
	caughtPokemon		map[string]pokeapi.Pokemon
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, string) error
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
		"explore": {
			name:        "explore",
			description: "It takes the name of a location area as an argument. Parse the Pokemon's names from the response and display them to the user",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "It takes the name of a Pokemon as an argument. Catch the Pokemon.",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "It takes the name of a Pokemon and prints the name, height, weight, stats and type(s) of the Pokemon.",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "It prints a list of all the names of the Pokemon the user has caught.",
			callback:    commandPokedex,
		},
	}
}

func runREPL(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)

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
			args := ""
			if len(words) > 1 {
				args = words[1]
			}
			err := command.callback(cfg, args)
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