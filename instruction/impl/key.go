package impl

import (
	"github.com/AlanRostem/mu-8/num"
	"github.com/AlanRostem/mu-8/system"
)

func SkpVx(args []num.DByte, sys *system.System) {
	x := args[0]
	pcDebugf(sys.Registers.PC(), "SKP V%X", x)
	// TODO implement
}

func SknpVx(args []num.DByte, sys *system.System) {
	x := args[0]
	pcDebugf(sys.Registers.PC(), "SKNP V%X", x)
	// TODO implement
}
