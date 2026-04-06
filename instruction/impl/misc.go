package impl

import (
	"github.com/AlanRostem/mu-8/mu8"
	"github.com/AlanRostem/mu-8/system"
)

func Sys(args []mu8.DByte, sys *system.System) {
	// it is intentional that this does nothing
	pcDebugf(sys.Registers.ProgramCounter, "SYS %03X (not implemented)", args[0])
}
