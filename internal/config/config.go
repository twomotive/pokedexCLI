package config

import "github.com/twomotive/pokedex/internal/api"

// AppConfig represents the application's configuration
type AppConfig struct {
	Pagination *api.PaginationConfig
	Pokedex    map[string]api.Pokemon
}

// NewAppConfig creates a new application configuration
func NewAppConfig() *AppConfig {
	return &AppConfig{
		Pagination: &api.PaginationConfig{},
		Pokedex:    make(map[string]api.Pokemon),
	}
}
