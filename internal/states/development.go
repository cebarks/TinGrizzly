package states

import (
	"log"

	"github.com/cebarks/TinGrizzly/internal/world"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

// StateDevelopment is a testing state for any dev work
type StateDevelopment struct {
	State
}

func (s StateDevelopment) Update(wm *StateContext, dt float64) {
	if wm.WindowManager.Pressed(pixelgl.Key1) {
		wm.StateManager.SetState("null")
	}
}

func (s StateDevelopment) Render(win *pixelgl.Window) {
	win.Clear(colornames.Darksalmon)
}

func (s StateDevelopment) Start() {
	w := world.NewWorld(15, 15)

	t, _ := w.Grid.At(3, 3)

	td := w.TileDataFromTile(t)

	td.Type = world.TileTypeStone
	td.State.AddFlag(world.FlagActive)
	// td.State.AddFlag(world.FlagStateful)

	log.Printf("Tile: %+v", t)
	log.Printf("TileData: %+v", td)
	log.Printf("bitmask: %b", td.State.TileBitmask)

	t2, _ := w.Grid.At(3, 3)

	td2 := w.TileDataFromTile(t2)

	log.Printf("Tile: %+v", t2)
	log.Printf("TileData: %+v", td2)
	log.Printf("bitmask: %b", td.State.TileBitmask)
}

func (s StateDevelopment) Stop() {

}

// StateNull is a testing state for any dev work
type StateNull struct {
	State
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

}

func (s StateNull) Stop() {

}
