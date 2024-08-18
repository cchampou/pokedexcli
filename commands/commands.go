package commands

type cliCommand struct {
	name        string
	description string
	Callback    func() error
}

func GetCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			Callback:    HelpExecutor,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			Callback:    ExitExecutor,
		},
	}
}
