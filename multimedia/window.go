package multimedia

import (
	"image/color"
	"log"

	"github.com/AlanRostem/mu-8/internal/system"
	"github.com/hajimehoshi/ebiten/v2"
)

type Window struct {
}

func NewWindow() *Window {
	return &Window{}
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
	return nil
}

func (w *Window) Draw(screen *ebiten.Image) {
	screen.Fill(color.Black)
	/*for i := range keyboard.KeyCount {
		x := i % 4
		y := i / 4
		w.system.FrameBuffer[y][x] = w.input.Get(uint8(i))
	}

	for y := range system.DisplayHeight {
		for x := range system.DisplayWidth {
			c := color.Black
			if w.system.FrameBuffer[y][x] {
				c = color.White
			}
			screen.Set(x, y, c)
		}
	}*/
}

func (w *Window) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return system.DisplayWidth, system.DisplayHeight
}
