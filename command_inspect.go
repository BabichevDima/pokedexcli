package main

import (
	"fmt"
)

func commandInspect(cfg *config, pokemon string) error{
	fmt.Println("------------------------------")
    fmt.Printf("You cautch: %v pokemons\n", len(cfg.caughtPokemon))
	fmt.Println("------------------------------")
	pokemonData, ok := cfg.caughtPokemon[pokemon]
	if !ok {
		return fmt.Errorf("You haven't caught %s yet", pokemon)
	}
	
    fmt.Println("==========================")
    fmt.Printf("Name: %s\n", pokemonData.Name)
    fmt.Printf("Base Experience: %d\n", pokemonData.BaseExperience)
    fmt.Printf("Height: %d\n", pokemonData.Height)
    fmt.Printf("Weight: %d\n", pokemonData.Weight)

    fmt.Println("Types:")
    for _, t := range pokemonData.Types {
        fmt.Printf("  - %s (slot %d)\n", t.Type.Name, t.Slot)
    }

    fmt.Println("Stats:")
    for _, s := range pokemonData.Stats {
        fmt.Printf("  - %s: %d\n", s.Stat.Name, s.BaseStat)
    }
    fmt.Println("==========================")
	return nil
}

