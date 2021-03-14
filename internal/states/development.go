package states

import (
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

// StateDevelopment is a testing state for any dev work
type StateDevelopment struct {
	State
}

func (s StateDevelopment) Update(dt float64) {

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

func (s StateNull) Update(dt float64) {

}

func (s StateNull) Render(win *pixelgl.Window) {
	win.Clear(colornames.Paleturquoise)
}

func (s StateNull) Start() {

}

func (s StateNull) Stop() {

}
