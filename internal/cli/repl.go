package cli

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/twomotive/pokedex/internal/api"
	"github.com/twomotive/pokedex/internal/cli/commands"
	"github.com/twomotive/pokedex/internal/config"
)

func StartREPL() {
	reader := bufio.NewScanner(os.Stdin)
	client := api.NewClient()
	appConfig := config.NewAppConfig()

	fmt.Println("Welcome to the Pokedex! Type 'help' for commands.")

	for {
		fmt.Print("Pokedex > ")
		reader.Scan()

		words := cleanInput(reader.Text())
		if len(words) == 0 {
			continue
		}

		commandName := words[0]
		command, exists := commands.GetCommands()[commandName]

		if exists {
			// Pass the rest of the words as arguments to the command callback
			args := []string{}
			if len(words) > 1 {
				args = words[1:]
			}
			err := command.Callback(appConfig, client, args)
			if err != nil {
				fmt.Printf("Error: %v\n", err)
			}
		} else {
			fmt.Println("Unknown command. Type 'help' for available commands.")
		}
	}
}

// cleanInput sanitizes user input
func cleanInput(text string) []string {
	output := strings.ToLower(text)
	return strings.Fields(output)
}
