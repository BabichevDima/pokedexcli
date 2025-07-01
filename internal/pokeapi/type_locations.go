package pokeapi

type RespShallowLocations struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}


type RespPokemonEncounter struct {
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
	} `json:"pokemon_encounters"`
}

type Pokemon struct {
	Name           string `json:"name"`
    BaseExperience int    `json:"base_experience"`
    Height         int    `json:"height"`
    Weight         int    `json:"weight"`
    Order          int    `json:"order"`
    Types          []PokemonType `json:"types"`
    Stats          []PokemonStat `json:"stats"`
}

type PokemonType struct {
    Slot int `json:"slot"`
    Type struct {
        Name string `json:"name"`
    } `json:"type"`
}

type PokemonStat struct {
    BaseStat int `json:"base_stat"`
    Stat     struct {
        Name string `json:"name"`
    } `json:"stat"`
}