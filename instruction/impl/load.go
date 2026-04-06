package impl

import (
	"github.com/AlanRostem/mu-8/mu8"
	"github.com/AlanRostem/mu-8/system"
)

// LdVxByte executes opcode "8xy0", aka "LD Vx, Vy"
func LdVxVy(args []mu8.DByte, sys *system.System) {
	x := args[0]
	y := args[1]
	sys.Registers.GeneralPurpose[x] = sys.Registers.GeneralPurpose[y]
	pcDebugf(sys.Registers.ProgramCounter, "LD V%X, V%X", x, y)
}

// LdVxByte executes opcode "8xy0", aka "LD Vx, Vy"
func LdVxDt(args []mu8.DByte, sys *system.System) {
	x := args[0]
	sys.Registers.GeneralPurpose[x] = sys.Registers.DelayTimer
	pcDebugf(sys.Registers.ProgramCounter, "LD V%X, DT", x)
}

// LdVxByte executes opcode "6xkk", aka "LD Vx, byte"
func LdVxByte(args []mu8.DByte, sys *system.System) {
	x := args[0]
	kk := args[1]
	sys.Registers.GeneralPurpose[x] = mu8.Byte(kk)
	pcDebugf(sys.Registers.ProgramCounter, "LD V%X, %d", x, kk)
}

// LdIAddr executes opcode "Annn", aka "LD I, addr"
func LdIAddr(args []mu8.DByte, sys *system.System) {
	addr := args[0]
	sys.Registers.Index = addr
	pcDebugf(sys.Registers.ProgramCounter, "LD I, 0x%03X", addr)
}
