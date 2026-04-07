package processor

import (
	"github.com/AlanRostem/mu-8/internal/decode"
	"github.com/AlanRostem/mu-8/internal/num"
	"github.com/AlanRostem/mu-8/internal/system"
)

const Frequency = 500

type Processor struct {
	System *system.System
}

func New(system *system.System) *Processor {
	return &Processor{
		System: system,
	}
}

func (p *Processor) Execute(opcode num.DByte) {
	info := decode.Decode(opcode)
	table := tableAll[info.Class]
	if table.IsSingle() {
		table.Single()(info.Args, p.System)
		return
	}
	identity := info.Idenitity
	inst := table[identity]
	inst(info.Args, p.System)
}

func (p *Processor) Cycle() {
	addr := num.NewUint12(p.System.Registers.PC().Int())
	opcode := p.System.Memory.FetchInstruction(addr)
	p.Execute(opcode)
	// increments it by 2 every cpu cycle
	// if the above opcode modifies PC, it will offset it by -2
	// and this line will accomodate for that
	p.System.Registers.IncrementPC()
}
