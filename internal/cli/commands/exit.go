package commands

import (
	"fmt"
	"os"

	"github.com/twomotive/pokedex/internal/api"
	"github.com/twomotive/pokedex/internal/config"
)

// Exit exits the application
func CommandExit(cfg *config.AppConfig, client *api.Client) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
