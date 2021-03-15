package world

type TileState struct {
	TileBitmask
}

func (ts TileState) IsActive() bool {
	return ts.TileBitmask.HasFlag(FlagActive)
}

func (ts *TileState) SetActive() {
	ts.TileBitmask.AddFlag(FlagActive)
}
