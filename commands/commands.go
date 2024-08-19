package commands

type cliCommand struct {
	name        string
	description string
	Callback    func(config *Config) error
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
		"map": {
			name:        "map",
			description: "Displays the names of 20 location areas in the Pokemon world (next page)",
			Callback:    MapExecutor,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the names of 20 location areas in the Pokemon world (prev page)",
			Callback:    MapBExecutor,
		},
	}
}
