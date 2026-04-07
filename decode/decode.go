package decode

import "github.com/AlanRostem/mu-8/num"

type Argument struct {
	NibblePosition uint8
	NibbleSize     uint8
	Value          uint16
}

type Info struct {
	Class          Class     // First nibble which is can be the same for several opcodes
	Idenitity      num.DByte // Usually the last few nibbles of the opcode
	Args           []num.DByte
	SingleIdentity bool
}

// TODO set the 0x000 instructions identity to 0 instead of
// making it a single identity
func class0Decoded(instruction num.DByte) Info {
	identity := instruction & 0xFFF
	const CLS = num.DByte(0x0E0)
	const RET = num.DByte(0x0EE)
	// jump wont be implemented (I think)
	isJump := identity != CLS && identity != RET
	return Info{
		Class:          Class0,
		Args:           make([]num.DByte, 0),
		SingleIdentity: isJump,
		Idenitity:      identity,
	}
}

func Decode(instruction num.DByte) Info {
	nibbles := instruction.GetNibbles()
	class := Class(nibbles[0])
	if class == Class0 {
		return class0Decoded(instruction)
	}
	classInfo := allClasses[class]
	argValues := make([]num.DByte, 0)
	for i, argSize := range classInfo.Args {
		nibIdx := uint8(i + 1)
		argVal := num.DByte(0)
		for j := range argSize {
			argVal.AppendNibble(nibbles[nibIdx+j])
		}
		argValues = append(argValues, argVal)
	}
	identity := num.DByte(0)
	if classInfo.IdentitySize != IdentitySizeNone {
		nibsToAppend := make([]uint8, 0)
		for i := range classInfo.IdentitySize {
			nibsToAppend = append(nibsToAppend, nibbles[3-i])
		}
		for i := range nibsToAppend {
			identity.AppendNibble(nibsToAppend[len(nibsToAppend)-i-1])
		}
	}
	return Info{
		Class:          class,
		Args:           argValues,
		Idenitity:      identity,
		SingleIdentity: classInfo.IdentitySize != IdentitySizeNone,
	}
}
