package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

// CatchPokemon
func (c *Client) CatchPokemon(pageURL *string, pokemon_name *string) (RespShallowPokemon, error) {
	url := baseURL + "/pokemon/" + *pokemon_name
	if pageURL != nil {
		url = *pageURL
	}

	if val, ok := c.cache.Get(url); ok {
		locationsResp := RespShallowPokemon{}
		err := json.Unmarshal(val, &locationsResp)
		if err != nil {
			return RespShallowPokemon{}, err
		}

		return locationsResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespShallowPokemon{}, nil
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespShallowPokemon{}, nil
	}

	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespShallowPokemon{}, nil
	}

	locationsResp := RespShallowPokemon{}
	err = json.Unmarshal(dat, &locationsResp)
	if err != nil {
		return RespShallowPokemon{}, nil
	}

	c.cache.Add(url, dat)
	return locationsResp, nil

}
