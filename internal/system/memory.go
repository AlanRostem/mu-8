package system

import "github.com/AlanRostem/mu-8/internal/num"

const MemorySize = num.Uint12Max

type MemoryBank struct {
	data [MemorySize]num.Byte
}

func newMemoryBank() *MemoryBank {
	m := &MemoryBank{
		data: [MemorySize]num.Byte{},
	}
	// copy font data to memory
	copy(m.data[:], fontData[:])
	return m
}

func (m *MemoryBank) Write(addr num.Uint12, value num.Byte) {
	m.data[addr.Value()] = value
}

func (m *MemoryBank) FetchInstruction(addr num.Uint12) num.DByte {
	left := m.Read(addr)
	addr.Add(1)
	right := m.Read(addr)
	opcode := left.Concat(right)
	return opcode
}

func (m *MemoryBank) Read(addr num.Uint12) num.Byte {
	return m.data[addr.Value()]
}
