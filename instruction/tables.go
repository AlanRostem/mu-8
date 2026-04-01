package instruction

import "github.com/AlanRostem/mu-8/mnemonic"

// TableAll contains instruction maps mapped to a category number.
var TableAll = map[mnemonic.Category]*InstructionMap{
	mnemonic.Category6: NewSingularInstructionMap(ldVxByte),
	mnemonic.CategoryA: NewSingularInstructionMap(ldIAddr),
}
