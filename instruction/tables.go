package instruction

import (
	"github.com/AlanRostem/mu-8/decode"
	"github.com/AlanRostem/mu-8/instruction/impl"
)

var tableClass0 = newInstructionTable()
var tableClass6 = newSingularInstructionTable(impl.LdVxByte)
var tableClass8 = newInstructionTable()
var tableClassA = newSingularInstructionTable(impl.LdIAddr)

var tableClassD = newSingularInstructionTable(impl.DrwVxVyN)
var tableClassE = newInstructionTable()

func init() {
	tableClass0.Add(0x000, impl.Sys)
	tableClass0.Add(0x0E0, impl.Cls)

	tableClass8.Add(0x0, impl.LdVxVy)

	tableClassE.Add(0x9E, impl.SkpVx)
	tableClassE.Add(0xA1, impl.SknpVx)
}

// tableAll contains instruction maps mapped to a class number.
var tableAll = map[decode.Class]instructionTable{
	decode.Class0: tableClass0,
	decode.Class6: tableClass6,
	decode.Class8: tableClass8,
	decode.ClassA: tableClassA,
	decode.ClassE: tableClassE,
	decode.ClassD: tableClassD,
}
