package memory

import "github.com/AlanRostem/mu-8/mu8"

const MemorySize = 0x1000

type Memory struct {
	data [MemorySize]mu8.Byte
}

func NewMemory() *Memory {
	return &Memory{
		data: [MemorySize]mu8.Byte{},
	}
}

func (m *Memory) Write(addr mu8.Address, value mu8.Byte) {
	m.data[addr.Value()] = value
}

func (m *Memory) Read(addr mu8.Address) mu8.Byte {
	return m.data[addr.Value()]
}
