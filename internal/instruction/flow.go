package instruction

import (
	"github.com/AlanRostem/mu-8/internal/num"
	"github.com/AlanRostem/mu-8/internal/system"
)

func CallAddr(args []num.DByte, sys *system.System) {
	nnn := args[0]
	pcDebugf(sys.Registers.PC(), "CALL 0x%03X", nnn)
	sys.Stack.Push(sys.Registers.PC())
	sys.Registers.JumpPC(nnn)
}

func Ret(args []num.DByte, sys *system.System) {
	panic("not implemented")
}

func JpAddr(args []num.DByte, sys *system.System) {
	nnn := args[0]
	pcDebugf(sys.Registers.PC(), "JP 0x%03X", nnn)
	sys.Registers.JumpPC(nnn)
}

func JpV0Addr(args []num.DByte, sys *system.System) {
	nnn := args[0]
	pcDebugf(sys.Registers.PC(), "JP V0, 0x%03X", nnn)
	v0 := sys.Registers.V[0]
	sys.Registers.JumpPC(num.DByte(v0) + nnn)
}

// 3xkk
func SeVxByte(args []num.DByte, sys *system.System) {
	x := args[0]
	kk := args[1]
	pcDebugf(sys.Registers.PC(), "SE V%01X, %d", x, kk)
	vx := sys.Registers.V[x]
	if num.DByte(vx) == kk {
		sys.Registers.IncrementPC()
	}
}

// 3xkk
func SneVxByte(args []num.DByte, sys *system.System) {
	x := args[0]
	kk := args[1]
	pcDebugf(sys.Registers.PC(), "SNE V%01X, %d", x, kk)
	vx := sys.Registers.V[x]
	if num.DByte(vx) != kk {
		sys.Registers.IncrementPC()
	}
}

func SeVxVy(args []num.DByte, sys *system.System) {
	panic("not implemented")
}

func SneVxVy(args []num.DByte, sys *system.System) {
	panic("not implemented")
}
