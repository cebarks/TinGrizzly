package ecs

import (
	"sync"

	"github.com/cebarks/TinGrizzly/internal/util"
	"github.com/panjf2000/ants/v2"
	"github.com/rs/zerolog/log"
)

type Manager struct {
	Systems  []System
	Entities []Entity

	workPool *ants.Pool
}

func NewManager() *Manager {
	pool, err := ants.NewPool(util.Cfg().Core.Tunables.MaxWorldPoolWorkers, ants.WithOptions(ants.Options{Nonblocking: true, Logger: &log.Logger}))
	if util.DebugError(err) {
		log.Fatal().Err(err).Msgf("An error occured when trying to initialize the world worker pool.")
	}

	return &Manager{workPool: pool}
}

func (manager *Manager) Update(delta float64) {
	var wg sync.WaitGroup

	for _, system := range manager.Systems {
		for _, entity := range manager.Entities {
			wg.Add(1)
			manager.workPool.Submit(func() {
				defer wg.Done()
				system.Update(delta, entity)
			})
		}
	}

	wg.Wait()
}
