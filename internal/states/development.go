package states

import (
	"fmt"

	"github.com/cebarks/TinGrizzly/internal/util/asset"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

// StateDevelopment is a testing state for any dev work
type StateDevelopment struct {
	batch *pixel.Batch
}

func (s StateDevelopment) Update(sc *StateContext, dt float64) {
	if sc.WindowManager.Pressed(pixelgl.Key1) {
		sc.StateManager.SetState("null")
	}
}

func (s *StateDevelopment) Render(win *pixelgl.Window) {
	win.Clear(colornames.Aqua)
	mat := pixel.IM.Moved(pixel.V(100, 100))
	for i := 0; i < 12; i++ {
		for j := 0; j < 43; j++ {
			asset.Sprites[fmt.Sprintf("%v:%v", i, j)].Draw(win, mat.Moved(pixel.V(float64(i)*16, float64(j)*16)))
		}
	}
}

func (s *StateDevelopment) Start() {
	// s.batch = pixel.NewBatch(&pixel.TrianglesData{}, )
}

func (s StateDevelopment) Stop() {
}

// StateNull is a testing state for any dev work
type StateNull struct {
}

func (s StateNull) Update(sc *StateContext, dt float64) {
	if sc.WindowManager.Pressed(pixelgl.Key2) {
		sc.StateManager.SetState("dev")
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
