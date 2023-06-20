package output

type nullOutput struct{}

func (nullOutput) Write(value int) error {
	return nil
}

// NullOutput returns a simple Output implementation that does nothing with values written.
func NullOutput() Output {
	return nullOutput{}
}
