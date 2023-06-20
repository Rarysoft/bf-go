package memory

type unsigned8BitMemory map[int]int

func (unsigned8BitMemory) MinAddress() int {
	return 0x0000
}

func (unsigned8BitMemory) MaxAddress() int {
	return 0x752F
}

func (unsigned8BitMemory) MinValue() int {
	return 0x00
}

func (unsigned8BitMemory) MaxValue() int {
	return 0xFF
}

func (u unsigned8BitMemory) Read(address int) (int, error) {
	return read(address, u.MinAddress(), u.MaxAddress(), u.readAddress)
}

func (u unsigned8BitMemory) Write(address int, value int) error {
	return write(address, value, u.MinAddress(), u.MaxAddress(), u.MinValue(), u.MaxValue(), u.writeAddress)
}

func (u unsigned8BitMemory) readAddress(address int) int {
	return u[address]
}

func (u unsigned8BitMemory) writeAddress(address int, value int) {
	u[address] = value
}

func Unsigned8BitMemory() Memory {
	return unsigned8BitMemory{}
}
