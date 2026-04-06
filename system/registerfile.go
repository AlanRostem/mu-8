package system

import "github.com/AlanRostem/mu-8/mu8"

const generalPurposeRegisterCount = 0x10
const stackSize = 16

type RegisterFile struct {
	GeneralPurpose [generalPurposeRegisterCount]mu8.Byte
	ProgramCounter mu8.DByte
	Index          mu8.DByte
	DelayTimer     mu8.Byte
	SoundTimer     mu8.Byte
}

func newRegisterFile() *RegisterFile {
	return &RegisterFile{
		GeneralPurpose: [generalPurposeRegisterCount]mu8.Byte{},
	}
}
