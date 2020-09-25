package main

import (
	"fmt"
	"time"

	"github.com/cebarks/TinGrizzly/states"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

var (
	//window settings
	win          pixelgl.Window
	windowTitle  = "Tin Grizzly!"
	windowBounds = pixel.R(0, 0, 1600, 900)
	windowVsync  = false
	winCfg       pixelgl.WindowConfig

	//fps vars
	frames  = 0
	lastFps = 0
	second  = time.Tick(time.Second)
	last    = time.Now()

	//StateManager global instacne
	StateManager states.StateManager
)

func run() {
	println("Launched!")

	winCfg := pixelgl.WindowConfig{
		Title:  windowTitle,
		Bounds: windowBounds,
		VSync:  windowVsync,
	}

	win, err := pixelgl.NewWindow(winCfg)
	if err != nil {
		panic(err)
	}

	win.Clear(colornames.Skyblue)

	StateManager = states.StateManager{}
	StateManager.Initialize()

	for !win.Closed() {
		win.Update()
		fpsUpdate()

		//Tick main game loop
		gameLoop()
	}
}

func gameLoop() {
	//Calculate delta time for update calculations
	dt := time.Since(last).Seconds()
	last = time.Now()

	update(dt)
	render(win)
}

func update(dt float64) {
	StateManager.ActiveState.Update(dt)
	x := StateManager.ActiveState
	x.Start()
}

func render(win pixelgl.Window) {
	StateManager.ActiveState.Render(win)
}

func fpsUpdate() {
	//fps calculations
	frames++
	select {
	case <-second:
		lastFps = frames
		win.SetTitle(fmt.Sprintf("%s | FPS: %d", winCfg.Title, lastFps))
		frames = 0
	default:
	}
}

func main() {
	println("Launching...")
	pixelgl.Run(run)
}
