package world

import (
	"github.com/cebarks/TinGrizzly/internal/util"
	"github.com/kelindar/tile"
)

type TileState interface {
	Update(world *World, p tile.Point, delta float64)
}

type TileStateEmpty struct {
	counter1 int64
}

func (ts *TileStateEmpty) Update(world *World, p tile.Point, delta float64) {
	ts.counter1++
	if ts.counter1 > util.TargetUPS*2 {
		world.SetTileTo(p, TileTypeStone)
		// log.Debug().Msg("changing to Stone")
	}
}

type TileStateStone struct {
	counter int64
}

func (ts *TileStateStone) Update(world *World, p tile.Point, delta float64) {
	ts.counter++
	if ts.counter > util.TargetUPS*2 {
		world.SetTileTo(p, TileTypeEmpty)
		// log.Debug().Msg("changing to Empty")
	}
}
