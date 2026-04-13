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

func LdFVx(args []num.DByte, sys *system.System) {
	x := args[0]
	pcDebugf(sys.Registers.PC(), "LD F, V%X", x)
	const fontStart = 0x50
	const fontSpriteHeight = 5
	digit := sys.Registers.V[x]
	sys.Registers.I = fontStart + num.DByte(digit)*fontSpriteHeight
}

func LdBVx(args []num.DByte, sys *system.System) {
	x := args[0]
	pcDebugf(sys.Registers.PC(), "LD B, V%X", x)
	vx := sys.Registers.V[x]
	hundred := vx / 100
	ten := (vx / 10) % 10
	one := vx % 10
	addr := num.NewUint12(sys.Registers.I.Int())
	sys.Memory.Write(addr, hundred)
	addr.Add(1)
	sys.Memory.Write(addr, ten)
	addr.Add(1)
	sys.Memory.Write(addr, one)

}

func LdIVx(args []num.DByte, sys *system.System) {
	x := args[0]
	pcDebugf(sys.Registers.PC(), "LD I, V%X", x)
	start := sys.Registers.I.Int()
	for i := range x + 1 {
		addr := num.NewUint12(start + i.Int())
		sys.Memory.Write(addr, sys.Registers.V[i])
	}
	sys.Registers.I += x + 1
}

func LdVxI(args []num.DByte, sys *system.System) {
	x := args[0]
	pcDebugf(sys.Registers.PC(), "LD V%X, I", x)
	start := sys.Registers.I.Int()
	for i := range x + 1 {
		addr := num.NewUint12(start + i.Int())
		sys.Registers.V[i] = sys.Memory.Read(addr)
	}
	sys.Registers.I += x + 1
}
