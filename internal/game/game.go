package game

import (
	"fmt"

	"github.com/rs/zerolog/log"

	"github.com/cebarks/TinGrizzly/internal/gfx"
	"github.com/cebarks/TinGrizzly/internal/states"
	"github.com/cebarks/TinGrizzly/internal/util"
	"github.com/cebarks/TinGrizzly/internal/util/asset"
	"github.com/dusk125/pixelutils"
	"github.com/faiface/pixel/pixelgl"
)

type Game struct {
	StateManager  *states.StateManager
	WindowManager *gfx.WindowManager
	debugMenu     *map[string]interface{}
}

func (game *Game) Run() {
	game.WindowManager = gfx.BuildWindowManager()
	game.StateManager = states.BuildStateManager()
	game.debugMenu = &map[string]interface{}{}
	asset.LoadTileSet()

	log.Info().Msg("Launched!")
	go game.renderLoop()
	game.updateLoop()

	log.Info().Msg("Closing.")
}

func (game *Game) updateLoop() {
	ticker := pixelutils.NewTickerV(util.TargetUPS, util.TargetUPS*60)

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

		game.StateManager.ActiveState.Update(ctx, dt)

		game.WindowManager.SetSubtitle("ups", fmt.Sprintf("%.f", ups))

		game.WindowManager.UpdateSubtitles()
		ticker.Wait()
	}
}

func (game *Game) renderLoop() {
	ticker := pixelutils.NewTickerV(util.TargetFPS, util.TargetFPS*60)

	for util.Running {
		_, fps := ticker.Tick()

		game.StateManager.ActiveState.Render(game.WindowManager.Window)
		game.WindowManager.SetSubtitle("fps", fmt.Sprintf("%.f", fps))

		game.WindowManager.Update()
		ticker.Wait()
	}
}
