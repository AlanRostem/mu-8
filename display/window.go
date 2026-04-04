package display

import (
	"image/color"
	"log"

	"github.com/AlanRostem/mu-8/system"
	"github.com/hajimehoshi/ebiten/v2"
)

type Window struct {
	system *system.System
}

func NewWindow(system *system.System) *Window {
	return &Window{system: system}
}

func (w *Window) Run() {
	ebiten.SetWindowSize(640, 320)
	ebiten.SetWindowTitle("MU-8")
	if err := ebiten.RunGame(w); err != nil {
		log.Fatal(err)
	}
}

func (w *Window) Update() error {
	return nil
}

func (w *Window) Draw(screen *ebiten.Image) {
	screen.Fill(color.Black)
	for y := range 32 {
		for x := range 64 {
			c := color.Black
			if w.system.FrameBuffer[y][x] {
				c = color.White
			}
			screen.Set(x, y, c)
		}
	}
}

func (w *Window) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 64, 32
}
