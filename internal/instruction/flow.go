package instruction

import (
	"github.com/AlanRostem/mu-8/internal/num"
	"github.com/AlanRostem/mu-8/internal/system"
)

func CallAddr(args []num.DByte, sys *system.System) {
	nnn := args[0]
	pcDebugf(sys.Registers.PC(), "CALL 0x%03X", nnn)
	sys.Stack.Push(sys.Registers.PC() + 2)
	sys.Registers.JumpPC(nnn)
}

func Ret(args []num.DByte, sys *system.System) {
	addr := sys.Stack.Pop()
	pcDebugf(sys.Registers.PC(), "RET 0x%03X", addr)
	sys.Registers.JumpPC(addr)
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

func SeVxByte(args []num.DByte, sys *system.System) {
	x := args[0]
	kk := args[1]
	pcDebugf(sys.Registers.PC(), "SE V%01X, %d", x, kk)
	vx := sys.Registers.V[x]
	if num.DByte(vx) == kk {
		sys.Registers.IncrementPC()
	}
}

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
	x := args[0]
	y := args[1]
	pcDebugf(sys.Registers.PC(), "SE V%X, V%X", x, y)
	if sys.Registers.V[x] == sys.Registers.V[y] {
		sys.Registers.IncrementPC()
	}
}

func SneVxVy(args []num.DByte, sys *system.System) {
	x := args[0]
	y := args[1]
	pcDebugf(sys.Registers.PC(), "SNE V%X, V%X", x, y)
	if sys.Registers.V[x] != sys.Registers.V[y] {
		sys.Registers.IncrementPC()
	}
}
