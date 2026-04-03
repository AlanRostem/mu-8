package cli

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

type Window struct {
	pixels [32][64]bool
}

func NewWindow() *Window {
	return &Window{}
}

func (w *Window) SetPixel(x, y uint8, black bool) {
	w.pixels[y][x] = black
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
	screen.Fill(color.White)
	for y := range 32 {
		for x := range 64 {
			c := color.White
			if w.pixels[y][x] {
				c = color.Black
			}
			screen.Set(x, y, c)
		}
	}
}

func (w *Window) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 64, 32
}
