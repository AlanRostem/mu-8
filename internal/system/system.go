package system

const (
	DisplayWidth  = 64
	DisplayHeight = 32
	KeyCount      = 0x10
)

type System struct {
	Registers   *RegisterFile
	Memory      *MemoryBank
	Stack       *Stack
	FrameBuffer [DisplayHeight][DisplayWidth]bool
	Keys        [KeyCount]bool
}

func New() *System {
	return &System{
		Registers:   newRegisterFile(),
		Memory:      newMemoryBank(),
		Stack:       newStack(),
		FrameBuffer: [DisplayHeight][DisplayWidth]bool{},
	}
}
