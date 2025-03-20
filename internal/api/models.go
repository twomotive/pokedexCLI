package api

// LocationAreasResponse represents the structure of the PokeAPI location-area response
type LocationAreasResponse struct {
	Count    int            `json:"count"`
	Next     *string        `json:"next"`
	Previous *string        `json:"previous"`
	Results  []LocationArea `json:"results"`
}

// LocationArea represents a single location area
type LocationArea struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

// PaginationConfig tracks pagination state
type PaginationConfig struct {
	Next     *string
	Previous *string
}

type PokemonLocationArea struct {
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
		} `json:"pokemon"`
	} `json:"pokemon_encounters"`
}
