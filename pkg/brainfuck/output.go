package brainfuck

import "bf/internal/output"

func NullOutput() output.Output {
	return output.NullOutput()
}

func ConsoleOutput() output.Output {
	return output.ConsoleOutput()
}
