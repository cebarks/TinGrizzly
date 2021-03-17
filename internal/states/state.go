package states

import (
	"reflect"

	"github.com/rs/zerolog/log"

	"github.com/cebarks/TinGrizzly/internal/gfx"
	"github.com/faiface/pixel/pixelgl"
)

// StateManager is reposible for holding information about the games current state and related information
type StateManager struct {
	ActiveState State
	stateMap    map[string]State
}

// State implements the base interface/struct for any State
type State interface {
	Update(win *StateContext, dt float64)
	Render(win *pixelgl.Window)
	Start()
	Stop()
	// Pause()
}

type StateContext struct {
	StateManager  *StateManager
	WindowManager *gfx.WindowManager
}

func (sm *StateManager) SetState(state string) {
	oldState := sm.ActiveState
	newState := sm.stateMap[state]

	if newState == nil {
		log.Fatal().Msgf("Couldn't set state: %s", state)
	}

	log.Debug().Msgf("Switching states: %+v -> %+v", reflect.TypeOf(oldState), reflect.TypeOf(newState))

	if oldState != nil {
		oldState.Stop()
	}
	sm.ActiveState = newState
	newState.Start()
	log.Fatal().Msgf("Switched states.")
}

// BuildStateManager returns a new StateManager already initialized with default values.
func BuildStateManager() *StateManager {
	sm := StateManager{
		stateMap: map[string]State{
			"null":     &StateNull{},
			"mainMenu": &StateMainMenu{},
			"dev":      &StateDevelopment{},
		},
	}

	// sm.SetState("null")
	sm.SetState("dev")

	return &sm
}
