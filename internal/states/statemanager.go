package states

// StateManager is reposible for holding information about the games current state and related information
type StateManager struct {
	ActiveState State
	stateMap    map[string]State
}

// Initialize all states
func (sm *StateManager) Initialize() {
	// create map of state instances
	sm.stateMap = map[string]State{}
	sm.stateMap["null"] = StateNull{}
	sm.stateMap["mainMenu"] = StateMainMenu{}
	sm.stateMap["dev"] = StateDevelopment{}
	// sm.stateMap["overworld"] = StateOverWorld
	// sm.stateMap["dungeon1"] = StateDungeon

	sm.ActiveState = sm.stateMap["null"]
}

// BuildStateManager returns a new StateManager already initialized with default values.
func BuildStateManager() StateManager {
	sm := StateManager{}
	sm.Initialize()
	return sm
}
