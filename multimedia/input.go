package multimedia

import (
	"github.com/AlanRostem/mu-8/internal/system"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

const KeyCount = 0x10

type UserInput struct {
	sys *system.System
}

func NewUserInput(sys *system.System) *UserInput {
	return &UserInput{
		sys: sys,
	}
}

func (ui *UserInput) Get(key uint8) bool {
	if key >= KeyCount {
		panic("key does not exist")
	}
	return ui.sys.Keys[key]
}

func (ui *UserInput) checkPressed(keyIdx uint8, ebKey ebiten.Key) {
	if inpututil.IsKeyJustPressed(ebKey) {
		ui.sys.Keys[keyIdx] = true
	} else if inpututil.IsKeyJustReleased(ebKey) {
		ui.sys.Keys[keyIdx] = false
	}
}

func (ui *UserInput) Update() {
	// first row
	ui.checkPressed(0x0, ebiten.Key1)
	ui.checkPressed(0x1, ebiten.Key2)
	ui.checkPressed(0x2, ebiten.Key3)
	ui.checkPressed(0x3, ebiten.Key4)
	// second row
	ui.checkPressed(0x4, ebiten.KeyQ)
	ui.checkPressed(0x5, ebiten.KeyW)
	ui.checkPressed(0x6, ebiten.KeyE)
	ui.checkPressed(0x7, ebiten.KeyR)
	// third row
	ui.checkPressed(0x8, ebiten.KeyA)
	ui.checkPressed(0x9, ebiten.KeyS)
	ui.checkPressed(0xA, ebiten.KeyD)
	ui.checkPressed(0xB, ebiten.KeyF)
	// last row
	ui.checkPressed(0xC, ebiten.KeyZ)
	ui.checkPressed(0xD, ebiten.KeyX)
	ui.checkPressed(0xE, ebiten.KeyC)
	ui.checkPressed(0xF, ebiten.KeyV)
}
