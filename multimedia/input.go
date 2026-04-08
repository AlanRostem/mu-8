package multimedia

import (
	"github.com/AlanRostem/mu-8/interpreter"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

const KeyCount = 0x10

type UserInput struct {
	interperter *interpreter.Interpreter
}

func NewUserInput(interperter *interpreter.Interpreter) *UserInput {
	return &UserInput{
		interperter: interperter,
	}
}

func (ui *UserInput) updateKeyState(keyIdx uint8, ebKey ebiten.Key) {
	if inpututil.IsKeyJustPressed(ebKey) {
		ui.interperter.SetKey(keyIdx, true)
	} else if inpututil.IsKeyJustReleased(ebKey) {
		ui.interperter.SetKey(keyIdx, false)
	}
}

func (ui *UserInput) Update() {
	// first row
	ui.updateKeyState(0x0, ebiten.Key1)
	ui.updateKeyState(0x1, ebiten.Key2)
	ui.updateKeyState(0x2, ebiten.Key3)
	ui.updateKeyState(0x3, ebiten.Key4)
	// second row
	ui.updateKeyState(0x4, ebiten.KeyQ)
	ui.updateKeyState(0x5, ebiten.KeyW)
	ui.updateKeyState(0x6, ebiten.KeyE)
	ui.updateKeyState(0x7, ebiten.KeyR)
	// third row
	ui.updateKeyState(0x8, ebiten.KeyA)
	ui.updateKeyState(0x9, ebiten.KeyS)
	ui.updateKeyState(0xA, ebiten.KeyD)
	ui.updateKeyState(0xB, ebiten.KeyF)
	// last row
	ui.updateKeyState(0xC, ebiten.KeyZ)
	ui.updateKeyState(0xD, ebiten.KeyX)
	ui.updateKeyState(0xE, ebiten.KeyC)
	ui.updateKeyState(0xF, ebiten.KeyV)
}
