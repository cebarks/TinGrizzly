package states

import (
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

// StateMainMenu is the main menu
type StateMainMenu struct {
}

func (s StateMainMenu) Update(wm *StateContext, dt float64) error {

	return nil
}

func (s StateMainMenu) Render(win *pixelgl.Window) error {
	win.Clear(colornames.Darksalmon)
	return nil
}

func (s StateMainMenu) Start() {

}

func (s StateMainMenu) Stop() {

}
