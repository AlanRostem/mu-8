package cli

import (
	"log"
	"os"

	"github.com/AlanRostem/mu-8/interpreter"
	"github.com/AlanRostem/mu-8/multimedia"
)

/*var sprite = []mu8.Byte{
	0b00100100,
	0b00000000,
	0b10000001,
	0b01000010,
	0b00111100,
}*/

/*var forLoopProgram = []mu8.Byte{
	0x65, 0x00, // 0x200: LD V5, 0
	0x75, 0x01, // 0x202: ADD V5, 1
	0x35, 0x05, // 0x204: SE V5, 5
	0x12, 0x00, // 0x206: JP 0x200
	0xFF, 0xFF, // 0x208: EXIT (custom)
}*/

var keySkipProgram = []byte{
	0x65, 0x00, // 0x200: LD V5, 0
	0xE5, 0x9E, // 0x202: SKP V5
	0x12, 0x02, // 0x204: JP 0x202
	0xFF, 0xFF, // 0x206: EXIT (custom)
}

var delayTimerProgram = []byte{
	0x65, 0xF0, // 0x200: LD V5, 240
	0xF5, 0x15, // 0x202: LD DT, V5
	0xF5, 0x07, // 0x204: LD V5, DT
	0x35, 0x00, // 0x206: SE V5, 0
	0x12, 0x04, // 0x208: JP 0x204
	0xFF, 0xFF, // 0x20A: EXIT (custom)
}

var soundTimerProgram = []byte{
	0x65, 0x40, // 0x200: LD V5, 64
	0xF5, 0x18, // 0x202: LD ST, V5
	0xFF, 0xFF, // 0x20A: EXIT (custom)
}

func Run() {
	const romPath = "programs/3-corax+.ch8"
	// const romPath = "programs/4-flags.ch8"
	// const romPath = "programs/Breakout (Brix hack) [David Winter, 1997].ch8"
	data, err := os.ReadFile(romPath)
	if err != nil {
		log.Fatal(err)
	}
	in := interpreter.New()
	in.Load(data)
	in.Run()
	w := multimedia.NewWindow(in)
	w.Run()
	in.Stop()
}
