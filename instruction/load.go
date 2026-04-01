package instruction

import (
	"encoding/binary"

	"github.com/AlanRostem/mu-8/mu8"
	"github.com/AlanRostem/mu-8/system"
)

// ldVxByte executes opcode "6xkk", aka "LD Vx, byte"
func ldVxByte(nibbles [MaxNibbles]uint8, sys *system.System) {
	x := nibbles[1]
	lastTwo := []byte{nibbles[2], nibbles[3]}
	kk := binary.BigEndian.Uint16(lastTwo)
	sys.Registers.GeneralPurpose[x] = mu8.Byte(kk)
}

// ldIAddr executes opcode "Annn", aka "LD I, addr"
func ldIAddr(nibbles [MaxNibbles]uint8, sys *system.System) {
	addr := mu8.DByte(0)
	// TODO: verify if this concatenation is correct
	n0 := mu8.DByte(nibbles[0]) << 3 * 4
	n1 := mu8.DByte(nibbles[1]) << 2 * 4
	n2 := mu8.DByte(nibbles[2]) << 1 * 4
	addr = addr & n0
	addr = addr & n1
	addr = addr & n2
	sys.Registers.Index = addr
}
