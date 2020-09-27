package main

import (
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/cebarks/TinGrizzly/gfx"
	"github.com/cebarks/TinGrizzly/states"

	"github.com/faiface/pixel/pixelgl"
)

var (
	// dt calculation
	lastTime = time.Now()

	//StateManager global instance
	StateManager states.StateManager
	//WindowManager global instance
	WindowManager *gfx.WindowManager
)

func main() {
	println("Launching...")
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	pixelgl.Run(run)
}

func gameLoop() {
	//Calculate delta time for update calculations
	dt := time.Since(lastTime).Seconds()
	lastTime = time.Now()

	update(dt)
	render(WindowManager.Window())
}

func update(dt float64) {
	as := StateManager.ActiveState
	as.Update(dt)
}

func render(win *pixelgl.Window) {
	as := StateManager.ActiveState
	as.Render(win)
}

func run() {
	println("Launched!")

	StateManager = states.BuildStateManager()
	WindowManager = &gfx.WindowManager{}

	WindowManager.Initalize()
	WindowManager.CreateWindow()

	log.Print("Made it here!  =======================================")
	for true {
		// gameLoop()

		WindowManager.Update()
	}
}
