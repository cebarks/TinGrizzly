package gfx

import (
	"fmt"
	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/rs/zerolog/log"
	"github.com/snwfdhmp/errlog"
	"golang.org/x/image/colornames"
)

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

func (wm *WindowManager) Initalize() {
	wm.windowTitle = "Tin Grizzly"
	wm.windowBounds = pixel.R(0, 0, 1600, 900)
	wm.windowVsync = false
	wm.frames = 0
}

func (wm *WindowManager) CreateWindow() {
	wm.winCfg = pixelgl.WindowConfig{
		Title:  wm.windowTitle,
		Bounds: wm.windowBounds,
		VSync:  wm.windowVsync,
	}

	log.Debug().Msgf("Created window config: %d", wm.winCfg)

	var err error
	wm.win, err = pixelgl.NewWindow(wm.winCfg)
	if errlog.Debug(err) {
		panic(err)
	}

	wm.win.Clear(colornames.Darkslateblue)
}

func (wm *WindowManager) Close() {
}

func (wm *WindowManager) Update() {
	wm.win.Update()
	if wm.shouldClose {
		wm.win.Destroy()
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

func (wm *WindowManager) SetTitle(title string) {
	log.Debug().Str("title", title).Msg("Window title updated.")
}

func (wm *WindowManager) Window() *pixelgl.Window {
	return wm.win
}
