package commands

import (
	"fmt"

	"github.com/twomotive/pokedex/internal/api"
	"github.com/twomotive/pokedex/internal/config"
)

// Map displays location areas
func CommandMap(cfg *config.AppConfig, client *api.Client) error {
	url := ""
	if cfg.Pagination.Next != nil {
		url = *cfg.Pagination.Next
	}

	resp, err := client.GetLocationAreas(url)
	if err != nil {
		return fmt.Errorf("cannot map: %v", err)
	}

	for _, area := range resp.Results {
		fmt.Println(area.Name)
	}

	// Update pagination
	cfg.Pagination.Next = resp.Next
	cfg.Pagination.Previous = resp.Previous
	return nil
}
