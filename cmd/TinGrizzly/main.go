package main

import (
	"github.com/cebarks/TinGrizzly/internal/game"
	"github.com/cebarks/TinGrizzly/internal/util"

	"github.com/faiface/pixel/pixelgl"
)

func main() {
	gam := &game.Game{}
	util.Startup()

	util.SetupCloseHandler()

	pixelgl.Run(gam.Run)
}
