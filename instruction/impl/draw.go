package impl

import (
	"github.com/AlanRostem/mu-8/mu8"
	"github.com/AlanRostem/mu-8/system"
)

func boolXor(a, b bool) bool {
	return a != b
}

func DrwVxVyN(args []mu8.DByte, sys *system.System) {
	x := args[0]
	y := args[1]
	n := args[2]
	start := sys.Registers.Index
	j := mu8.DByte(0)
	for i := start; i < start+n; i++ {
		addr := mu8.NewUint12(uint(i))
		rowByte := sys.Memory.Read(addr)
		row := rowByte.BoolArray()
		for k := range row {
			cx := x + mu8.DByte(k)
			cy := y + j
			current := sys.FrameBuffer[cy][cx]
			newVal := boolXor(current, row[k])
			sys.FrameBuffer[cy][cx] = newVal
		}
		j++
	}
}
