package brainfuck

import "bf/internal/input"

func NullInput() input.Input {
	return input.NullInput()
}

func StringInput(data string) input.Input {
	return input.StringInput(data, true)
}

func StringInputWithHeader(data string) input.Input {
	return input.StringInput(data, false)
}

func FileInput(filename string) input.Input {
	return input.FileInput(filename)
}
