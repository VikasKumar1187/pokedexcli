// pkg/repl/help.go
package pokedex

import "fmt"

// helpCommand displays help information.
func helpCommand(*Config) error {
	fmt.Println("Available commands:")
	fmt.Println("  help: prints this help message")
	fmt.Println("  exit: exits the Pokedex")
	return nil
}
