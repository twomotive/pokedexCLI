package main

import "fmt"

func commandMap(cfg *PageConfig) error {
	url := "https://pokeapi.co/api/v2/location-area"
	if cfg.Next != nil {
		url = *cfg.Next
	}

	results, newCfg, err := getAreas(url)
	if err != nil {
		return fmt.Errorf("cannot map: %v", err)
	}

	for i := range results {
		fmt.Printf("%s\n", results[i].Name)
	}

	cfg.Next = newCfg.Next
	cfg.Previous = newCfg.Previous

	return nil
}
