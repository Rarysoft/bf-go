package memory

type unsigned16BitMemory map[int]int

func (unsigned16BitMemory) MinAddress() int {
	return 0x0000
}

func (unsigned16BitMemory) MaxAddress() int {
	return 0xFFFF
}

func (unsigned16BitMemory) MinValue() int {
	return 0x0000
}

func (unsigned16BitMemory) MaxValue() int {
	return 0xFFFF
}

func (u unsigned16BitMemory) Read(address int) (int, error) {
	return read(address, u.MinAddress(), u.MaxAddress(), u.readAddress)
}

func (u unsigned16BitMemory) Write(address int, value int) error {
	return write(address, value, u.MinAddress(), u.MaxAddress(), u.MinValue(), u.MaxValue(), u.writeAddress)
}

func (u unsigned16BitMemory) readAddress(address int) int {
	return u[address]
}

func (u unsigned16BitMemory) writeAddress(address int, value int) {
	u[address] = value
}

func Unsigned16BitMemory() Memory {
	return unsigned16BitMemory{}
}
