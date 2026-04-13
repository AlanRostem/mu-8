package instruction

import (
	"math/rand/v2"

	"github.com/AlanRostem/mu-8/internal/num"
	"github.com/AlanRostem/mu-8/internal/system"
)

// 7xkk
func AddVxByte(args []num.DByte, sys *system.System) {
	x := args[0]
	kk := args[1]
	pcDebugf(sys.Registers.PC(), "ADD V%01X, %d", x, kk)
	sys.Registers.V[x] += num.Byte(kk)
}

func RndVxByte(args []num.DByte, sys *system.System) {
	x := args[0]
	kk := args[1]
	pcDebugf(sys.Registers.PC(), "RND V%01X, %d", x, kk)
	r := rand.UintN(256)
	sys.Registers.V[x] = num.Byte(r) & num.Byte(kk)
}

func OrVxVy(args []num.DByte, sys *system.System) {
	x := args[0]
	y := args[1]
	pcDebugf(sys.Registers.PC(), "OR, V%X, V%X", x, y)
	sys.Registers.V[x] |= sys.Registers.V[y]
}

func AndVxVy(args []num.DByte, sys *system.System) {
	x := args[0]
	y := args[1]
	pcDebugf(sys.Registers.PC(), "AND, V%X, V%X", x, y)
	sys.Registers.V[x] &= sys.Registers.V[y]
}

func XorVxVy(args []num.DByte, sys *system.System) {
	x := args[0]
	y := args[1]
	pcDebugf(sys.Registers.PC(), "XOR, V%X, V%X", x, y)
	sys.Registers.V[x] ^= sys.Registers.V[y]
}

func AddVxVy(args []num.DByte, sys *system.System) {
	x := args[0]
	y := args[1]
	pcDebugf(sys.Registers.PC(), "ADD, V%X, V%X", x, y)
	result := num.DByte(sys.Registers.V[x]) + num.DByte(sys.Registers.V[y])
	if result > 0xFF {
		result &= 0x00FF
		sys.Registers.V[0xF] = 1
	} else {
		sys.Registers.V[0xF] = 0
	}
	sys.Registers.V[x] = num.Byte(result)
}

func AddIVx(args []num.DByte, sys *system.System) {
	x := args[0]
	pcDebugf(sys.Registers.PC(), "ADD, I, V%X", x)
	sys.Registers.I += num.DByte(sys.Registers.V[x])
}

func SubVxVy(args []num.DByte, sys *system.System) {
	x := args[0]
	y := args[1]
	pcDebugf(sys.Registers.PC(), "SUB, V%X, V%X", x, y)
	if int16(sys.Registers.V[x]) >= int16(sys.Registers.V[y]) {
		sys.Registers.V[0xF] = 1
	} else {
		sys.Registers.V[0xF] = 0
	}
	sys.Registers.V[x] -= sys.Registers.V[y]
}

func SubnVxVy(args []num.DByte, sys *system.System) {
	x := args[0]
	y := args[1]
	pcDebugf(sys.Registers.PC(), "SUB, V%X, V%X", x, y)
	if int16(sys.Registers.V[y]) >= int16(sys.Registers.V[x]) {
		sys.Registers.V[0xF] = 1
	} else {
		sys.Registers.V[0xF] = 0
	}
	sys.Registers.V[x] = sys.Registers.V[y] - sys.Registers.V[x]
}

func ShrVxVy(args []num.DByte, sys *system.System) {
	x := args[0]
	y := args[1]
	pcDebugf(sys.Registers.PC(), "SHR V%X, V%X", x, y)
	sys.Registers.V[x] >>= 1
	sys.Registers.V[0xF] = sys.Registers.V[x] & 0b10000000
}

func ShlVxVy(args []num.DByte, sys *system.System) {
	x := args[0]
	y := args[1]
	pcDebugf(sys.Registers.PC(), "SHL V%X, V%X", x, y)
	sys.Registers.V[x] <<= 1
	sys.Registers.V[0xF] = sys.Registers.V[x] & 0b00000001
}
