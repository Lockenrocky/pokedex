package pokeapi

// RespShallowLocations
type RespShallowPokemon struct {
	Exp   int `json:"base_experience"`
	Forms []struct {
		Pokemon_Name string `json:"name"`
	} `json:"forms"`
	Height int `json:"height"`
	Stats  []struct {
		Base_stat int `json:"base_stat"`
		Stat      struct {
			Stat_Name string `json:"name"`
		} `json:"stat"`
	} `json:"stats"`
	Types []struct {
		Type struct {
			Type_Name string `json:"name"`
		} `json:"type"`
	} `json:"types"`
	Weight int `json:"weight"`
}
