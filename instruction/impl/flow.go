package impl

import (
	"github.com/AlanRostem/mu-8/mu8"
	"github.com/AlanRostem/mu-8/system"
)

func Call(args []mu8.DByte, sys *system.System) {
	nnn := args[0]
	pcDebugf(sys.Registers.ProgramCounter, "CALL 0x%03X", nnn)
	sys.Stack.Push(sys.Registers.ProgramCounter)
	// HACK since the PC is incremented always, we offset it by -2 so it lands on
	// the desired address later
	sys.Registers.ProgramCounter = nnn - 2
}

func Jp(args []mu8.DByte, sys *system.System) {
	nnn := args[0]
	pcDebugf(sys.Registers.ProgramCounter, "JP 0x%03X", nnn)
	// HACK since the PC is incremented always, we offset it by -2 so it lands on
	// the desired address later
	sys.Registers.ProgramCounter = nnn - 2
}

// 3xkk
func SeVxByte(args []mu8.DByte, sys *system.System) {
	x := args[0]
	kk := args[1]
	pcDebugf(sys.Registers.ProgramCounter, "SE V%01X, %d", x, kk)
	vx := sys.Registers.GeneralPurpose[x]
	if mu8.DByte(vx) == kk {
		sys.Registers.ProgramCounter += 2
	}
}

// 3xkk
func SneVxByte(args []mu8.DByte, sys *system.System) {
	x := args[0]
	kk := args[1]
	pcDebugf(sys.Registers.ProgramCounter, "SNE V%01X, %d", x, kk)
	vx := sys.Registers.GeneralPurpose[x]
	if mu8.DByte(vx) != kk {
		sys.Registers.ProgramCounter += 2
	}
}
