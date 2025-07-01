package main

import (
	"fmt"
)

func commandPokedex(cfg *config, arg string) error{
    if len(cfg.caughtPokemon) == 0 {
        fmt.Println("You catch no pokemons yet!")
        return nil
    }
    
    fmt.Println("Your Pokedex:")
    for _, pokemon := range cfg.caughtPokemon {
        fmt.Printf(" - %s\n", pokemon.Name)
    }
	return nil
}

