package system

const (
	DisplayWidth  = 64
	DisplayHeight = 32
)

type System struct {
	Registers   *RegisterFile
	Memory      *MemoryBank
	Stack       *Stack
	FrameBuffer [DisplayHeight][DisplayWidth]bool
}

func New() *System {
	return &System{
		Registers:   newRegisterFile(),
		Memory:      newMemoryBank(),
		Stack:       newStack(),
		FrameBuffer: [DisplayHeight][DisplayWidth]bool{},
	}
}
