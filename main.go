package main

import (
	"bufio"
	commands "github.com/cchampou/pokedexcli/commands"
	"log"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	config := commands.Config{
		Cursor: 0,
	}
	commands.PrintPrompt()
	for scanner.Scan() {
		textInput := scanner.Text()
		commandMap := commands.GetCommands()
		selectedCommand, ok := commandMap[textInput]
		if ok {
			err := selectedCommand.Callback(&config)
			if err != nil {
				log.Fatalf("%v", err)
			}
		}
		commands.PrintPrompt()
	}
}
