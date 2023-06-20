package looper

import (
	"bf/internal/dialect"
	"errors"
)

type LoopStartFinder func(code string, position int) (int, error)
type LoopEndFinder func(code string, position int) (int, error)

// FindStartOfLoopDelta returns the delta value, as a negative integer, indicating the distance to the start loop
// command that matches the end loop command at the position indicated. It is not necessary, although possible, to
// provide the entire code to this method. However, it is necessary to provide at least the code ending at the end
// of loop and beginning either all the way at the start of the program, or at least far enough back to locate the
// start of the loop. Since the purpose of this method is to find the start of the loop, it is unlikely that
// callers will know how much code to include in order to ensure location of the start of the loop, so this method
// will usually be called with code that extends to the start of the program. There is no benefit in performance
// if a subset of code is provided.
func FindStartOfLoopDelta(code string, position int) (int, error) {
	err := validateArgs(code, position)
	if err != nil {
		return 0, err
	}
	done := false
	delta := 0
	loopCount := 1
	for !done {
		delta--
		if position+delta < 0 {
			return 0, errors.New("loop has no start")
		}
		token := string([]rune(code)[position+delta])
		if token == dialect.StartLoop {
			loopCount--
			if loopCount == 0 {
				done = true
				continue
			}
		}
		if token == dialect.EndLoop {
			loopCount++
		}
	}
	return delta, nil
}

// FindEndOfLoopDelta returns the delta value, as a positive integer, indicating the distance to the end loop
// command that matches the start loop command at the position indicated. It is not necessary, although possible,
// to provide the entire code to this method. However, it is necessary to provide at least the code beginning at
// the start of loop and continuing either all the way to the end of the program, or at least far enough to locate
// the end of the loop. Since the purpose of this method is to find the end of the loop, it is unlikely that
// callers will know how much code to include in order to ensure location of the end of the loop, so this method
// will usually be called with code that extends to the end of the program. There is no benefit in performance if a
// subset of code is provided.
func FindEndOfLoopDelta(code string, position int) (int, error) {
	err := validateArgs(code, position)
	if err != nil {
		return 0, err
	}
	done := false
	delta := 0
	loopCount := 1
	for !done {
		delta++
		if position+delta == len(code) {
			return 0, errors.New("loop has no end")
		}
		token := string([]rune(code)[position+delta])
		if token == dialect.EndLoop {
			loopCount--
			if loopCount == 0 {
				done = true
				continue
			}
		}
		if token == dialect.StartLoop {
			loopCount++
		}
	}
	return delta, nil
}

func validateArgs(code string, position int) error {
	if code == "" {
		return errors.New("code is missing")
	}
	if position < 0 {
		return errors.New("position is before start of code")
	}
	if position >= len(code) {
		return errors.New("position is after end of code")
	}
	return nil
}
