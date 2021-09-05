package states

import (
	"math/rand"

	"github.com/cebarks/TinGrizzly/internal/world"
	"github.com/cebarks/TinGrizzly/resources"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/kelindar/tile"
	"github.com/rs/zerolog/log"
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
	camera *tile.View
}

func (s StateDevelopment) Update(sc *StateContext, dt float64) error {
	s.wm.Update(dt)
	return nil
}

func (s *StateDevelopment) Render(win *pixelgl.Window) error {
	win.Clear(colornames.Aqua)
	c := canvasFromView(s.world, s.camera)
	c.Draw(win, pixel.IM.Scaled(pixel.ZV, 4).Moved(win.Bounds().Center()))
	return nil
}

func (s *StateDevelopment) Start() {
	s.world = world.NewWorld(36, 36)
	s.wm = world.NewManager(s.world)
	s.wm.AddSystem(devSystem)
	s.camera = s.world.Grid.View(tile.NewRect(0, 0, 9, 9), nil)

	for i := 0; i < 400; i++ {

		s.world.TileDataLookup(int16(rand.Intn(36)), int16(rand.Intn(36)))
	}
}

func (s StateDevelopment) Stop() {
}

func canvasFromView(w *world.World, view *tile.View) *pixelgl.Canvas {
	canvas := pixelgl.NewCanvas(pixel.R(0, 0, 9*16, 9*16))

	if view == nil {
		log.Panic().Msg("view is nil")
	}

	view.Each(func(p tile.Point, t tile.Tile) {
		td := w.TileDataLookupFromTile(t)

		tt, err := td.State.Get("tile_type")
		if err != nil {
			log.Panic().Err(err).Msg("")
			return
		}

		typ := tt.(world.TileType)

		var sprite *pixel.Sprite

		if typ == world.TileTypeStone {
			sprite = resources.GetSprite("stone")
		} else if typ == world.TileTypeEmpty {
			sprite = resources.GetSprite("grass")
		}
		sprite.Draw(canvas, pixel.IM.Moved(pixel.V(8, 8)).Moved(pixel.V(float64(p.X*16), float64(p.Y*16))))
	})

	return canvas
}
