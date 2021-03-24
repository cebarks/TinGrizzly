package states

import (
	"github.com/cebarks/TinGrizzly/internal/world"
	"github.com/faiface/pixel/pixelgl"
	"github.com/kelindar/tile"
	"golang.org/x/image/colornames"
)

// StateDevelopment is a testing state for any dev work
type StateDevelopment struct {
	w *world.World
}

func (s StateDevelopment) Update(wm *StateContext, dt float64) {
	if wm.WindowManager.Pressed(pixelgl.Key1) {
		wm.StateManager.SetState("null")
	}
	s.w.Update(dt)
}

func (s *StateDevelopment) Render(win *pixelgl.Window) {
	win.Clear(colornames.Darksalmon)
	s.w.Render(win)
}

func (s *StateDevelopment) Start() {
	s.w = world.NewWorld(39, 39)

	s.w.Grid.Each(func(p tile.Point, t tile.Tile) {
		if (p.X%2 == 0 && p.Y%2 == 0) || (p.X%2 == 1 && p.Y%2 == 1) {
			return
		}
		s.w.SetTileTo(p, world.TileTypeStone)
	})

	var i int16
	var j int16
	for i = 0; i < 13; i++ {
		for j = 0; j < 13; j++ {
			td := s.w.TileDataLookup(i, j)
			td.Header.Bitmask.RemoveFlag(world.FlagActive)
		}
	}
}

func (s StateDevelopment) Stop() {
	s.w = nil
}

// StateNull is a testing state for any dev work
type StateNull struct {
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
	// log.Error().Err(fmt.Errorf("test: %v", []int{1, 2, 3})).Msg("Test Error")
}

func (s StateNull) Stop() {

}
