package cli

import (
	"github.com/AlanRostem/mu-8/display"
	"github.com/AlanRostem/mu-8/instruction"
	"github.com/AlanRostem/mu-8/mu8"
	"github.com/AlanRostem/mu-8/system"
)

var sprite = []mu8.Byte{
	0b00100100,
	0b00000000,
	0b10000001,
	0b01000010,
	0b00111100,
}

func Run() {
	sys := system.New()
	exec := instruction.NewExecutor(sys)
	sys.Registers.Index = 100
	sys.Registers.GeneralPurpose[0x0] = 16
	sys.Registers.GeneralPurpose[0x1] = 8
	for i, row := range sprite {
		addr := mu8.NewUint12(uint(i + sys.Registers.Index.Int()))
		sys.Memory.Write(addr, row)
	}
	exec.Exec(0xD015)
	w := display.NewWindow(sys)
	w.Run()
}
