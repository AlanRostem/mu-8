package main

import (
	"fmt"

	"github.com/AlanRostem/mu-8/cpu"
)

func main() {
	regFile := cpu.NewRegisterFile()
	regFile.CurrentOperand = 0x1
	fmt.Println("The current operand is", regFile.CurrentOperand)
}
