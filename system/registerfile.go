package system

import "github.com/AlanRostem/mu-8/num"

const generalPurposeRegisterCount = 0x10
const stackSize = 16

const ProgramOffset = 0x200
const indexOffset = 0x050

type RegisterFile struct {
	V  [generalPurposeRegisterCount]num.Byte
	I  num.DByte
	DT num.Byte
	ST num.Byte

	pc num.DByte
}

func newRegisterFile() *RegisterFile {
	return &RegisterFile{
		V: [generalPurposeRegisterCount]num.Byte{},
	}
}

func (rf *RegisterFile) Clear() {
	rf.pc = ProgramOffset
	rf.I = indexOffset
}

func (rf *RegisterFile) SetPC(value num.DByte) {
	rf.pc = value - 2 // incremented next cpu cycle anyway, so this cancels out
}

func (rf *RegisterFile) IncrementPC() {
	rf.pc += 2
}

func (rf *RegisterFile) PC() num.DByte {
	return rf.pc
}
