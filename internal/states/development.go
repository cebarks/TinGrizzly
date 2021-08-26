package states

import (
	"github.com/cebarks/TinGrizzly/internal/world"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

var devSystem *world.System = &world.System{
	Priority:         1,
	Id:               "dev",
	Type:             world.SystemTypeEntity | world.SystemTypeTile,
	Concurrent:       false,
	TileUpdateFunc:   tileUpdate,
	EntityUpdateFunc: entityUpdate,
}

func tileUpdate(sys *world.System, targets *world.TileData, delta float64) error {
	return nil
}

func entityUpdate(sys *world.System, targets *world.Entity, delta float64) error {
	return nil
}

// StateDevelopment is a testing state for any dev work
type StateDevelopment struct {
	world *world.World
	wm    *world.Manager
}

func (s StateDevelopment) Update(sc *StateContext, dt float64) error {
	s.wm.Update(dt)
	return nil
}

func (s *StateDevelopment) Render(win *pixelgl.Window) error {
	win.Clear(colornames.Aqua)
	return nil
}

func (s *StateDevelopment) Start() {
	s.world = world.NewWorld(36, 36)
	s.wm = world.NewManager(s.world)
	s.wm.AddSystem(devSystem)
}

func (s StateDevelopment) Stop() {
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
	// log.Error().Err(fmt.Errorf("test: %v", []int{1, 2, 3})).Msg("Test Error")
}

func (s StateNull) Stop() {

}
