package commands

import (
	"github.com/twomotive/pokedex/internal/api"
	"github.com/twomotive/pokedex/internal/config"
)

// Command represents a CLI command
type Command struct {
	Name        string
	Description string
	Callback    func(*config.AppConfig, *api.Client, []string) error
}

// GetCommands returns all registered commands
func GetCommands() map[string]Command {
	return map[string]Command{
		"mapb": {
			Name:        "mapb",
			Description: "Displays previous 20 location areas in the Pokemon world",
			Callback:    CommandMapBack,
		},
		"map": {
			Name:        "map",
			Description: "Displays the names of 20 location areas in the Pokemon world",
			Callback:    CommandMap,
		},
		"explore": {
			Name:        "explore",
			Description: "Explore a location area in the Pokemon world",
			Callback:    CommandExplore,
		},
		"catch": {
			Name:        "catch",
			Description: "Catch a pokemon",
			Callback:    CommandCatch,
		},
		"pokedex": {
			Name:        "pokedex",
			Description: "View your caught Pokemon",
			Callback:    CommandPokedex,
		},
		"inspect": {
			Name:        "inspect",
			Description: "Inspect a caught pokemon",
			Callback:    CommandInspect,
		},
		"help": {
			Name:        "help",
			Description: "Displays a help message",
			Callback:    CommandHelp,
		},
		"exit": {
			Name:        "exit",
			Description: "Exit the Pokedex",
			Callback:    CommandExit,
		},
	}
}
