package impl

import (
	"github.com/AlanRostem/mu-8/logger"
	"github.com/AlanRostem/mu-8/mu8"
	"github.com/AlanRostem/mu-8/system"
)

func Call(args []mu8.DByte, sys *system.System) {
	nnn := args[0]
	sys.Stack.Push(sys.Registers.ProgramCounter)
	sys.Registers.ProgramCounter = nnn
	logger.Debugf("CALL 0x%03X", nnn)
}

func Jp(args []mu8.DByte, sys *system.System) {
	nnn := args[0]
	sys.Registers.ProgramCounter = nnn
	logger.Debugf("JP 0x%03X", nnn)
}

// 3xkk
func SeVxByte(args []mu8.DByte, sys *system.System) {
	x := args[0]
	kk := args[1]
	vx := sys.Registers.GeneralPurpose[x]
	if mu8.DByte(vx) == kk {
		sys.Registers.ProgramCounter += 2
	}
	logger.Debugf("SE V%01X, 0x%02X", x, kk)
}
