package main

import (
	"time"

	"github.com/Lockenrocky/pokedexcli/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5*time.Second, time.Minute*5)
	cfg := &config{
		pokeapiClient: pokeClient,
		pokedex:       make(map[string]pokeapi.RespShallowPokemon),
	}
	startRepl(cfg)
}
