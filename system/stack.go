package system

import "github.com/AlanRostem/mu-8/mu8"

type Stack struct {
	pointer mu8.Byte
	levels  [stackSize]mu8.DByte
}

func newStack() *Stack {
	return &Stack{
		pointer: 0,
		levels:  [stackSize]mu8.DByte{},
	}
}

func (s *Stack) Push(value mu8.DByte) {
	s.pointer++
	s.levels[s.pointer] = value
}

func (s *Stack) Pop() mu8.DByte {
	s.pointer--
	return s.levels[s.pointer]
}
