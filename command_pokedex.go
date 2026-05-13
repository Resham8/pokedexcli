package main

import (	
	"fmt"
)

func commandPokedex(cfg *config, args []string) error {
	pokemons := cfg.caughtPokemon

	fmt.Println("Your Pokedex:")
	for _, pokemon := range pokemons {
		fmt.Printf("  - %s\n", pokemon.Name)
	}	

	return nil
}
