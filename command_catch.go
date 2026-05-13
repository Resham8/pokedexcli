package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, args []string) error {
	if len(args) != 1 {
		return errors.New("No pokemon provided")
	}

	pokemonName := args[0]

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonName)

	pokemon, err := cfg.pokeapiClient.GetPokemon(pokemonName)

	if err != nil {
		return fmt.Errorf("pokemon %s does not exist", pokemonName)
	}

	const threashold = 40
	baseExp := pokemon.BaseExperience

	randomNum := rand.Intn(baseExp)

	if randomNum > threashold {
		return fmt.Errorf("failed to catch %s", pokemonName)
	}

	cfg.caughtPokemon[pokemonName] = pokemon
	fmt.Printf("%s was caught! \n", pokemonName)
	fmt.Println("You may now inspect it with the inspect command.")
	return nil
}
