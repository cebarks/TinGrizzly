package states

import (
	"log"

	"github.com/cebarks/TinGrizzly/internal/gfx"
	"github.com/faiface/pixel/pixelgl"
)

// StateManager is reposible for holding information about the games current state and related information
type StateManager struct {
	ActiveState State
	stateMap    map[string]State
}

// State implements the base interface/struct for  any State
type State interface {
	Update(win *StateContext, dt float64)
	Render(win *pixelgl.Window)
	Start()
	Pause()
	Stop()
}

type StateContext struct {
	StateManager  *StateManager
	WindowManager *gfx.WindowManager
}

// Initialize all states
func (sm *StateManager) Initialize() {
	// create map of state instances
	sm.stateMap = map[string]State{
		"null":     StateNull{},
		"mainMenu": StateMainMenu{},
		"dev":      StateDevelopment{},
	}

	// sm.SetState("null")
	sm.SetState("dev")
}

func (sm *StateManager) SetState(state string) {
	oldState := sm.ActiveState
	newState := sm.stateMap[state]

	if newState == nil {
		log.Fatalf("Couldn't set state: %s", state)
	}

	if oldState != nil {
		oldState.Stop()
	}
	sm.ActiveState = newState
	newState.Start()
}

// BuildStateManager returns a new StateManager already initialized with default values.
func BuildStateManager() *StateManager {
	sm := StateManager{}
	sm.Initialize()
	return &sm
}
