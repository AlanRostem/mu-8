package impl

import (
	"github.com/AlanRostem/mu-8/mu8"
	"github.com/AlanRostem/mu-8/system"
)

func SkpVx(args []mu8.DByte, sys *system.System) {
	x := args[0]
	pcDebugf(sys.Registers.ProgramCounter, "SKP V%X", x)
	// TODO implement
}

func SknpVx(args []mu8.DByte, sys *system.System) {
	x := args[0]
	pcDebugf(sys.Registers.ProgramCounter, "SKNP V%X", x)
	// TODO implement
}
