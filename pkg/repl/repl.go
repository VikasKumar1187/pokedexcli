// pkg/repl/repl.go
package repl

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Pokedex represents the Pokedex application.
type Repl struct {
	commands map[string]func()
}

// NewPokedex creates a new instance of Pokedex.
func NewRepl() *Repl {
	return &Repl{
		commands: map[string]func(){
			"help": helpCommand,
			"exit": exitCommand,
		},
	}
}

// Execute is the main entry point for the Pokedex CLI.
func Execute() {
	repl := NewRepl()
	repl.StartREPL()
}

// StartREPL starts the Pokedex REPL.
func (p *Repl) StartREPL() {
	fmt.Println("Welcome to the Pokedex REPL!")
	fmt.Println("Type 'help' for available commands and usage.")

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("pokedex > ")
		scanner.Scan()
		userInput := strings.TrimSpace(scanner.Text())

		if command, exists := p.commands[userInput]; exists {
			command()
		} else {
			fmt.Printf("Unknown command: %s. Type 'help' for usage.\n", userInput)
		}
	}
}
