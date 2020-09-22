package states

import (
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

// GameState implements the base interface/struct for  any GameState
type GameState struct {
	active bool
}

func (gs GameState) Update(dt float64) {

}

func (gs GameState) Render(win pixelgl.Window) {

}

type GameStateMainMenu struct {
	GameState
}

func (gs GameStateMainMenu) Update(dt float64) {

}

func (gs GameStateMainMenu) Render(win pixelgl.Window) {
	win.Clear(colornames.Darksalmon)
}

func (gs GameStateMainMenu) Start() {

}
