package cli

import (
	"github.com/AlanRostem/mu-8/display"
	"github.com/AlanRostem/mu-8/instruction"
	"github.com/AlanRostem/mu-8/system"
)

func Run() {
	sys := system.New()
	exec := instruction.NewExecutor(sys)
	exec.Exec(0xE2A1)
	exec.Exec(0xE59E)
	w := display.NewWindow(sys)
	sys.FrameBuffer[5][10] = true
	w.Run()
}
