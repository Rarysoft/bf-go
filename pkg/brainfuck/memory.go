package brainfuck

import "bf/internal/memory"

func Signed8BitMemory() memory.Memory {
	return memory.Signed8BitMemory()
}

func Signed16BitMemory() memory.Memory {
	return memory.Signed16BitMemory()
}

func Unsigned8BitMemory() memory.Memory {
	return memory.Unsigned8BitMemory()
}

func Unsigned16BitMemory() memory.Memory {
	return memory.Unsigned16BitMemory()
}
