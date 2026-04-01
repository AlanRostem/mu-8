package cli

import (
	"fmt"

	"github.com/AlanRostem/mu-8/cpu"
	"github.com/AlanRostem/mu-8/memory"
	"github.com/AlanRostem/mu-8/mu8"
)

func Run() {
	regFile := cpu.NewRegisterFile()
	regFile.CurrentOperand = 0x1
	fmt.Println("The current operand is", regFile.CurrentOperand)
	addr := mu8.NewUint12(0x0FFF)
	mem := memory.New()
	mem.Write(addr, 123)
	fmt.Printf("Address %s has value %d\n", &addr, mem.Read(addr))
	db := mu8.DByte(0xABCD)
	fmt.Printf("N0=%X\n", db.Nibble(0))
	fmt.Printf("N1=%X\n", db.Nibble(1))
	fmt.Printf("N2=%X\n", db.Nibble(2))
	fmt.Printf("N3=%X\n", db.Nibble(3))
}
