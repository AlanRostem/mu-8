package system

import (
	"sync"
	"time"

	"github.com/AlanRostem/mu-8/internal/logger"
	"github.com/AlanRostem/mu-8/internal/num"
)

type Timers struct {
	delay   num.Byte
	sound   num.Byte
	running bool

	mut sync.RWMutex
}

func newTimers() *Timers {
	return &Timers{}
}

func (t *Timers) Stop() {
	t.mut.Lock()
	defer t.mut.Unlock()
	t.running = false
	logger.Infof("Stopped timers.")
}

func (t *Timers) Run() {
	go t.RunBlocking()
}

func (t *Timers) RunBlocking() {
	t.running = true
	for t.running {
		t.Update()
		time.Sleep(time.Second / 60)
	}
}

func (t *Timers) Update() {
	if t.delay > 0 {
		t.delay--
	}
	if t.sound > 0 {
		t.sound--
	}
}

func (t *Timers) SetDT(value num.Byte) {
	t.mut.Lock()
	defer t.mut.Unlock()
	t.delay = value
}

func (t *Timers) SetST(value num.Byte) {
	t.mut.Lock()
	defer t.mut.Unlock()
	t.sound = value
}

func (t *Timers) DT() num.Byte {
	t.mut.RLock()
	defer t.mut.RUnlock()
	return t.delay
}

func (t *Timers) ST() num.Byte {
	t.mut.RLock()
	defer t.mut.RUnlock()
	return t.sound
}
