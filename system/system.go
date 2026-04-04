package system

type System struct {
	Registers   *RegisterFile
	Memory      *MemoryBank
	FrameBuffer [32][64]bool
}

func New() *System {
	return &System{
		Registers:   newRegisterFile(),
		Memory:      newMemoryBank(),
		FrameBuffer: [32][64]bool{},
	}
}
