package ram

const (
	// FrameSize defines the number of words in a frame
	FrameSize = 32

	// MemorySize determines how many words of memory are in memory
	MemorySize = 1024
)

type (
	// Address describes a physical address to be provided to RAM
	Address uint32

	// Word describes the storage of a physial space in RAM
	Word uint32
)

// GetData gets the data at the given address
func GetData(a Address) Word {
	return physicalMemory[a]
}

// SetData sets the data at the given address
func SetData(a Address, d Word) {
	physicalMemory[a] = d
}
