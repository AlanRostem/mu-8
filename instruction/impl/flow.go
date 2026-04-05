package impl

import (
	"github.com/AlanRostem/mu-8/mu8"
	"github.com/AlanRostem/mu-8/system"
)

func Jp(args []mu8.DByte, sys *system.System) {
	nnn := args[0]
	sys.Registers.ProgramCounter = nnn
}
