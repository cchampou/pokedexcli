package commands

import "os"

func ExitExecutor() error {
	os.Exit(0)
	return nil
}
