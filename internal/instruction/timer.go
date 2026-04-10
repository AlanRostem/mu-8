package instruction

import (
	"github.com/AlanRostem/mu-8/internal/num"
	"github.com/AlanRostem/mu-8/internal/system"
)

// LdVxDt executes opcode "Fx07	", aka "LD Vx, DT"
func LdVxDt(args []num.DByte, sys *system.System) {
	x := args[0]
	pcDebugf(sys.Registers.PC(), "LD V%X, ST", x)
	sys.Registers.V[x] = sys.Timers.DT()
}

// LdDtVx executes opcode "Fx15", aka "LD DT, Vx"
func LdDtVx(args []num.DByte, sys *system.System) {
	x := args[0]
	pcDebugf(sys.Registers.PC(), "LD DT, V%X", x)
	sys.Timers.SetDT(sys.Registers.V[x])
}

// LdStVx executes opcode "Fx18", aka "LD ST, Vx"
func LdStVx(args []num.DByte, sys *system.System) {
	x := args[0]
	pcDebugf(sys.Registers.PC(), "LD ST, V%X", x)
	sys.Timers.SetST(sys.Registers.V[x])
}
