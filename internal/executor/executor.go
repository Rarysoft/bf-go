package executor

import (
	"bf/internal/input"
	"bf/internal/memory"
	"bf/internal/output"
)

type Executor interface {
	// PerformIncrement performs the increment operation.
	PerformIncrement() error

	// PerformDecrement performs the decrement operation.
	PerformDecrement() error

	// PerformIncrementPointer performs the increment pointer operation.
	PerformIncrementPointer()

	// PerformDecrementPointer performs the decrement pointer operation.
	PerformDecrementPointer()

	// PerformStartLoop performs the start loop operation.
	PerformStartLoop() (bool, error)

	// PerformEndLoop performs the end loop operation.
	PerformEndLoop() (bool, error)

	// PerformInput performs the input operation.
	PerformInput() error

	// PerformOutput performs the output operation.
	PerformOutput() error
}

type bfExecutor struct {
	input   input.Input
	output  output.Output
	memory  memory.Memory
	pointer int
}

// GetPointer gets the value of the pointer. This method is not normally used when executing code, and is not used at
// all by the BF interpreter. It can be used for testing or to otherwise gain insight into the internal workings of the
// executor.
func (b *bfExecutor) GetPointer() int {
	return b.pointer
}

// SetPointer sets the value of the pointer. This method is not normally used when executing code, and is not used at
// all by the BF interpreter. It can be used for testing or to manipulate the internal workings of the executor. Note
// that doing so is dangerous, and may alter the code execution in ways that violate the brainfuck standard, and may
// result in an error.
func (b *bfExecutor) SetPointer(p int) {
	b.pointer = p
}

func (b *bfExecutor) PerformIncrement() error {
	current, err := b.memory.Read(b.pointer)
	if err != nil {
		return err
	}
	if current == b.memory.MaxValue() {
		return b.memory.Write(b.pointer, b.memory.MinValue())
	}
	return b.memory.Write(b.pointer, current+1)
}

func (b *bfExecutor) PerformDecrement() error {
	current, err := b.memory.Read(b.pointer)
	if err != nil {
		return err
	}
	if current == b.memory.MinValue() {
		return b.memory.Write(b.pointer, b.memory.MaxValue())
	}
	return b.memory.Write(b.pointer, current-1)
}

func (b *bfExecutor) PerformIncrementPointer() {
	if b.pointer == b.memory.MaxAddress() {
		b.pointer = b.memory.MinAddress()
		return
	}
	b.pointer++
}

func (b *bfExecutor) PerformDecrementPointer() {
	if b.pointer == b.memory.MinAddress() {
		b.pointer = b.memory.MaxAddress()
		return
	}
	b.pointer--
}

func (b *bfExecutor) PerformStartLoop() (bool, error) {
	value, err := b.memory.Read(b.pointer)
	if err != nil {
		return false, err
	}
	return value == 0, nil
}

func (b *bfExecutor) PerformEndLoop() (bool, error) {
	value, err := b.memory.Read(b.pointer)
	if err != nil {
		return false, err
	}
	return value != 0, nil
}

func (b bfExecutor) PerformInput() error {
	value, err := b.input.Read()
	if err != nil {
		return err
	}
	return b.memory.Write(b.pointer, value)
}

func (b *bfExecutor) PerformOutput() error {
	value, err := b.memory.Read(b.pointer)
	if err != nil {
		return err
	}
	return b.output.Write(value)
}

func Default() Executor {
	return &bfExecutor{
		input:   input.NullInput(),
		output:  output.ConsoleOutput(),
		memory:  memory.Unsigned8BitMemory(),
		pointer: 0,
	}
}

func New(input input.Input, output output.Output, memory memory.Memory) Executor {
	return &bfExecutor{
		input:   input,
		output:  output,
		memory:  memory,
		pointer: 0,
	}
}
