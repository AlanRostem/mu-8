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
	panic("not implemented")
}

func AndVxVy(args []num.DByte, sys *system.System) {
	panic("not implemented")
}

func XorVxVy(args []num.DByte, sys *system.System) {
	panic("not implemented")
}

func AddVxVy(args []num.DByte, sys *system.System) {
	panic("not implemented")
}

func SubVxVy(args []num.DByte, sys *system.System) {
	panic("not implemented")
}

func ShrVxVy(args []num.DByte, sys *system.System) {
	panic("not implemented")
}

func SubnVxVy(args []num.DByte, sys *system.System) {
	panic("not implemented")
}

func ShlVxVy(args []num.DByte, sys *system.System) {
	panic("not implemented")
}

func AddIVx(args []num.DByte, sys *system.System) {
	panic("not implemented")
}
