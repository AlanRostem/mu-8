package memory

import "github.com/AlanRostem/mu-8/mu8"

const MemorySize = mu8.Uint12Max

type Memory struct {
	data [MemorySize]mu8.Byte
}

func New() *Memory {
	return &Memory{
		data: [MemorySize]mu8.Byte{},
	}
}

func (m *Memory) Write(addr mu8.Uint12, value mu8.Byte) {
	m.data[addr.Value()] = value
}

func (m *Memory) Read(addr mu8.Uint12) mu8.Byte {
	return m.data[addr.Value()]
}
