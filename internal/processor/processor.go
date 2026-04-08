package processor

import (
	"fmt"

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

func (p *Processor) execute(opcode num.DByte) {
	info := decode.Decode(opcode)
	table := tableAll[info.Class]
	if table.IsSingle() {
		table.Single()(info.Args, p.System)
		return
	}
	identity := info.Idenitity
	inst, ok := table[identity]
	if !ok {
		panic(fmt.Errorf("instruction not found: class=0x%X, identity=0x%X", info.Class, identity))
	}
	inst(info.Args, p.System)
}

func (p *Processor) Cycle() {
	addr := num.NewUint12(p.System.Registers.PC().Int())
	opcode := p.System.Memory.FetchInstruction(addr)
	p.execute(opcode)
	// increments it by 2 every cpu cycle
	// if the above opcode modifies PC, it will offset it by -2
	// and this line will accomodate for that
	p.System.Registers.IncrementPC()
}
