package system

import (
	"sync"
	"time"

	"github.com/AlanRostem/mu-8/internal/num"
)

type Timers struct {
	delay num.Byte
	sound num.Byte

	mut sync.Mutex
}

func newTimers() *Timers {
	return &Timers{}
}

func (t *Timers) Run() {
	for {
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
