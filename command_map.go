package main

import (
	"errors"
	"fmt"	
)

func commandMap(cfg *config) error {		
	resp, err := cfg.pokeapiClient.ListLocationAreas(cfg.nextLocationAreaURL)

	if err != nil {
		return err
	}

	for _, area := range resp.Results {
		fmt.Printf("%s\n", area.Name)
	}
	cfg.nextLocationAreaURL = resp.Next
	cfg.previousLocationAreaURL = resp.Previous
	return nil
}

func commandMapback(cfg *config) error {
	if cfg.previousLocationAreaURL == nil {
		return errors.New("You are on the first page.")
	}		
	resp, err := cfg.pokeapiClient.ListLocationAreas(cfg.previousLocationAreaURL)

	if err != nil {
		return err
	}

	for _, area := range resp.Results {
		fmt.Printf("%s\n", area.Name)
	}
	cfg.nextLocationAreaURL = resp.Next
	cfg.previousLocationAreaURL = resp.Previous
	return nil
}