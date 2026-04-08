package system

import "github.com/AlanRostem/mu-8/internal/num"

const (
	DisplayWidth  = 64
	DisplayHeight = 32
	KeyCount      = 0x10
	NoHaltKey     = KeyCount
)

type System struct {
	Registers          *RegisterFile
	Memory             *MemoryBank
	Stack              *Stack
	FrameBuffer        [DisplayHeight][DisplayWidth]bool
	Keys               [KeyCount]bool
	HaltingKeyRegister num.DByte
}

func New() *System {
	return &System{
		Registers:          newRegisterFile(),
		Memory:             newMemoryBank(),
		Stack:              newStack(),
		FrameBuffer:        [DisplayHeight][DisplayWidth]bool{},
		HaltingKeyRegister: NoHaltKey,
	}
}
