package cpu

import (
	"github.com/AlanRostem/mu-8/mnemonic"
	"github.com/AlanRostem/mu-8/mu8"
)

type ParseInfo struct {
	Mnemonic mnemonic.Mnemonic
	Category mnemonic.Category
	// Nibbles is an array of the full opcode separated into each 4bit value
	Nibbles [4]uint8
}

func ParseOpcode(opcode mu8.DByte) ParseInfo {
	info := ParseInfo{}
	for i := range uint8(4) {
		info.Nibbles[i] = opcode.Nibble(i)
	}
	// TODO implement other categories, for now we only support load instructions
	info.Category = mnemonic.CategoryLoad
	// TODO set the Mnemonic field to the correct one
	return info
}
