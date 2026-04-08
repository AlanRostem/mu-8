package instruction

import (
	"github.com/AlanRostem/mu-8/internal/num"
	"github.com/AlanRostem/mu-8/internal/system"
)

func SkpVx(args []num.DByte, sys *system.System) {
	x := args[0]
	pcDebugf(sys.Registers.PC(), "SKP V%X", x)
	key := sys.Registers.V[x] & 0x0F
	keyState := sys.Keys[key]
	if keyState {
		sys.Registers.IncrementPC()
	}
}

func SknpVx(args []num.DByte, sys *system.System) {
	x := args[0]
	pcDebugf(sys.Registers.PC(), "SKNP V%X", x)
	key := sys.Registers.V[x] & 0x0F
	keyState := sys.Keys[key]
	if !keyState {
		sys.Registers.IncrementPC()
	}
}

func LdVxK(args []num.DByte, sys *system.System) {
	x := args[0]
	pcDebugf(sys.Registers.PC(), "LD V%X, K", x)
	// assuming the keyboard is handled on a different thread
	// in the go application
	anyKeyPressed := func() int8 {
		for i, k := range sys.Keys {
			if k {
				return int8(i)
			}
		}
		return -1
	}
	var key int8
	for key = anyKeyPressed(); key == -1; key = anyKeyPressed() {
		// halting until a key is pressed
	}
	sys.Registers.V[x] = num.Byte(key)
}
