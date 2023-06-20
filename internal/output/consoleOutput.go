package output

import (
	"fmt"
)

type consoleOutput struct{}

func (c consoleOutput) Write(value int) error {
	fmt.Print(string(rune(value)))
	return nil
}

// ConsoleOutput returns a simple Output implementation that outputs to the system console.
func ConsoleOutput() Output {
	return consoleOutput{}
}
