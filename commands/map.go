package commands

import "github.com/cchampou/pokedexcli/remote"

func MapExecutor(config *Config) error {
	strings, err := remote.FetchLocationArea(config.Cursor)
	if err != nil {
		return err
	}
	for _, area := range strings {
		println(area)
	}
	return nil
}
