package main

import (
	"bufio"
	commands "github.com/cchampou/pokedexcli/commands"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	commands.PrintPrompt()
	for scanner.Scan() {
		textInput := scanner.Text()
		commandMap := commands.GetCommands()
		selectedCommand, ok := commandMap[textInput]
		if ok {
			err := selectedCommand.Callback()
			if err != nil {
				return
			}
		}
		commands.PrintPrompt()
	}
}
