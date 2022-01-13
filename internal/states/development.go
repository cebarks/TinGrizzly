package states

import (
	"math/rand"

	"github.com/cebarks/TinGrizzly/internal/gfx"
	"github.com/cebarks/TinGrizzly/internal/world"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/kelindar/tile"
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
	world  *world.World
	wm     *world.Manager
	camera *gfx.WorldCamera
}

func (s StateDevelopment) Update(sc *StateContext, dt float64) error {
	s.wm.Update(dt)

	s.camera.Update(s.world)
	return nil
}

func (s *StateDevelopment) Render(win *pixelgl.Window) error {
	win.Clear(colornames.Aqua)
	s.camera.Canvas.Draw(win, pixel.IM.Scaled(pixel.ZV, 3).Moved(win.Bounds().Center()))
	return nil
}

func (s *StateDevelopment) Start() {
	s.world = world.NewWorld(36, 36)
	s.wm = world.NewManager(s.world)
	s.wm.AddSystem(devSystem)
	s.camera = gfx.NewCamera(s.world.Grid.View(tile.NewRect(0, 0, 21, 21), nil))

	for i := 0; i < 400; i++ {
		s.world.TileDataLookup(int16(rand.Intn(36)), int16(rand.Intn(36)))
	}
	s.world.Entities = append(s.world.Entities)
}

func (s StateDevelopment) Stop() {
}

func NewPlayer() *world.Entity {
	player := world.NewEntity()

	player.State.Set("sprite", "sword")
	player.State.Set("location", &tile.Point{X: 0, Y: 0})

	return player
}
