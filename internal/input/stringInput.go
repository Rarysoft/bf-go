package input

import (
	"errors"
)

type stringInput struct {
	data       string
	headerRead bool
	index      int
}

func (s *stringInput) Read() (int, error) {
	if !s.headerRead {
		s.headerRead = true
		return len(s.data), nil
	}
	if s.index >= len(s.data) {
		return 0, errors.New("read past end of data")
	}
	value := int([]rune(s.data)[s.index])
	s.index++
	return value, nil
}

// StringInput returns an implementation of Input that reads from a predefined String. Instances can be created with or
// without header data. With header data, the first read will return the size of the String. This size indicates the
// number of additional reads that can safely be executed before the end of the String is reached. Without a header,
// the String length must be known.
func StringInput(data string, skipHeader bool) Input {
	return &stringInput{
		data:       data,
		headerRead: skipHeader,
		index:      0,
	}
}
