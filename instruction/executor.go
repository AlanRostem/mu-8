package instruction

import (
	"github.com/AlanRostem/mu-8/decode"
	"github.com/AlanRostem/mu-8/mu8"
	"github.com/AlanRostem/mu-8/system"
)

const ProgramOffset = 0x200
const IndexOffset = 0x050

type Executor struct {
	System *system.System
}

func NewExecutor(system *system.System) *Executor {
	return &Executor{
		System: system,
	}
}

func (e *Executor) Exec(instruction mu8.DByte) {
	info := decode.Decode(instruction)
	table := tableAll[info.Class]
	if table.IsSingle() {
		table.Single()(info.Args, e.System)
		return
	}
	identity := info.Idenitity
	inst := table[identity]
	inst(info.Args, e.System)
}

func (e *Executor) Load(program [system.MemorySize - ProgramOffset]mu8.Byte) {
	for i, value := range program {
		addr := mu8.NewUint12(ProgramOffset + i)
		e.System.Memory.Write(addr, value)
	}
	e.System.Registers.ProgramCounter = ProgramOffset
	e.System.Registers.Index = IndexOffset
}
