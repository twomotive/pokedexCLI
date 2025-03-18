package commands

import (
	"fmt"

	"github.com/twomotive/pokedex/internal/api"
	"github.com/twomotive/pokedex/internal/config"
)

// MapBack displays previous location areas
func CommandMapBack(cfg *config.AppConfig, client *api.Client) error {
	if cfg.Pagination.Previous == nil {
		fmt.Println("No previous location areas.")
		return nil
	}

	url := *cfg.Pagination.Previous
	resp, err := client.GetLocationAreas(url)
	if err != nil {
		return fmt.Errorf("cannot mapb: %v", err)
	}

	for _, area := range resp.Results {
		fmt.Println(area.Name)
	}

	// Update pagination
	cfg.Pagination.Next = resp.Next
	cfg.Pagination.Previous = resp.Previous
	return nil
}
