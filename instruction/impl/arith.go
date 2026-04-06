package impl

import (
	"github.com/AlanRostem/mu-8/mu8"
	"github.com/AlanRostem/mu-8/system"
)

// 7xkk
func AddVxByte(args []mu8.DByte, sys *system.System) {
	x := args[0]
	kk := args[1]
	pcDebugf(sys.Registers.ProgramCounter, "ADD V%01X, %d", x, kk)
	sys.Registers.GeneralPurpose[x] += mu8.Byte(kk)
}
