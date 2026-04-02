package instruction

const singleInstKey = 0xFF

type InstructionMap map[uint8]Instruction

func NewInstructionMap() InstructionMap {
	return make(map[uint8]Instruction)
}

func NewSingularInstructionMap(inst Instruction) InstructionMap {
	return map[uint8]Instruction{singleInstKey: inst}
}

func (im InstructionMap) IsSingle() bool {
	_, ok := im[singleInstKey]
	return ok
}

func (im InstructionMap) Single() Instruction {
	if !im.IsSingle() {
		panic("cannot get categorized instruction when map is singluar")
	}
	return im[singleInstKey]
}

func (im InstructionMap) Get(cat uint8) Instruction {
	if im.IsSingle() {
		panic("cannot get multiple instructions for singular map")
	}
	return im[cat]
}

func (im InstructionMap) Add(subcat uint8, inst Instruction) {
	if im.IsSingle() {
		panic("cannot add more than one instruction to a singular map")
	}
	im[subcat] = inst
}
