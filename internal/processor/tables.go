package processor

import (
	"github.com/AlanRostem/mu-8/internal/decode"
	"github.com/AlanRostem/mu-8/internal/instruction"
)

var tableClass0 = newInstructionTable()
var tableClass1 = newSingularInstructionTable(instruction.Jp)
var tableClass2 = newSingularInstructionTable(instruction.Call)
var tableClass3 = newSingularInstructionTable(instruction.SeVxByte)
var tableClass4 = newSingularInstructionTable(instruction.SneVxByte)
var tableClass6 = newSingularInstructionTable(instruction.LdVxByte)
var tableClass7 = newSingularInstructionTable(instruction.AddVxByte)
var tableClass8 = newInstructionTable()
var tableClassA = newSingularInstructionTable(instruction.LdIAddr)

var tableClassD = newSingularInstructionTable(instruction.DrwVxVyN)
var tableClassE = newInstructionTable()
var tableClassF = newInstructionTable()

func init() {
	tableClass0.Add(0x000, instruction.Sys)
	tableClass0.Add(0x0E0, instruction.Cls)

	tableClass8.Add(0x0, instruction.LdVxVy)

	tableClassE.Add(0x9E, instruction.SkpVx)
	tableClassE.Add(0xA1, instruction.SknpVx)

	tableClassF.Add(0x07, instruction.LdVxDt)
}

// tableAll contains instruction maps mapped to a class number.
var tableAll = map[decode.Class]instructionTable{
	decode.Class0: tableClass0,
	decode.Class1: tableClass1,
	decode.Class2: tableClass2,
	decode.Class3: tableClass3,
	decode.Class4: tableClass4,
	decode.Class6: tableClass6,
	decode.Class7: tableClass7,
	decode.Class8: tableClass8,
	decode.ClassA: tableClassA,
	decode.ClassD: tableClassD,
	decode.ClassE: tableClassE,
	decode.ClassF: tableClassF,
}
