package instruction

import (
	"github.com/AlanRostem/mu-8/internal/num"
	"github.com/AlanRostem/mu-8/internal/system"
)

func Sys(args []num.DByte, sys *system.System) {
	// it is intentional that this does nothing
	pcDebugf(sys.Registers.PC(), "SYS %03X (not implemented)", args[0])
}
