package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// PokeAreas represents the structure of the PokeAPI location-area response
type PokeAreas struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func getAreas(url string) ([]struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}, *PageConfig, error) {

	resp, err := http.Get(url)
	if err != nil {
		return nil, nil, fmt.Errorf("error making request: %v", err)
	}
	defer resp.Body.Close()

	var pokeAreas PokeAreas
	err = json.NewDecoder(resp.Body).Decode(&pokeAreas)
	if err != nil {
		return nil, nil, fmt.Errorf("JSON decode error: %v", err)
	}

	cfg := &PageConfig{
		Next:     pokeAreas.Next,
		Previous: pokeAreas.Previous,
	}

	return pokeAreas.Results, cfg, nil
}
