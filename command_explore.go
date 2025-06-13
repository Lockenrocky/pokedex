package main

import (
	"errors"
	"fmt"
)

func commandExplore(cfg *config, area_name string) error {

	if area_name == "" {
		return errors.New("no area name was provided! Pleas provide an area name or area id")
	}

	pokemonsResp, err := cfg.pokeapiClient.ListPokemons(cfg.areaNameURL, &area_name)
	if err != nil {
		return err
	}

	for _, pokemon := range pokemonsResp.Pokemons {
		fmt.Println(pokemon.Pokemon.Name)
	}
	return nil
}
