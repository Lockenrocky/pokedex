package main

import (
	"errors"
	"fmt"
)

func commandInspect(cfg *config, pokemon_name string) error {

	if pokemon_name == "" {
		return errors.New("no pokemon name was provided! Pleas provide an pokemon name or pokemon id")
	}

	value, ok := cfg.pokedex[pokemon_name]
	if !ok {
		fmt.Print("you have not caught that pokemon")
		return nil
	}

	fmt.Printf("Name: %v\n", value.Forms[0].Pokemon_Name)
	fmt.Printf("Height: %v\n", value.Height)
	fmt.Printf("Weight: %v\n", value.Weight)
	fmt.Println("Stats:")
	for _, values := range value.Stats {
		fmt.Printf(" - %v: %v\n", values.Stat.Stat_Name, values.Base_stat)
	}
	fmt.Println("Types:")
	for _, values := range value.Types {
		fmt.Printf(" - %v\n", values.Type.Type_Name)
	}

	return nil
}
