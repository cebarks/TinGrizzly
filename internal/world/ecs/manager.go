package ecs

import (
	"sort"
	"sync"

	"github.com/cebarks/TinGrizzly/internal/util"
	"github.com/panjf2000/ants/v2"
	"github.com/rs/zerolog/log"
)

type Manager struct {
	Systems  systems
	Entities []Entity

	//sysCompMap stores component <-> system mappings
	sysCompMap map[Component]System

	//workPool is the pool used for processing any concurrent work during the update loop
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
	var cSystems []System

	for _, system := range manager.Systems {
		if system.IsConcurrent() {
			cSystems = append(cSystems, system)
		} else {
			for _, entity := range manager.Entities {
				system.Update(delta, entity)
			}
		}
	}

	for _, system := range cSystems {
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

func (manager *Manager) AddSystem(system System, components ...Component) {
	manager.Systems = append(manager.Systems, system)

	if len(components) < 1 {
		for _, comp := range components {
			manager.sysCompMap[comp] = system
		}
	}

	sort.Sort(manager.Systems)
}