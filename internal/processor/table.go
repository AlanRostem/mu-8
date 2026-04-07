package processor

import "github.com/AlanRostem/mu-8/internal/num"

const singleInstKey = 0xFF

type instructionTable map[num.DByte]Instruction

func newInstructionTable() instructionTable {
	return make(map[num.DByte]Instruction)
}

func newSingularInstructionTable(inst Instruction) instructionTable {
	return map[num.DByte]Instruction{singleInstKey: inst}
}

func (im instructionTable) IsSingle() bool {
	_, ok := im[singleInstKey]
	return ok
}

func (im instructionTable) Single() Instruction {
	if !im.IsSingle() {
		panic("cannot get categorized instruction when map is singluar")
	}
	return im[singleInstKey]
}

func (im instructionTable) Get(id num.DByte) Instruction {
	if im.IsSingle() {
		panic("cannot get multiple instructions for singular map")
	}
	return im[id]
}

func (im instructionTable) Add(id num.DByte, inst Instruction) {
	if im.IsSingle() {
		panic("cannot add more than one instruction to a singular map")
	}
	im[id] = inst
}
