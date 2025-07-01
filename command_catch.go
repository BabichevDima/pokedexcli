package main

import (
	"math/rand"
	"fmt"
	"time"
	"github.com/BabichevDima/pokedexcli/internal/pokeapi"
)

func commandCatch(cfg *config, pokemon string) error{
	if pokemon == "" {
		return fmt.Errorf("missing pokemon name")
	}

	fmt.Printf("Throwing a Pokeball at %v...\n", pokemon)

	pokemonInfoResp, err := cfg.pokeapiClient.PokemonInfo(pokemon)
	if err != nil {
		return err
	}

	baseExp := pokemonInfoResp.BaseExperience
    catchChance := calculateCatchChance(baseExp)

    rand.Seed(time.Now().UnixNano())
    roll := rand.Intn(100) + 1

    if roll <= catchChance {
        fmt.Printf("%s was caught!\n", pokemon)
        if cfg.caughtPokemon == nil {
            cfg.caughtPokemon = make(map[string]pokeapi.Pokemon)
        }
        cfg.caughtPokemon[pokemon] = pokemonInfoResp
        fmt.Println("You may now inspect it with the 'inspect' command.")  
    } else {
        fmt.Printf("%s escaped!\n", pokemon)
    }
	return nil
}

func calculateCatchChance(baseExp int) int {
    baseChance := 50
    reduction := baseExp / 10
    finalChance := baseChance - reduction

    if finalChance < 10 {
        return 10
    }
    return finalChance
}