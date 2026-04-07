package keyboard

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

const KeyCount = 0x10

type Keyboard struct {
	keys [KeyCount]bool
}

func New() *Keyboard {
	return &Keyboard{
		keys: [KeyCount]bool{},
	}
}

func (kb *Keyboard) Get(key uint8) bool {
	if key >= KeyCount {
		panic("key does not exist")
	}
	return kb.keys[key]
}

func (kb *Keyboard) checkPressed(keyIdx uint8, ebKey ebiten.Key) {
	if inpututil.IsKeyJustPressed(ebKey) {
		kb.keys[keyIdx] = true
	} else if inpututil.IsKeyJustReleased(ebKey) {
		kb.keys[keyIdx] = false
	}
}

func (kb *Keyboard) Update() {
	// first row
	kb.checkPressed(0x0, ebiten.Key1)
	kb.checkPressed(0x1, ebiten.Key2)
	kb.checkPressed(0x2, ebiten.Key3)
	kb.checkPressed(0x3, ebiten.Key4)
	// second row
	kb.checkPressed(0x4, ebiten.KeyQ)
	kb.checkPressed(0x5, ebiten.KeyW)
	kb.checkPressed(0x6, ebiten.KeyE)
	kb.checkPressed(0x7, ebiten.KeyR)
	// third row
	kb.checkPressed(0x8, ebiten.KeyA)
	kb.checkPressed(0x9, ebiten.KeyS)
	kb.checkPressed(0xA, ebiten.KeyD)
	kb.checkPressed(0xB, ebiten.KeyF)
	// last row
	kb.checkPressed(0xC, ebiten.KeyZ)
	kb.checkPressed(0xD, ebiten.KeyX)
	kb.checkPressed(0xE, ebiten.KeyC)
	kb.checkPressed(0xF, ebiten.KeyV)
}
