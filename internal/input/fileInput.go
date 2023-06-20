package input

import (
	"os"
)

// FileInput returns an implementation of Input that reads input from the specified file. An error will be returned if
// the file does not exist. If the second argument is true, the header indicating the size of the file will not be
// included, and instead the first read will return the first byte in the file.
//
// Internally, this is not truly a separate implementation. This function simply reads the file into a string and then
// returns a stringInput.
func FileInput(filename string, skipHeader bool) (Input, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &stringInput{
		data:       string(data),
		headerRead: skipHeader,
		index:      0}, nil
}
