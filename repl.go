package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/Lockenrocky/pokedexcli/internal/pokeapi"
)

type config struct {
	pokeapiClient    pokeapi.Client
	nextLocationsURL *string
	prevLocationsURL *string
	areaNameURL      *string
	pokedex          map[string]pokeapi.RespShallowPokemon
}

func startRepl(cfg *config) {

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		words := CleanInput(scanner.Text())
		if len(words) == 0 {
			continue
		}
		commandName := words[0]
		area_name := ""

		if len(words) == 2 {
			area_name = words[1]
		}

		command, exits := getCommands()[commandName]

		if exits {
			err := command.callback(cfg, area_name)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unkown command")
			continue
		}
	}
}

func CleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

type cliCommand struct {
	name        string
	description string
	callback    func(cfg *config, area_name string) error
}

func getCommands() map[string]cliCommand {

	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Displays 20 names of location areas in the Pokemon world",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the previous 20 names of location areas in the Pokemon world",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore",
			description: "Display all Pokemons for location",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Try to catch a Pokemon",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "Inspects a Pokemon",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "displays all catched Pokemons",
			callback:    commandPokedex,
		},
	}
}
