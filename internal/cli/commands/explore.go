package commands

import (
	"fmt"

	"github.com/twomotive/pokedex/internal/api"
	"github.com/twomotive/pokedex/internal/config"
)

func CommandExplore(cfg *config.AppConfig, client *api.Client, args []string) error {
	url := ""

	if len(args) == 0 {
		return fmt.Errorf("missing location area name")
	}

	areaName := args[0]

	fmt.Printf("Exploring %s...\n", areaName)

	resp, err := client.GetPokemons(url, areaName)
	if err != nil {
		return fmt.Errorf("cannot explore: %v", err)
	}

	fmt.Printf("Found Pokemon:\n")
	for _, encounter := range resp.PokemonEncounters {
		fmt.Printf(" - %s\n", encounter.Pokemon.Name)
	}

	return nil

}
