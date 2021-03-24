package world

type TileType byte

const (
	TileTypeEmpty = iota
	TileTypeStone
)

type TileBitmask byte

const (
	FlagActive = 1 << iota
	// FlagStateful
)

// AddFlag add a flag to the bitmask
func (f *TileBitmask) AddFlag(flag TileBitmask) { *f |= flag }

// HasFlag deteremine if bitmask has a flag type
func (f TileBitmask) HasFlag(flag TileBitmask) bool { return f&flag != 0 }

// RemoveFlag removes a flag from the bitmask
func (f *TileBitmask) RemoveFlag(flag TileBitmask) { *f &^= flag }

func newTileStateForType(typ TileType) TileState {
	switch typ {
	case TileTypeStone:
		return &TileStateStone{}
	default:
		return &TileStateEmpty{}
	}
}
