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
	ebiten.SetWindowSize(system.DisplayWidth*10, system.DisplayHeight*10)
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
	for y := range system.DisplayHeight {
		for x := range system.DisplayWidth {
			c := color.Black
			if w.system.FrameBuffer[y][x] {
				c = color.White
			}
			screen.Set(x, y, c)
		}
	}
}

func (w *Window) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return system.DisplayWidth, system.DisplayHeight
}
