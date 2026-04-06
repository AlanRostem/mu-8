package impl

import (
	"github.com/AlanRostem/mu-8/logger"
	"github.com/AlanRostem/mu-8/mu8"
	"github.com/AlanRostem/mu-8/system"
)

func Sys(args []mu8.DByte, _ *system.System) {
	// it is intentional that this does nothing
	logger.Debugf("SYS %03X (not implemented)", args[0])
}
