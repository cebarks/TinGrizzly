package states

import (
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

// StateResource is a testing state for any dev work
type StateResource struct {
}

func (s StateResource) Update(sc *StateContext, dt float64) {
	if sc.WindowManager.Pressed(pixelgl.Key1) {
		sc.StateManager.SetState("null")
	}
}

func (s *StateResource) Render(win *pixelgl.Window) {
	win.Clear(colornames.Aqua)
}

func (s *StateResource) Start() {

}

func (s StateResource) Stop() {
}
