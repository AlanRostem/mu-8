package system

import (
	"github.com/AlanRostem/mu-8/cpu"
	"github.com/AlanRostem/mu-8/memory"
)

type System struct {
	Registers cpu.RegisterFile
	Memmory   memory.Memory
}
