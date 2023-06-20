package memory

import "errors"

// Memory is an interface for an Executor implementation to use for internal program memory. This memory stores, in
// some unspecified internal format, the data values stored in cells pointed to by the program pointer.
type Memory interface {
	// MinAddress returns the lowest address value. The program pointer will initially point to this address.
	MinAddress() int

	// MaxAddress returns the highest address value.
	MaxAddress() int

	// MinValue returns the lowest value that can be stored in a memory cell.
	MinValue() int

	// MaxValue returns the highest value that can be stored in a memory cell.
	MaxValue() int

	// Read returns the value stored in the cell at the provided address, or pointer index. The address must be within
	// the range specified by the minAddress and maxAddress values. By convention, if an address outside of the valid
	// range is specified, the implementation should return an error. Also by convention, addresses not previously
	// initialized should return 0.
	Read(address int) (int, error)

	// Write stores the specified value in the cell at the provided address, or pointer index. The address must be
	// within the range specified by the minAddress and maxAddress values. The value must be within the range specified
	// by the minValue and maxValue values. By convention, if an address outside the valid range is specified, the
	// implementation should return an error. Also by convention, if a value outside the valid range is specified, the
	// implementation should return an error.
	Write(address int, value int) error
}

type readFromMemory func(int) int
type writeToMemory func(int, int)

func read(address int, minAddress int, maxAddress int, memoryReader readFromMemory) (int, error) {
	if address > maxAddress || address < minAddress {
		return 0, errors.New("address out of range")
	}
	return memoryReader(address), nil
}

func write(address int, value int, minAddress int, maxAddress int, minValue int, maxValue int, memoryWriter writeToMemory) error {
	if address > maxAddress || address < minAddress {
		return errors.New("address out of range")
	}
	if value > maxValue || value < minValue {
		return errors.New("value out of range")
	}
	memoryWriter(address, value)
	return nil
}
