package system

const (
	DisplayWidth  = 64
	DisplayHeight = 32
)

type System struct {
	Registers   *RegisterFile
	Memory      *MemoryBank
	FrameBuffer [DisplayHeight][DisplayWidth]bool
}

func New() *System {
	return &System{
		Registers:   newRegisterFile(),
		Memory:      newMemoryBank(),
		FrameBuffer: [DisplayHeight][DisplayWidth]bool{},
	}
}
