package instruction

import (
	"github.com/AlanRostem/mu-8/decode"
	"github.com/AlanRostem/mu-8/instruction/impl"
)

var tableClass6 = newSingularInstructionTable(impl.LdVxByte)
var tableClassA = newSingularInstructionTable(impl.LdIAddr)

var tableClassD = newSingularInstructionTable(impl.DrwVxVyN)
var tableClassE = newInstructionTable()

func init() {
	tableClassE.Add(0x9E, impl.SkpVx)
	tableClassE.Add(0xA1, impl.SknpVx)
}

// tableAll contains instruction maps mapped to a class number.
var tableAll = map[decode.Class]instructionTable{
	decode.Class6: tableClass6,
	decode.ClassA: tableClassA,
	decode.ClassE: tableClassE,
	decode.ClassD: tableClassD,
}
