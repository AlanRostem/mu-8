package instruction

import (
	"github.com/AlanRostem/mu-8/num"
	"github.com/AlanRostem/mu-8/system"
)

type Instruction func(args []num.DByte, sys *system.System)
