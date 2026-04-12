package instruction

import (
	"github.com/AlanRostem/mu-8/internal/num"
	"github.com/AlanRostem/mu-8/internal/system"
)

func Cls(args []num.DByte, sys *system.System) {
	pcDebugf(sys.Registers.PC(), "CLS")
	for y := range system.DisplayHeight {
		for x := range system.DisplayWidth {
			sys.FrameBuffer[y][x] = false
		}
	}
}

func DrwVxVyN(args []num.DByte, sys *system.System) {
	x := args[0]
	y := args[1]
	n := args[2]
	pcDebugf(sys.Registers.PC(), "DRW V%X, V%X, %d", x, y, n)
	sys.Registers.V[0xF] = 0
	start := sys.Registers.I
	j := num.Byte(0)
	for i := start; i < start+n; i++ {
		addr := num.NewUint12(int(i))
		rowByte := sys.Memory.Read(addr)
		row := rowByte.BoolArray()
		for k := range row {
			cx := sys.Registers.V[x] + num.Byte(k)
			cy := sys.Registers.V[y] + j
			cx %= system.DisplayWidth
			cy %= system.DisplayHeight
			current := sys.FrameBuffer[cy][cx]
			if current && row[k] {
				sys.Registers.V[0xF] = 1
			}
			newVal := num.BoolXor(current, row[k])
			sys.FrameBuffer[cy][cx] = newVal
		}
		j++
	}
}
