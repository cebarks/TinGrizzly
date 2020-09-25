package states

// StateManager is reposible for holding information about the games current state and related information
type StateManager struct {
	ActiveState State
	stateMap    map[string]State
}

// Initialize all states
func (st StateManager) Initialize() {
	st.stateMap = map[string]State{}
	st.stateMap["mainMenu"] = StateMainMenu{}
	st.stateMap["dev"] = StateDevelopment{}
	// gs.stateMap["overworld"] = StateOverWorld
	// gs.stateMap["dungeon1"] = StateDungeon
}
