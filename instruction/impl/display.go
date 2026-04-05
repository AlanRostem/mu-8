package impl

import (
	"github.com/AlanRostem/mu-8/mu8"
	"github.com/AlanRostem/mu-8/system"
)

func Cls(args []mu8.DByte, sys *system.System) {
	for y := range system.DisplayHeight {
		for x := range system.DisplayWidth {
			sys.FrameBuffer[y][x] = false
		}
	}
}

func DrwVxVyN(args []mu8.DByte, sys *system.System) {
	x := args[0]
	y := args[1]
	n := args[2]
	start := sys.Registers.Index
	j := mu8.Byte(0)
	for i := start; i < start+n; i++ {
		addr := mu8.NewUint12(uint(i))
		rowByte := sys.Memory.Read(addr)
		row := rowByte.BoolArray()
		for k := range row {
			cx := sys.Registers.GeneralPurpose[x] + mu8.Byte(k)
			cy := sys.Registers.GeneralPurpose[y] + j
			current := sys.FrameBuffer[cy][cx]
			newVal := mu8.BoolXor(current, row[k])
			sys.FrameBuffer[cy][cx] = newVal
		}
		j++
	}
}
