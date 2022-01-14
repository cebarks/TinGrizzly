package states

import (
	"github.com/cebarks/TinGrizzly/internal/world"
	"github.com/faiface/pixel/pixelgl"
	"github.com/kelindar/tile"
	"golang.org/x/image/colornames"
)

// StateDevelopment is a testing state for any dev work
type StateWorld struct {
	w *world.World
}

func (s StateWorld) Update(sc *StateContext, dt float64) {
	s.w.Update(dt)
}

func (s *StateWorld) Render(win *pixelgl.Window) {
	win.Clear(colornames.Darksalmon)
	// s.w.Render(win)
}

func (s *StateWorld) Start() {
	s.w = world.NewWorld(39, 39)

	s.w.Grid.Each(func(p tile.Point, t tile.Tile) {
		if (p.X%2 == 0 && p.Y%2 == 0) || (p.X%2 == 1 && p.Y%2 == 1) {
			return
		}
		s.w.SetTileTo(p, "stone")
	})

	var i int16
	var j int16
	for i = 0; i < 13; i++ {
		for j = 0; j < 13; j++ {
			td := s.w.TileDataLookup(i, j)
			td.Header.Bitmask.ClearFlag(world.FlagActive)
		}
	}
}

func (s StateWorld) Stop() {
}
