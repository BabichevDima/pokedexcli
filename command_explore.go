package main

import (
	"fmt"
)

func commandExplore(cfg *config, locationArea string) error{
	fmt.Printf("Exploring %v...\n", locationArea)

	explorePokemonResp, err := cfg.pokeapiClient.ListExplorePokemon(locationArea)
	if err != nil {
		return err
	}

	fmt.Println()
	fmt.Println("Found Pokemon:")
	for _, pokemonEncounter := range explorePokemonResp.PokemonEncounters {
		fmt.Println("- " + pokemonEncounter.Pokemon.Name)
	}
	fmt.Println()
	return nil
}