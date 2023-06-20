package brainfuck

import (
	"bf/internal/dialect"
	"bf/internal/executor"
	"bf/internal/input"
	"bf/internal/looper"
	"bf/internal/memory"
	"bf/internal/output"
)

// BF is the core of the BF library, the brainfuck interpreter. This is the type to use in order to execute brainfuck
// code. It requires an implementation of the Executor interface and an implementation of the two Looper functions.
type BF struct {
	executor        executor.Executor
	loopStartFinder looper.LoopStartFinder
	loopEndFinder   looper.LoopEndFinder
}

// Run runs a brainfuck program. If any coding errors are encountered while executing the code, an error will be
// returned. This method will attempt to execute the code, whether or not any actual brainfuck code is found. An error
// will only be returned if invalid brainfuck code is found, such as incorrect matching of [ and ] commands.
func (b BF) Run(code string) error {
	position := 0
	for position < len(code) {
		delta, err := b.perform(code, position)
		if err != nil {
			return err
		}
		position += delta
	}
	return nil
}

func (b BF) perform(code string, position int) (int, error) {
	token := string([]rune(code)[position])
	switch token {
	case dialect.Increment:
		err := b.executor.PerformIncrement()
		if err != nil {
			return 0, err
		}
		break
	case dialect.Decrement:
		err := b.executor.PerformDecrement()
		if err != nil {
			return 0, err
		}
		break
	case dialect.IncrementPointer:
		b.executor.PerformIncrementPointer()
		break
	case dialect.DecrementPointer:
		b.executor.PerformDecrementPointer()
		break
	case dialect.StartLoop:
		jumpToEnd, err := b.executor.PerformStartLoop()
		if err != nil {
			return 0, err
		}
		if jumpToEnd {
			return b.loopEndFinder(code, position)
		}
		break
	case dialect.EndLoop:
		jumpToStart, err := b.executor.PerformEndLoop()
		if err != nil {
			return 0, err
		}
		if jumpToStart {
			return b.loopStartFinder(code, position)
		}
		break
	case dialect.Input:
		err := b.executor.PerformInput()
		if err != nil {
			return 0, err
		}
		break
	case dialect.Output:
		err := b.executor.PerformOutput()
		if err != nil {
			return 0, err
		}
		break
	default:
		break
	}
	return 1, nil
}

// Default returns an instance of the interpreter that uses a BFExecutor and the default looper functions. The
// BFExecutor instance will use the NullInput, ConsoleOutput, and Unsigned8BitMemory implementations.
func Default() BF {
	return BF{
		executor:        executor.Default(),
		loopStartFinder: looper.FindStartOfLoopDelta,
		loopEndFinder:   looper.FindEndOfLoopDelta,
	}
}

// New returns an instance of the interpreter that uses the provided Executor implementation and looper functions.
func New(executor executor.Executor, loopStartFinder looper.LoopStartFinder, loopEndFinder looper.LoopEndFinder) BF {
	return BF{
		executor:        executor,
		loopStartFinder: loopStartFinder,
		loopEndFinder:   loopEndFinder,
	}
}

// Custom returns an instance of the interpreter that uses a BFExecutor and the default looper functions. The
// BFExecutor instance will use the provided Input, Output, and Memory implementations.
func Custom(input input.Input, output output.Output, memory memory.Memory) BF {
	return BF{
		executor:        executor.New(input, output, memory),
		loopStartFinder: looper.FindStartOfLoopDelta,
		loopEndFinder:   looper.FindEndOfLoopDelta,
	}
}
