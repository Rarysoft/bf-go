package memory

type signed16BitMemory map[int]int

func (signed16BitMemory) MinAddress() int {
	return 0x0000
}

func (signed16BitMemory) MaxAddress() int {
	return 0xFFFF
}

func (signed16BitMemory) MinValue() int {
	return -0x8000
}

func (signed16BitMemory) MaxValue() int {
	return 0x7FFF
}

func (s signed16BitMemory) Read(address int) (int, error) {
	return read(address, s.MinAddress(), s.MaxAddress(), s.readAddress)
}

func (s signed16BitMemory) Write(address int, value int) error {
	return write(address, value, s.MinAddress(), s.MaxAddress(), s.MinValue(), s.MaxValue(), s.writeAddress)
}

func (s signed16BitMemory) readAddress(address int) int {
	return s[address]
}

func (s signed16BitMemory) writeAddress(address int, value int) {
	s[address] = value
}

func Signed16BitMemory() Memory {
	return signed16BitMemory{}
}
