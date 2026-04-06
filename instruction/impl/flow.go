package impl

import (
	"github.com/AlanRostem/mu-8/mu8"
	"github.com/AlanRostem/mu-8/system"
)

func Call(args []mu8.DByte, sys *system.System) {
	nnn := args[0]
	pcDebugf(sys.Registers.ProgramCounter, "CALL 0x%03X", nnn)
	sys.Stack.Push(sys.Registers.ProgramCounter)
	sys.Registers.ProgramCounter = nnn
}

func Jp(args []mu8.DByte, sys *system.System) {
	nnn := args[0]
	sys.Registers.ProgramCounter = nnn
	pcDebugf(sys.Registers.ProgramCounter, "JP 0x%03X", nnn)
}

// 3xkk
func SeVxByte(args []mu8.DByte, sys *system.System) {
	x := args[0]
	kk := args[1]
	vx := sys.Registers.GeneralPurpose[x]
	if mu8.DByte(vx) == kk {
		sys.Registers.ProgramCounter += 2
	}
	pcDebugf(sys.Registers.ProgramCounter, "SE V%01X, %d", x, kk)
}
