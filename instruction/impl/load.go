package impl

import (
	"github.com/AlanRostem/mu-8/mu8"
	"github.com/AlanRostem/mu-8/system"
)

// LdVxByte executes opcode "8xy0", aka "LD Vx, Vy"
func LdVxVy(args []mu8.DByte, sys *system.System) {
	x := args[0]
	y := args[1]
	pcDebugf(sys.Registers.PC(), "LD V%X, V%X", x, y)
	sys.Registers.V[x] = sys.Registers.V[y]
}

// LdVxByte executes opcode "8xy0", aka "LD Vx, Vy"
func LdVxDt(args []mu8.DByte, sys *system.System) {
	x := args[0]
	pcDebugf(sys.Registers.PC(), "LD V%X, DT", x)
	sys.Registers.V[x] = sys.Registers.DT
}

// LdVxByte executes opcode "6xkk", aka "LD Vx, byte"
func LdVxByte(args []mu8.DByte, sys *system.System) {
	x := args[0]
	kk := args[1]
	pcDebugf(sys.Registers.PC(), "LD V%X, %d", x, kk)
	sys.Registers.V[x] = mu8.Byte(kk)
}

// LdIAddr executes opcode "Annn", aka "LD I, addr"
func LdIAddr(args []mu8.DByte, sys *system.System) {
	addr := args[0]
	pcDebugf(sys.Registers.PC(), "LD I, 0x%03X", addr)
	sys.Registers.I = addr
}
