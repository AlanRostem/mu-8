package instruction

import "github.com/AlanRostem/mu-8/system"

const Count = 35
const MaxNibbles = 4

type Instruction func(nibbles [MaxNibbles]uint8, sys *system.System)
