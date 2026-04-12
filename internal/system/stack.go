package system

import "github.com/AlanRostem/mu-8/internal/num"

type Stack struct {
	pointer num.Byte
	levels  [stackSize]num.DByte
}

func newStack() *Stack {
	return &Stack{
		pointer: 0,
		levels:  [stackSize]num.DByte{},
	}
}

func (s *Stack) Push(value num.DByte) {
	s.pointer++
	s.levels[s.pointer] = value
}

func (s *Stack) Pop() num.DByte {
	defer func() { s.pointer-- }()
	return s.levels[s.pointer]
}
