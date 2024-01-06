// pkg/repl/exit.go
package pokedex

import (
	"fmt"
	"os"
)

// exitCommand exits the Pokedex CLI.
func exitCommand(*Config) {
	fmt.Println("Goodbye! Exiting Pokedex.")
	os.Exit(0)
}
