package cli

import (
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

/* var forLoopProgram = [instruction.ProgramSize]mu8.Byte{
	0x65, 0x00, // LD V5, 0
	0x75, 0x01, // ADD V5, 1
	0x35, 0x05, // SE V5, 2
	0x12, 0x00, // JP 0x200
	0xFF, 0xFF, // end of program here using custom opcode
}*/

var keyHaltProgram = [interpreter.ProgramSize]byte{
	0xF5, 0x0A, // LD V5, K
	0x12, 0x02, // JP 0x201
}

func Run() {
	// const romPath = "programs/IBM Logo.ch8"
	// data, err := os.ReadFile(romPath)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	in := interpreter.New()
	in.Load(keyHaltProgram[:])
	go in.Run()
	w := multimedia.NewWindow(in)
	w.Run()
}
