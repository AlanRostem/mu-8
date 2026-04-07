package impl

import (
	"github.com/AlanRostem/mu-8/num"
	"github.com/AlanRostem/mu-8/system"
)

func Sys(args []num.DByte, sys *system.System) {
	// it is intentional that this does nothing
	pcDebugf(sys.Registers.PC(), "SYS %03X (not implemented)", args[0])
}
