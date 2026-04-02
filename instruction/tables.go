package instruction

import "github.com/AlanRostem/mu-8/decode"

var MapClass6 = NewSingularInstructionMap(ldVxByte)
var MapClassA = NewSingularInstructionMap(ldIAddr)

// TableAll contains instruction maps mapped to a class number.
var TableAll = map[decode.Class]InstructionMap{
	decode.Class6: MapClass6,
	decode.ClassA: MapClassA,
}
