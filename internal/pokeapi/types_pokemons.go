package pokeapi

// RespShallowLocations
type RespShallowPokemons struct {
	Pokemons []struct {
		Pokemon struct {
			Name string `json:"name"`
			Url  string `json:"url"`
		} `json:"pokemon"`
	} `json:"pokemon_encounters"`
}
