package system

type System struct {
	Registers   *RegisterFile
	Memmory     *MemoryBank
	FrameBuffer [32][64]bool
}

func New() *System {
	return &System{
		Registers:   newRegisterFile(),
		Memmory:     newMemoryBank(),
		FrameBuffer: [32][64]bool{},
	}
}
