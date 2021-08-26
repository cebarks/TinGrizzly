package game

import (
	"fmt"

	"github.com/rs/zerolog/log"

	"github.com/cebarks/TinGrizzly/internal/gfx"
	"github.com/cebarks/TinGrizzly/internal/states"
	"github.com/cebarks/TinGrizzly/internal/util"
	"github.com/cebarks/TinGrizzly/resources"
	"github.com/dusk125/pixelutils"
	"github.com/faiface/pixel/pixelgl"
)

type Game struct {
	StateManager  *states.StateManager
	WindowManager *gfx.WindowManager
	debugMap      *map[string]interface{}
}

func (game *Game) Run() {
	game.WindowManager = gfx.BuildWindowManager()
	game.StateManager = states.BuildStateManager()
	game.debugMap = &map[string]interface{}{}

	resources.Setup()

	log.Info().Msg("Launched!")
	go game.renderLoop()
	game.updateLoop()

	log.Info().Msg("Closing.")
}

func (game *Game) updateLoop() {
	ticker := pixelutils.NewTickerV(util.Cfg().Core.Tunables.Ups, util.Cfg().Core.Tunables.Ups*60)

	for util.Running {
		if game.WindowManager.Pressed(pixelgl.KeyEscape) || game.WindowManager.Closed() {
			util.Running = false
			return
		}
		dt, ups := ticker.Tick()

		ctx := &states.StateContext{
			StateManager:  game.StateManager,
			WindowManager: game.WindowManager,
		}
		states.CheckStateKeys(ctx)

		stateErr := game.StateManager.ActiveState.Update(ctx, dt)

		if util.DebugError(stateErr) {
			log.Error().Err(stateErr).Msg("error during game loop.")
		}

		game.WindowManager.SetSubtitle("ups", fmt.Sprintf("%.f", ups))

		game.WindowManager.UpdateSubtitles()
		game.WindowManager.UpdateInput()
		ticker.Wait()
	}
}

func (game *Game) renderLoop() {
	ticker := pixelutils.NewTickerV(util.Cfg().Core.Tunables.Fps, util.Cfg().Core.Tunables.Fps*60)

	for util.Running {
		_, fps := ticker.Tick()
		game.WindowManager.SetSubtitle("fps", fmt.Sprintf("%.f", fps))

		if game.WindowManager.Window.Focused() {
			game.StateManager.ActiveState.Render(game.WindowManager.Window)
		}

		game.WindowManager.SwapBuffers()

		ticker.Wait()
	}
}
