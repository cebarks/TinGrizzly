package states

import (
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

// StateDevelopment is a testing state for any dev work
type StateDevelopment struct {
	State
}

func (s StateDevelopment) Update(wm *StateContext, dt float64) {
	if wm.WindowManager.Pressed(pixelgl.Key1) {
		wm.StateManager.SetState("null")
	}
}

func (s StateDevelopment) Render(win *pixelgl.Window) {
	win.Clear(colornames.Darksalmon)
}

func (s StateDevelopment) Start() {

}

func (s StateDevelopment) Stop() {

}

// StateNull is a testing state for any dev work
type StateNull struct {
	State
}

func (s StateNull) Update(wm *StateContext, dt float64) {
	if wm.WindowManager.Pressed(pixelgl.Key2) {
		wm.StateManager.SetState("dev")
	}
}

func (s StateNull) Render(win *pixelgl.Window) {
	win.Clear(colornames.Paleturquoise)
}

func (s StateNull) Start() {

}

func (s StateNull) Stop() {

}
