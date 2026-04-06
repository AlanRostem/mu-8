package system

import "github.com/AlanRostem/mu-8/mu8"

const generalPurposeRegisterCount = 0x10
const stackSize = 16

const ProgramOffset = 0x200
const indexOffset = 0x050

type RegisterFile struct {
	V  [generalPurposeRegisterCount]mu8.Byte
	I  mu8.DByte
	DT mu8.Byte
	ST mu8.Byte

	pc mu8.DByte
}

func newRegisterFile() *RegisterFile {
	return &RegisterFile{
		V: [generalPurposeRegisterCount]mu8.Byte{},
	}
}

func (rf *RegisterFile) Clear() {
	rf.pc = ProgramOffset
	rf.I = indexOffset
}

func (rf *RegisterFile) SetPC(value mu8.DByte) {
	rf.pc = value - 2 // incremented next cpu cycle anyway, so this cancels out
}

func (rf *RegisterFile) IncrementPC() {
	rf.pc += 2
}

func (rf *RegisterFile) PC() mu8.DByte {
	return rf.pc
}
