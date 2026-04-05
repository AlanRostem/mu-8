package impl

import (
	"fmt"

	"github.com/AlanRostem/mu-8/mu8"
	"github.com/AlanRostem/mu-8/system"
)

// LdVxByte executes opcode "8xy0", aka "LD Vx, Vy"
func LdVxVy(args []mu8.DByte, sys *system.System) {
	x := args[0]
	y := args[1]
	sys.Registers.GeneralPurpose[x] = sys.Registers.GeneralPurpose[y]
	fmt.Printf("LD V%X, V%X\n", x, y)
}

// LdVxByte executes opcode "6xkk", aka "LD Vx, byte"
func LdVxByte(args []mu8.DByte, sys *system.System) {
	x := args[0]
	kk := args[1]
	sys.Registers.GeneralPurpose[x] = mu8.Byte(kk)
	fmt.Printf("LD V%X, 0x%02X\n", x, kk)
}

// LdIAddr executes opcode "Annn", aka "LD I, addr"
func LdIAddr(args []mu8.DByte, sys *system.System) {
	addr := args[0]
	sys.Registers.Index = addr
	fmt.Printf("LD I, 0x%03X\n", addr)
}
