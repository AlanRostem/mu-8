package decode

import "github.com/AlanRostem/mu-8/mu8"

type Argument struct {
	NibblePosition uint8
	NibbleSize     uint8
	Value          uint16
}

type DecodeInfo struct {
	Class     Class  // First nibble which is can be the same for several opcodes
	Idenitity uint16 // Usually the last few nibbles of the opcode
	Args      []uint16
}

func Decode(instruction mu8.DByte) DecodeInfo {
	nibbles := [4]uint8{}
	for i := range uint8(4) {
		// TODO check if this ordering makes sense
		nibbles[4-i] = instruction.Nibble(i)
	}
	class := Class(nibbles[0])
	// classInfo := Classes[class]
	// TODO finish implementation
	return DecodeInfo{
		Class: class,
	}
}
