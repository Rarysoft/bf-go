package config

import (
	"errors"
	"os"
)

type Config struct {
	SourceFilename string
}

func FromArgs() (Config, error) {
	if len(os.Args) < 2 {
		return Config{}, errors.New("missing arguments")
	}
	return Config{
		SourceFilename: os.Args[1],
	}, nil
}
