package states

import (
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

// StateMainMenu is the main menu
type StateMainMenu struct {
}

func (s StateMainMenu) Update(wm *StateContext, dt float64) {

}

func (s StateMainMenu) Render(win *pixelgl.Window) {
	win.Clear(colornames.Darksalmon)
}

func (s StateMainMenu) Start() {

}

func (s StateMainMenu) Stop() {

}
