package instruction

import (
	"time"

	"github.com/AlanRostem/mu-8/decode"
	"github.com/AlanRostem/mu-8/num"
	"github.com/AlanRostem/mu-8/system"
)

const ProgramSize = system.MemorySize - system.ProgramOffset
const CpuFrequency = 500

// OpcodeExit is a special opcode used by the Executor to
// exit the program for testing purposes. This will be
// removed in prod.
const OpcodeExit = num.DByte(0xFFFF)

type Executor struct {
	System *system.System
}

func NewExecutor(system *system.System) *Executor {
	return &Executor{
		System: system,
	}
}

func (e *Executor) ExecOpcode(opcode num.DByte) {
	info := decode.Decode(opcode)
	table := tableAll[info.Class]
	if table.IsSingle() {
		table.Single()(info.Args, e.System)
		return
	}
	identity := info.Idenitity
	inst := table[identity]
	inst(info.Args, e.System)
}

func (e *Executor) LoadProgram(program [ProgramSize]num.Byte) {
	for i, value := range program {
		addr := num.NewUint12(system.ProgramOffset + i)
		e.System.Memory.Write(addr, value)
	}
	e.System.Registers.Clear()
}

func (e *Executor) ExecProgram() {
	for {
		addr := num.NewUint12(e.System.Registers.PC().Int())
		opcode := e.System.Memory.FetchInstruction(addr)
		// TODO remove after prod
		if opcode == OpcodeExit {
			break
		}
		e.ExecOpcode(opcode)
		// increments it by 2 every cpu cycle
		// if the above opcode modifies PC, it will offset it by -2
		// and this line will accomodate for that
		e.System.Registers.IncrementPC()
		// simulate chip8 "clock speed"
		time.Sleep(time.Second / CpuFrequency)
	}
}
