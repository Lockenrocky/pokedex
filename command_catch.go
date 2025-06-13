package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, pokemon_name string) error {

	if pokemon_name == "" {
		return errors.New("no pokemon name was provided! Pleas provide an pokemon name or pokemon id")
	}

	pokemonResp, err := cfg.pokeapiClient.CatchPokemon(cfg.areaNameURL, &pokemon_name)
	if err != nil {
		return err
	}

	// Try to catch the pokemon

	fmt.Printf("Throwing a Pokeball at %v...\n", pokemon_name)

	// Calculate chance based on experience
	exp := pokemonResp.Exp
	r := rand.Intn(exp)
	if r > exp/2 {
		fmt.Printf("%s was caught!\nYou may now inspect it with the inspect command.\n", pokemon_name)
		cfg.pokedex[pokemon_name] = pokemonResp
	} else {
		fmt.Printf("%s escaped!\n", pokemon_name)
	}

	return nil
}
