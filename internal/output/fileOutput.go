package output

// FileOutput returns an implementation that writes output to the specified file. If the second argument is true, data
// will be appended to the file if it already exists. If the second argument is false, the file will be overwritten if
// it already exists. In either case, if the file does not exist, a new file will be created.
func FileOutput(filename string, append bool) (Output, error) {
	// TODO: to be implemented
	return nil, nil
}
