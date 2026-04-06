package system

import "github.com/AlanRostem/mu-8/mu8"

const MemorySize = mu8.Uint12Max

type MemoryBank struct {
	data [MemorySize]mu8.Byte
}

func newMemoryBank() *MemoryBank {
	m := &MemoryBank{
		data: [MemorySize]mu8.Byte{},
	}
	// copy font data to memory
	copy(m.data[:], fontData[:])
	return m
}

func (m *MemoryBank) Write(addr mu8.Uint12, value mu8.Byte) {
	m.data[addr.Value()] = value
}

func (m *MemoryBank) FetchInstruction(addr mu8.Uint12) mu8.DByte {
	left := m.Read(addr)
	addr.Add(1)
	right := m.Read(addr)
	opcode := left.Concat(right)
	return opcode
}

func (m *MemoryBank) Read(addr mu8.Uint12) mu8.Byte {
	return m.data[addr.Value()]
}
