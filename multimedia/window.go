package multimedia

import (
	"image/color"
	"log"

	"github.com/AlanRostem/mu-8/internal/system"
	"github.com/AlanRostem/mu-8/interpreter"
	"github.com/hajimehoshi/ebiten/v2"
)

type Window struct {
	interpreter *interpreter.Interpreter
	input       *UserInput
	sound       *SoundEngine
}

func NewWindow(interpreter *interpreter.Interpreter) *Window {
	return &Window{
		interpreter: interpreter,
		input:       NewUserInput(interpreter),
		sound:       NewSoundEngine(interpreter),
	}
}

func (w *Window) Run() {
	ebiten.SetWindowSize(system.DisplayWidth*10, system.DisplayHeight*10)
	ebiten.SetWindowTitle("MU-8")
	ebiten.SetTPS(60)
	if err := ebiten.RunGame(w); err != nil {
		log.Fatal(err)
	}
}

func (w *Window) Update() error {
	w.input.Update()
	w.sound.Update()
	return nil
}

func (w *Window) Draw(screen *ebiten.Image) {
	screen.Fill(color.Black)
	frame := w.interpreter.DisplayBuffer()
	for y := range system.DisplayHeight {
		for x := range system.DisplayWidth {
			c := color.Black
			if frame[y][x] {
				c = color.White
			}
			screen.Set(x, y, c)
		}
	}
}

func (w *Window) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return system.DisplayWidth, system.DisplayHeight
}
