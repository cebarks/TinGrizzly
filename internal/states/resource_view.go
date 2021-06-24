package states

import (
	"github.com/cebarks/TinGrizzly/resources"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

// StateResource is a testing state for any dev work
type StateResource struct {
}

func (s *StateResource) Update(sc *StateContext, dt float64) {
}

func (s *StateResource) Render(win *pixelgl.Window) {
	win.Clear(colornames.Brown)
	resources.Sheet.Packr.Draw(win, pixel.IM.Moved(resources.Sheet.Packr.Center()))
}

func (s *StateResource) Start() {
}

func (s StateResource) Stop() {
}
