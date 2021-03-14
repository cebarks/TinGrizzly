package gfx

import (
	"log"

	"github.com/cebarks/TinGrizzly/internal/util"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

type WindowManager struct {
	*pixelgl.Window
}

var (
	title string = ""
)

func BuildWindowManager() *WindowManager {
	window, err := pixelgl.NewWindow(pixelgl.WindowConfig{
		Title:  title,
		Bounds: pixel.R(0, 0, 1920, 1080),
	})

	if util.DebugError(err) {
		log.Fatalf("Couldn't create window: %v", err)
	}

	winm := WindowManager{Window: window}
	return &winm
}
