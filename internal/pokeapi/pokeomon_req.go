package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)


func (c *Client) GetPokemon(pokemonName string) (Pokemon, error) {
	endpoint := "/pokemon/" + pokemonName
	fullUrl := baseURL + endpoint

	data, ok := c.cache.Get(fullUrl)

	if ok {		
		pokemonData := Pokemon{}
		err := json.Unmarshal(data, &pokemonData)

		if err != nil {
			return Pokemon{}, err
		}

		return pokemonData, nil
	}	

	req, err := http.NewRequest("GET", fullUrl, nil)

	if err != nil {
		return Pokemon{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return Pokemon{}, err
	}

	defer resp.Body.Close()

	if resp.StatusCode > 399 {
		return Pokemon{}, fmt.Errorf("bad status code: %v", resp.StatusCode)
	}

	data, err = io.ReadAll(resp.Body)
	if err != nil {
		return Pokemon{}, err
	}

	pokemonData := Pokemon{}
	err = json.Unmarshal(data, &pokemonData)

	if err != nil {
		return Pokemon{}, err
	}

	c.cache.Add(fullUrl, data)

	return pokemonData, nil
}
