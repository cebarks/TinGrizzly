package states

import (
	"github.com/faiface/pixel/pixelgl"
)

// State implements the base interface/struct for  any State
type State interface {
	Update(dt float64)
	Render(win *pixelgl.Window)
	Start()
	Pause()
	Stop()
}
