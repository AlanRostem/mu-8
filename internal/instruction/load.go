package instruction

import (
	"github.com/AlanRostem/mu-8/internal/num"
	"github.com/AlanRostem/mu-8/internal/system"
)

// LdVxByte executes opcode "8xy0", aka "LD Vx, Vy"
func LdVxVy(args []num.DByte, sys *system.System) {
	x := args[0]
	y := args[1]
	pcDebugf(sys.Registers.PC(), "LD V%X, V%X", x, y)
	sys.Registers.V[x] = sys.Registers.V[y]
}

// LdVxByte executes opcode "Fx15", aka "LD DT, Vx"
func LdVxDt(args []num.DByte, sys *system.System) {
	x := args[0]
	pcDebugf(sys.Registers.PC(), "LD DT, V%X", x)
	sys.Timers.SetDT(sys.Registers.V[x])
}

// LdVxByte executes opcode "6xkk", aka "LD Vx, byte"
func LdVxByte(args []num.DByte, sys *system.System) {
	x := args[0]
	kk := args[1]
	pcDebugf(sys.Registers.PC(), "LD V%X, %d", x, kk)
	sys.Registers.V[x] = num.Byte(kk)
}

// LdIAddr executes opcode "Annn", aka "LD I, addr"
func LdIAddr(args []num.DByte, sys *system.System) {
	addr := args[0]
	pcDebugf(sys.Registers.PC(), "LD I, 0x%03X", addr)
	sys.Registers.I = addr
}
