package main

import (
	"bufio"
	"fmt"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	printPrompt()
	for scanner.Scan() {
		textInput := scanner.Text()
		commands := getCommands()
		selectedCommand, ok := commands[textInput]
		if ok {
			err := selectedCommand.callback()
			if err != nil {
				return
			}
		}
		printPrompt()
	}
}

func printPrompt() {
	print("Pokedex > ")
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    helpExecutor,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    exitExecutor,
		},
	}
}

func helpExecutor() error {
	println("")
	println("Welcome to the Pokedex!")
	println("Usage:")
	println("")
	commands := getCommands()
	for _, command := range commands {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}
	println("")
	return nil
}

func exitExecutor() error {
	os.Exit(0)
	return nil
}
