package commands

import "os"

func ExitExecutor(config *Config) error {
	os.Exit(0)
	return nil
}
