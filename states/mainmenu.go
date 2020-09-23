package states

import (
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

type StateMainMenu struct {
	State
}

func (s StateMainMenu) Update(dt float64) {

}

func (s StateMainMenu) Render(win pixelgl.Window) {
	win.Clear(colornames.Darksalmon)
}

func (s StateMainMenu) Start() {

}
