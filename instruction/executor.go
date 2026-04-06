package instruction

import (
	"time"

	"github.com/AlanRostem/mu-8/decode"
	"github.com/AlanRostem/mu-8/mu8"
	"github.com/AlanRostem/mu-8/system"
)

const ProgramOffset = 0x200
const ProgramSize = system.MemorySize - ProgramOffset
const IndexOffset = 0x050
const CpuFrequency = 500

// OpcodeExit is a special opcode used by the Executor to
// exit the program for testing purposes. This will be
// removed in prod.
const OpcodeExit = mu8.DByte(0xFFFF)

type Executor struct {
	System *system.System
}

func NewExecutor(system *system.System) *Executor {
	return &Executor{
		System: system,
	}
}

func (e *Executor) ExecOpcode(opcode mu8.DByte) {
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

func (e *Executor) LoadProgram(program [ProgramSize]mu8.Byte) {
	for i, value := range program {
		addr := mu8.NewUint12(ProgramOffset + i)
		e.System.Memory.Write(addr, value)
	}
	e.System.Registers.ProgramCounter = ProgramOffset
	e.System.Registers.Index = IndexOffset
}

func (e *Executor) ExecProgram() {
	for {
		addr := mu8.NewUint12(e.System.Registers.ProgramCounter.Int())
		left := e.System.Memory.Read(addr)
		addr.Add(1)
		right := e.System.Memory.Read(addr)
		opcode := left.Concat(right)
		if opcode == OpcodeExit {
			break
		}
		e.ExecOpcode(opcode)
		// TODO verify if incrementing PC here is correct
		e.System.Registers.ProgramCounter += 2
		time.Sleep(time.Second / CpuFrequency)
	}
}
