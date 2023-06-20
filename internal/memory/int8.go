package memory

type signed8BitMemory map[int]int

func (signed8BitMemory) MinAddress() int {
	return 0x0000
}

func (signed8BitMemory) MaxAddress() int {
	return 0x752F
}

func (signed8BitMemory) MinValue() int {
	return -0x80
}

func (signed8BitMemory) MaxValue() int {
	return 0x7F
}

func (s signed8BitMemory) Read(address int) (int, error) {
	return read(address, s.MinAddress(), s.MaxAddress(), s.readAddress)
}

func (s signed8BitMemory) Write(address int, value int) error {
	return write(address, value, s.MinAddress(), s.MaxAddress(), s.MinValue(), s.MaxValue(), s.writeAddress)
}

func (s signed8BitMemory) readAddress(address int) int {
	return s[address]
}

func (s signed8BitMemory) writeAddress(address int, value int) {
	s[address] = value
}

func Signed8BitMemory() Memory {
	return signed8BitMemory{}
}
