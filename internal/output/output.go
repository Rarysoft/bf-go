package output

// Output is an interface for an {@link Executor} implementation to write to when it encounters code that indicates to
// output a value.
type Output interface {
	// Write writes the provided value to the output.
	Write(value int) error
}
