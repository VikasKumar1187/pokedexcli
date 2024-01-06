// pkg/repl/exit.go
package pokedex

import (
	"context"
	"errors"
	"fmt"
)

// exitCommand exits the Pokedex CLI.
func mapCommand(cfg *Config) error {

	shallowLocations, err := cfg.PokeApiClient.ListLocations(context.Background(), cfg.NextLocationsURL)

	if err != nil {
		return err
	}

	cfg.NextLocationsURL = shallowLocations.Next
	cfg.PrevLocationsURL = shallowLocations.Previous

	for _, loc := range shallowLocations.Results {
		fmt.Println(loc.Name)
	}

	return nil
}

// exitCommand exits the Pokedex CLI.
func mapBackCommand(cfg *Config) error {

	if cfg.PrevLocationsURL == nil {
		return errors.New("you're on the first page")
	}

	shallowLocations, err := cfg.PokeApiClient.ListLocations(context.Background(), cfg.PrevLocationsURL)
	if err != nil {
		return err
	}

	cfg.NextLocationsURL = shallowLocations.Next
	cfg.PrevLocationsURL = shallowLocations.Previous

	for _, loc := range shallowLocations.Results {
		fmt.Println(loc.Name)
	}

	return nil
}
