package processor

import (
	"github.com/AlanRostem/mu-8/internal/num"
	"github.com/AlanRostem/mu-8/internal/system"
)

type Instruction func(args []num.DByte, sys *system.System)
