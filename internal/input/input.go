package input

// Input is an interface for an Executor implementation to read from when it encounters code that indicates to input a
// value.
type Input interface {
	// Read reads from the input.
	Read() (int, error)
}
