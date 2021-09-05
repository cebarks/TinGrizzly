package states

import (
	"image/color"

	"github.com/cebarks/TinGrizzly/resources"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
	"github.com/rs/zerolog/log"
	"golang.org/x/image/colornames"
)

type StateLoading struct {
	text *text.Text
}

func (s StateLoading) Update(wm *StateContext, dt float64) error {
	resources.Setup()
	wm.StateManager.SetState("dev")
	log.Logger.Info().Msg("Done Loading!")
	return nil
}

func (s StateLoading) Render(win *pixelgl.Window) error {
	win.Clear(colornames.Black)
	s.text.WriteString("Loading...") //TODO: doesn't work?
	return nil
}

func (s *StateLoading) Start() {
	s.text = text.New(pixel.V(100, 100), text.Atlas7x13)
	s.text.Color = color.White
}

func (s StateLoading) Stop() {

}

// StateNull is a testing state for any dev work
type StateNull struct {
}

func (s StateNull) Update(sc *StateContext, dt float64) error {
	return nil
}

func (s StateNull) Render(win *pixelgl.Window) error {
	win.Clear(colornames.Paleturquoise)
	return nil
}

func (s StateNull) Start() {

}

func (s StateNull) Stop() {

}
