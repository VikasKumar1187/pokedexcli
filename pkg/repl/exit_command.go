// pkg/repl/exit.go
package repl

import (
	"fmt"
	"os"
)

// exitCommand exits the Pokedex CLI.
func exitCommand() {
	fmt.Println("Goodbye! Exiting Pokedex.")
	os.Exit(0)
}
