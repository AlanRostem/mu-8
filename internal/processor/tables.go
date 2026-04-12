package processor

import (
	"github.com/AlanRostem/mu-8/internal/decode"
	"github.com/AlanRostem/mu-8/internal/instruction"
)

var tableClass0 = newInstructionTable()
var tableClass1 = newSingularInstructionTable(instruction.JpAddr)
var tableClass2 = newSingularInstructionTable(instruction.CallAddr)
var tableClass3 = newSingularInstructionTable(instruction.SeVxByte)
var tableClass4 = newSingularInstructionTable(instruction.SneVxByte)
var tableClass5 = newSingularInstructionTable(instruction.SeVxVy)
var tableClass6 = newSingularInstructionTable(instruction.LdVxByte)
var tableClass7 = newSingularInstructionTable(instruction.AddVxByte)
var tableClass8 = newInstructionTable()
var tableClass9 = newSingularInstructionTable(instruction.SneVxVy)
var tableClassA = newSingularInstructionTable(instruction.LdIAddr)
var tableClassB = newSingularInstructionTable(instruction.JpV0Addr)
var tableClassC = newSingularInstructionTable(instruction.RndVxByte)
var tableClassD = newSingularInstructionTable(instruction.DrwVxVyN)
var tableClassE = newInstructionTable()
var tableClassF = newInstructionTable()

func init() {
	tableClass0.Add(0x000, instruction.Sys)
	tableClass0.Add(0x0E0, instruction.Cls)
	tableClass0.Add(0x0EE, instruction.Ret)

	tableClass8.Add(0x0, instruction.LdVxVy)
	tableClass8.Add(0x1, instruction.OrVxVy)
	tableClass8.Add(0x2, instruction.AndVxVy)
	tableClass8.Add(0x3, instruction.XorVxVy)
	tableClass8.Add(0x4, instruction.AddVxVy)
	tableClass8.Add(0x5, instruction.SubVxVy)
	tableClass8.Add(0x6, instruction.ShrVxVy)
	tableClass8.Add(0x7, instruction.SubnVxVy)
	tableClass8.Add(0xE, instruction.ShlVxVy)

	tableClassE.Add(0x9E, instruction.SkpVx)
	tableClassE.Add(0xA1, instruction.SknpVx)

	tableClassF.Add(0x07, instruction.LdVxDt)
	tableClassF.Add(0x0A, instruction.LdVxK)
	tableClassF.Add(0x15, instruction.LdDtVx)
	tableClassF.Add(0x18, instruction.LdStVx)
	tableClassF.Add(0x1E, instruction.AddIVx)
	tableClassF.Add(0x29, instruction.LdFVx)
	tableClassF.Add(0x33, instruction.LdBVx)
	tableClassF.Add(0x55, instruction.LdIVx)
	tableClassF.Add(0x65, instruction.LdVxI)
}

// tableAll contains instruction maps mapped to a class number.
var tableAll = map[decode.Class]instructionTable{
	decode.Class0: tableClass0,
	decode.Class1: tableClass1,
	decode.Class2: tableClass2,
	decode.Class3: tableClass3,
	decode.Class4: tableClass4,
	decode.Class5: tableClass5,
	decode.Class6: tableClass6,
	decode.Class7: tableClass7,
	decode.Class8: tableClass8,
	decode.Class9: tableClass9,
	decode.ClassA: tableClassA,
	decode.ClassB: tableClassB,
	decode.ClassC: tableClassC,
	decode.ClassD: tableClassD,
	decode.ClassE: tableClassE,
	decode.ClassF: tableClassF,
}
