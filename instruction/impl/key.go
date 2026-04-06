package impl

import (
	"github.com/AlanRostem/mu-8/mu8"
	"github.com/AlanRostem/mu-8/system"
)

func SkpVx(args []mu8.DByte, sys *system.System) {
	x := args[0]
	// TODO implement
	pcDebugf(sys.Registers.ProgramCounter, "SKP V%X", x)
}

func SknpVx(args []mu8.DByte, sys *system.System) {
	x := args[0]
	// TODO implement
	pcDebugf(sys.Registers.ProgramCounter, "SKNP V%X", x)
}
