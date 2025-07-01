package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) ListLocations(pageURL *string) (RespShallowLocations, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	if cache, ok := c.cache.Get(url); ok {
		locationsResp := RespShallowLocations{}
		err := json.Unmarshal(cache, &locationsResp)
		if err != nil {
			return RespShallowLocations{}, err
		}
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespShallowLocations{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespShallowLocations{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespShallowLocations{}, err
	}

	locationsResp := RespShallowLocations{}
	err = json.Unmarshal(dat, &locationsResp)
	if err != nil {
		return RespShallowLocations{}, err
	}


	c.cache.Add(url, dat)
	return locationsResp, nil
}

func (c *Client) ListExplorePokemon(locationArea string) (RespPokemonEncounter, error) {
	url := baseURL + "/location-area/"
	if locationArea != "" {
		url += locationArea
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespPokemonEncounter{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespPokemonEncounter{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespPokemonEncounter{}, err
	}

	pokemonEncounterResp := RespPokemonEncounter{}
	err = json.Unmarshal(dat, &pokemonEncounterResp)
	if err != nil {
		return RespPokemonEncounter{}, err
	}

	return pokemonEncounterResp, nil
}

func (c *Client) PokemonInfo(pokemon string) (Pokemon, error) {
	url := baseURL + "/pokemon/"
	if pokemon != "" {
		url += pokemon
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Pokemon{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return Pokemon{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return Pokemon{}, err
	}

	pokemonInfoResp := Pokemon{}
	err = json.Unmarshal(dat, &pokemonInfoResp)
	if err != nil {
		return Pokemon{}, err
	}

	return pokemonInfoResp, nil
}