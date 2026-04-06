package impl

import (
	"github.com/AlanRostem/mu-8/mu8"
	"github.com/AlanRostem/mu-8/system"
)

func Cls(args []mu8.DByte, sys *system.System) {
	pcDebugf(sys.Registers.PC(), "CLS")
	for y := range system.DisplayHeight {
		for x := range system.DisplayWidth {
			sys.FrameBuffer[y][x] = false
		}
	}
}

func DrwVxVyN(args []mu8.DByte, sys *system.System) {
	// TODO implement wrapping
	x := args[0]
	y := args[1]
	n := args[2]
	pcDebugf(sys.Registers.PC(), "DRW V%X, V%X, %d", x, y, n)
	start := sys.Registers.I
	j := mu8.Byte(0)
	for i := start; i < start+n; i++ {
		addr := mu8.NewUint12(int(i))
		rowByte := sys.Memory.Read(addr)
		row := rowByte.BoolArray()
		for k := range row {
			cx := sys.Registers.V[x] + mu8.Byte(k)
			cy := sys.Registers.V[y] + j
			current := sys.FrameBuffer[cy][cx]
			newVal := mu8.BoolXor(current, row[k])
			sys.FrameBuffer[cy][cx] = newVal
		}
		j++
	}
}
