package game

import (
	"time"

	"github.com/rs/zerolog/log"

	"github.com/cebarks/TinGrizzly/internal/gfx"
	"github.com/cebarks/TinGrizzly/internal/states"
	"github.com/cebarks/TinGrizzly/internal/util"
	"github.com/dusk125/pixelutils"
	"github.com/faiface/pixel/pixelgl"
)

type Game struct {
	StateManager  *states.StateManager
	WindowManager *gfx.WindowManager
}

func (game *Game) Run() {
	game.WindowManager = gfx.BuildWindowManager()
	game.StateManager = states.BuildStateManager()

	log.Info().Msg("Initialized!")

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
		util.DebugMap.Store("ups", ups)

		ctx := &states.StateContext{
			StateManager:  game.StateManager,
			WindowManager: game.WindowManager,
		}
		states.CheckStateKeys(ctx)

		stateErr := game.StateManager.ActiveState.Update(ctx, dt)

		if util.DebugError(stateErr) {
			log.Error().Err(stateErr).Msg("error during game loop.")
		}

		game.WindowManager.SetSubtitle(util.BuildDebugSubtitle())
		game.WindowManager.UpdateInput()
		ticker.Wait()
	}
}

func (game *Game) renderLoop() {
	ticker := pixelutils.NewTickerV(util.Cfg().Core.Tunables.Fps, util.Cfg().Core.Tunables.Fps*60)

	for util.Running {
		_, fps := ticker.Tick()
		util.DebugMap.Store("fps", fps)

		if game.WindowManager.Window.Focused() && util.ShouldRender {
			game.StateManager.ActiveState.Render(game.WindowManager.Window)
		} else {
			time.Sleep(time.Second / 30)
		}

		game.WindowManager.SwapBuffers()

		if !util.Cfg().Core.Tunables.UnlockFps {
			ticker.Wait()
		}
	}
}
