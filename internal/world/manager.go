package world

import (
	"sort"

	"github.com/cebarks/TinGrizzly/internal/util"
	"github.com/kelindar/tile"
	"github.com/panjf2000/ants/v2"
	"github.com/rs/zerolog/log"
)

type Manager struct {
	World   *World
	Systems systems

	//workPool is the pool used for processing any concurrent work during the update loop
	workPool *ants.Pool
}

func NewManager(world *World) *Manager {
	pool, err := ants.NewPool(util.Cfg().Core.Tunables.MaxWorldPoolWorkers, ants.WithOptions(ants.Options{Nonblocking: true, Logger: &log.Logger}))
	if util.DebugError(err) {
		log.Fatal().Err(err).Msgf("An error occured when trying to initialize the world worker pool.")
	}

	return &Manager{
		Systems:  make(systems, 0),
		workPool: pool,
		World:    world,
	}
}

func (manager *Manager) Update(delta float64) error {
	for _, sys := range manager.Systems {
		if sys.Type&SystemTypeDummy != 0 {
			log.Debug().Msg("ignoring dummy system during update tick")
			continue
		}

		if sys.Type&SystemTypeEntity != 0 {
			for _, ent := range manager.World.Entities {
				sys.UpdateEntity(delta, ent)
			}
		}

		if sys.Type&SystemTypeTile != 0 {
			manager.World.Grid.Each(func(p tile.Point, t tile.Tile) {
				sys.UpdateTile(delta, manager.World.TileDataLookupFromTile(t))
			})
		}
	}
	return manager.World.Update(delta)
}

func (manager *Manager) AddSystem(system *System) {
	manager.Systems = append(manager.Systems, system)

	sort.Sort(manager.Systems)
}
