// pkg/repl/pokedex.go
package pokedex

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/VikasKumar1187/pokedexcli/internal/pokeapi"
)

type Config struct {
	PokeApiClient    *pokeapi.Client
	NextLocationsURL *string
	PrevLocationsURL *string
}

// Pokedex represents the Pokedex application.
type Repl struct {
	commands map[string]func(*Config)
}

// NewPokedex creates a new instance of Pokedex.
func NewRepl() *Repl {
	return &Repl{
		commands: map[string]func(*Config){
			"help": helpCommand,
			"exit": exitCommand,
			"map":  mapCommand,
			"mapb": mapBackCommand,
		},
	}
}

// Execute is the main entry point for the Pokedex CLI.
func StartPokedex(cfg *Config) {
	repl := NewRepl()
	repl.StartREPL(cfg)
}

// StartREPL starts the Pokedex REPL.
func (rpl *Repl) StartREPL(cfg *Config) {
	fmt.Println("Welcome to the Pokedex REPL!")
	fmt.Println("Type 'help' for available commands and usage.")

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("pokedex > ")
		scanner.Scan()
		userInput := strings.TrimSpace(scanner.Text())

		if command, exists := rpl.commands[userInput]; exists {
			command(cfg)
		} else {
			fmt.Printf("Unknown command: %s. Type 'help' for usage.\n", userInput)
		}
	}
}
