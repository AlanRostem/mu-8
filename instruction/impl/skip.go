package impl

import (
	"fmt"

	"github.com/AlanRostem/mu-8/mu8"
	"github.com/AlanRostem/mu-8/system"
)

func SkpVx(args []mu8.DByte, sys *system.System) {
	x := args[0]
	// TODO implement
	fmt.Printf("SKP V%X\n", x)
}

func SknpVx(args []mu8.DByte, sys *system.System) {
	x := args[0]
	// TODO implement
	fmt.Printf("SKNP V%X\n", x)
}
