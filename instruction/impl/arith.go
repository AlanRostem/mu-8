package impl

import (
	"github.com/AlanRostem/mu-8/num"
	"github.com/AlanRostem/mu-8/system"
)

// 7xkk
func AddVxByte(args []num.DByte, sys *system.System) {
	x := args[0]
	kk := args[1]
	pcDebugf(sys.Registers.PC(), "ADD V%01X, %d", x, kk)
	sys.Registers.V[x] += num.Byte(kk)
}
