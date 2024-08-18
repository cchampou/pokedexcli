package commands

import "fmt"

func HelpExecutor() error {
	println("")
	println("Welcome to the Pokedex!")
	println("Usage:")
	println("")
	commands := GetCommands()
	for _, command := range commands {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}
	println("")
	return nil
}
