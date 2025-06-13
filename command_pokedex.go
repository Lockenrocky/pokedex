package main

import (
	"fmt"
)

func commandPokedex(cfg *config, _ string) error {

	fmt.Println("Your Pokedex: ")

	for _, pokemon := range cfg.pokedex {
		fmt.Printf(" - %v\n", pokemon.Forms[0].Pokemon_Name)
	}

	return nil
}
