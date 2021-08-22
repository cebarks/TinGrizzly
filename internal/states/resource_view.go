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

func (s *StateResource) Update(sc *StateContext, dt float64) error {
	return nil
}

func (s *StateResource) Render(win *pixelgl.Window) error {
	win.Clear(colornames.Brown)
	resources.Sheet.Packr.Draw(win, pixel.IM.Moved(resources.Sheet.Packr.Center()))
	return nil
}

func (s *StateResource) Start() {
}

func (s StateResource) Stop() {
}
