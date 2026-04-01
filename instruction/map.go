package instruction

type InstructionMap struct {
	instructions map[uint8]Instruction
	isSingular   bool
}

func NewInstructionMap() *InstructionMap {
	return &InstructionMap{
		instructions: make(map[uint8]Instruction),
		isSingular:   false,
	}
}

func NewSingularInstructionMap(inst Instruction) *InstructionMap {
	return &InstructionMap{
		instructions: map[uint8]Instruction{0: inst},
		isSingular:   true,
	}
}

func (im *InstructionMap) Single() Instruction {
	if !im.isSingular {
		panic("cannot get indexed instruction when map is singluar")
	}
	return im.instructions[0]
}

func (im *InstructionMap) Get(subcat uint8) Instruction {
	if im.isSingular {
		panic("cannot get multiple instructions for singular map")
	}
	return im.instructions[subcat]
}
