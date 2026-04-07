package interpreter

import (
	"time"

	"github.com/AlanRostem/mu-8/internal/num"
	"github.com/AlanRostem/mu-8/internal/processor"
	"github.com/AlanRostem/mu-8/internal/system"
)

const ProgramSize = system.MemorySize - system.ProgramOffset

type Interpreter struct {
	system    *system.System
	processor *processor.Processor
}

func New() *Interpreter {
	s := system.New()
	return &Interpreter{
		system:    s,
		processor: processor.New(s),
	}
}

func (in *Interpreter) SetKey(key uint8, state bool) {
	in.system.Keys[key] = state
}

func (in *Interpreter) Load(program [ProgramSize]num.Byte) {
	in.system.Registers.Clear()
	for i, value := range program {
		addr := num.NewUint12(system.ProgramOffset + i)
		in.system.Memory.Write(addr, value)
	}
}

// Run starts the processor cycle "loop". This method is blocking
func (in *Interpreter) Run() {
	for {
		in.processor.Cycle()
		// simulate chip8 "clock speed"
		time.Sleep(time.Second / processor.Frequency)
	}
}
