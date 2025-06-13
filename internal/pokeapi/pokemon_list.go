package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

// ListPokemons
func (c *Client) ListPokemons(pageURL *string, area_name *string) (RespShallowPokemons, error) {
	url := baseURL + "/location-area/" + *area_name
	if pageURL != nil {
		url = *pageURL
	}

	if val, ok := c.cache.Get(url); ok {
		locationsResp := RespShallowPokemons{}
		err := json.Unmarshal(val, &locationsResp)
		if err != nil {
			return RespShallowPokemons{}, err
		}

		return locationsResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespShallowPokemons{}, nil
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespShallowPokemons{}, nil
	}

	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespShallowPokemons{}, nil
	}

	locationsResp := RespShallowPokemons{}
	err = json.Unmarshal(dat, &locationsResp)
	if err != nil {
		return RespShallowPokemons{}, nil
	}

	c.cache.Add(url, dat)
	return locationsResp, nil

}
