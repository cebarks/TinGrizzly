package gfx

import (
	"fmt"
	"time"

	"github.com/cebarks/TinGrizzly/util"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/rs/zerolog/log"
	"golang.org/x/image/colornames"
)

// WindowManager is a pixelgl.Window wrapper responsible for all window related logic and control.
type WindowManager struct {
	win          *pixelgl.Window
	winCfg       pixelgl.WindowConfig
	windowVsync  bool
	windowBounds pixel.Rect
	windowTitle  string
	shouldClose  bool

	// fps vars
	frames  int
	lastFps int
}

var second = time.Tick(time.Second)

// Initialize sets var defaults
func (wm *WindowManager) Initialize() {
	wm.windowTitle = "Tin Grizzly"
	wm.windowBounds = pixel.R(0, 0, 1600, 900)
	wm.windowVsync = false
	wm.frames = 0
}

// CreateWindow creates the game window  with config options
func (wm *WindowManager) CreateWindow() {
	wm.winCfg = pixelgl.WindowConfig{
		Title:  wm.windowTitle,
		Bounds: wm.windowBounds,
		VSync:  wm.windowVsync,
	}

	log.Debug().Msgf("Created window config: %d", wm.winCfg)
	log.

	win, err := pixelgl.NewWindow(wm.winCfg)

	if util.DebugError(err) {
		panic(err)
	} else {
		wm.win = win
	}

	wm.win.Clear(colornames.Darkslateblue)
}

// Close sets shouldClose in WindowManager to true
func (wm *WindowManager) Close() {
	wm.shouldClose = true
}

// Update - handles window  update and closing logic
func (wm *WindowManager) Update() {
	wm.win.Update()

	wm.shouldClose = wm.shouldClose || wm.win.Closed()

	if wm.shouldClose {
		util.Running = false
	}
}

//fps calculations
func (wm *WindowManager) fpsUpdate() {
	wm.frames++
	select {
	case <-second:
		wm.lastFps = wm.frames

		var newTitle = fmt.Sprintf("%s | FPS: %d", wm.windowTitle, wm.lastFps)

		log.Print(newTitle)

		wm.SetTitle(newTitle)
		wm.frames = 0
	default:
	}
}

// SetTitle sets the window's title
func (wm *WindowManager) SetTitle(title string) {
	wm.win.SetTitle(title)
	log.Debug().Str("title", title).Msg("Window title updated.")
}

// Window returns a pointer to the current pixelgl.Window instnace
func (wm *WindowManager) Window() *pixelgl.Window {
	return wm.win
}
