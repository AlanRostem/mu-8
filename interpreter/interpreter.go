package interpreter

import (
	"sync"
	"time"

	"github.com/AlanRostem/mu-8/internal/logger"
	"github.com/AlanRostem/mu-8/internal/num"
	"github.com/AlanRostem/mu-8/internal/processor"
	"github.com/AlanRostem/mu-8/internal/system"
)

const ProgramSize = system.MemorySize - system.ProgramOffset
const DisplayHeight = system.DisplayHeight
const DisplayWidth = system.DisplayWidth

type Interpreter struct {
	system    *system.System
	processor *processor.Processor

	running bool

	mut sync.RWMutex
}

func New() *Interpreter {
	s := system.New()
	return &Interpreter{
		system:    s,
		processor: processor.New(s),
	}
}

func (in *Interpreter) DisplayBuffer() [DisplayHeight][DisplayWidth]bool {
	in.mut.RLock()
	defer in.mut.RUnlock()
	return in.system.FrameBuffer
}

func (in *Interpreter) SetKey(key uint8, state bool) {
	in.mut.Lock()
	defer in.mut.Unlock()
	in.system.Keys[key] = state
}

func (in *Interpreter) Load(program []byte) {
	in.mut.Lock()
	defer in.mut.Unlock()
	if len(program) > ProgramSize {
		panic("program too large for chip8 memory space")
	}
	in.system.Registers.Clear()
	for i, value := range program {
		addr := num.NewUint12(system.ProgramOffset + i)
		in.system.Memory.Write(addr, num.Byte(value))
	}
}

func (in *Interpreter) Stop() {
	in.mut.Lock()
	defer in.mut.Unlock()
	in.running = false
	in.system.Timers.Stop()
	logger.Debugf("Stopped interpreter.")
}

func (in *Interpreter) Run() {
	go in.RunBlocking()
}

// Run starts the processor cycle "loop". This method is blocking
func (in *Interpreter) RunBlocking() {
	in.system.Timers.Run()
	in.running = true
	for in.running {
		in.processor.Cycle()
		// check for exit opcode when debugging
		addr := num.NewUint12(in.system.Registers.PC().Int())
		opcode := in.system.Memory.FetchInstruction(addr)
		if opcode == 0xFFFF {
			logger.Debugf("Interpreter terminated.")
			in.running = false
			break
		}
		// simulate chip8 "clock speed"
		time.Sleep(time.Second / processor.Frequency)
	}
}
