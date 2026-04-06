package cli

import (
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

var program = [instruction.ProgramSize]mu8.Byte{
	0x61,
	0x11,
	0xFF, // end of program here
	0xFF, // end of program here
}

func Run() {
	sys := system.New()
	exec := instruction.NewExecutor(sys)
	exec.LoadProgram(program)
	exec.ExecProgram()
	// w := display.NewWindow(sys)
	// w.Run()
}
