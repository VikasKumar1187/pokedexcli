// pkg/repl/exit.go
package pokedex

import (
	"fmt"
	"os"
)

// exitCommand exits the Pokedex CLI.
func exitCommand(*Config) error {
	fmt.Println("Goodbye! Exiting Pokedex.")
	os.Exit(0)
	return nil
}
