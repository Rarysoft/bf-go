package brainfuck

import (
	"bf/internal/executor"
	"bf/internal/input"
	"bf/internal/memory"
	"bf/internal/output"
)

func DefaultExecutor() executor.Executor {
	return executor.Default()
}

func NewExecutor(input input.Input, output output.Output, memory memory.Memory) executor.Executor {
	return executor.New(input, output, memory)
}
