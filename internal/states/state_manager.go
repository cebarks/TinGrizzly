package states

import (
	"reflect"

	"github.com/rs/zerolog/log"

	"github.com/cebarks/TinGrizzly/internal/gfx"
	"github.com/cebarks/TinGrizzly/internal/util"
	"github.com/faiface/pixel/pixelgl"
)

// StateManager is reposible for holding information about the games current state and related information
type StateManager struct {
	ActiveState State
	stateMap    map[string]State
}

// State implements the base interface/struct for any State
type State interface {
	Update(win *StateContext, dt float64) error
	Render(win *pixelgl.Window) error
	Start()
	Stop()
}

type StateContext struct {
	StateManager  *StateManager
	WindowManager *gfx.WindowManager
}

func (sm *StateManager) SetState(state string) {
	oldState := sm.ActiveState
	newState := sm.stateMap[state]

	if newState == nil {
		log.Panic().Msgf("State not registered. Couldn't set new state: %s", state)
	}

	log.Debug().Msgf("Switching states: %+v -> %+v", reflect.TypeOf(oldState), reflect.TypeOf(newState))

	util.ShouldRender = false

	if oldState != nil {
		oldState.Stop()
	}
	sm.ActiveState = newState
	newState.Start()

	util.ShouldRender = true

	log.Debug().Msgf("Switched states.")
}

// BuildStateManager returns a new StateManager already initialized with default values.
func BuildStateManager() *StateManager {
	sm := StateManager{
		stateMap: map[string]State{
			"null":     &StateNull{},
			"mainMenu": &StateMainMenu{},
			"dev":      &StateDevelopment{},
			"resource": &StateResource{},
			"conway":   &StateConway{},
			"loading":  &StateLoading{},
		},
	}

	sm.SetState("loading")

	return &sm
}

//CheckStateKeys switches the active state based on number keys
func CheckStateKeys(sc *StateContext) {
	if sc.WindowManager.JustPressed(pixelgl.Key0) {
		sc.StateManager.SetState("null")
	} else if sc.WindowManager.JustPressed(pixelgl.Key1) {
		sc.StateManager.SetState("dev")
	} else if sc.WindowManager.JustPressed(pixelgl.Key2) {
		sc.StateManager.SetState("resource")
	} else if sc.WindowManager.JustPressed(pixelgl.Key9) {
		sc.StateManager.SetState("conway")
	}
}
