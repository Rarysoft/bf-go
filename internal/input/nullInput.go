package input

type nullInput struct{}

func (nullInput) Read() (int, error) {
	return 0, nil
}

// NullInput returns a simple Input implementation that accepts no input from anywhere and simply returns 0 to any read
// attempt.
func NullInput() Input {
	return nullInput{}
}
