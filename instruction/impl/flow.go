package impl

import (
	"github.com/AlanRostem/mu-8/mu8"
	"github.com/AlanRostem/mu-8/system"
)

func Call(args []mu8.DByte, sys *system.System) {
	nnn := args[0]
	pcDebugf(sys.Registers.PC(), "CALL 0x%03X", nnn)
	sys.Stack.Push(sys.Registers.PC())
	sys.Registers.SetPC(nnn)
}

func Jp(args []mu8.DByte, sys *system.System) {
	nnn := args[0]
	pcDebugf(sys.Registers.PC(), "JP 0x%03X", nnn)
	sys.Registers.SetPC(nnn)
}

// 3xkk
func SeVxByte(args []mu8.DByte, sys *system.System) {
	x := args[0]
	kk := args[1]
	pcDebugf(sys.Registers.PC(), "SE V%01X, %d", x, kk)
	vx := sys.Registers.V[x]
	if mu8.DByte(vx) == kk {
		sys.Registers.IncrementPC()
	}
}

// 3xkk
func SneVxByte(args []mu8.DByte, sys *system.System) {
	x := args[0]
	kk := args[1]
	pcDebugf(sys.Registers.PC(), "SNE V%01X, %d", x, kk)
	vx := sys.Registers.V[x]
	if mu8.DByte(vx) != kk {
		sys.Registers.IncrementPC()
	}
}
