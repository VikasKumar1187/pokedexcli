package main

import (
	"time"

	"github.com/VikasKumar1187/pokedexcli/internal/pokeapi"
	"github.com/VikasKumar1187/pokedexcli/pkg/pokedex"
)

func main() {
	pokeClient := pokeapi.NewClient(5 * time.Second)
	cfg := &pokedex.Config{
		PokeApiClient: pokeClient,
	}

	pokedex.StartPokedex(cfg)
}
