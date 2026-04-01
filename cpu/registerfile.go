package cpu

import "github.com/AlanRostem/mu-8/mu8"

const GeneralPurposeRegisterCount = 0x10

type RegisterFile struct {
	GeneralPurpose [GeneralPurposeRegisterCount]mu8.Byte
	Index          mu8.Word
	ProgramCounter mu8.Word
	StackPointer   mu8.Word
	DelayTimer     mu8.Byte
	SoundTime      mu8.Byte
	CurrentOperand mu8.Word
}

func NewRegisterFile() *RegisterFile {
	return &RegisterFile{
		GeneralPurpose: [GeneralPurposeRegisterCount]mu8.Byte{},
	}
}
