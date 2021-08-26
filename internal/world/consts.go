package world

type TileType uint

const (
	TileTypeEmpty TileType = iota
	TileTypeStone
)

type TileBitmask byte

const (
	FlagActive TileBitmask = 1 << iota
	FlagStateless
)

// AddFlag add a flag to the bitmask
func (f *TileBitmask) AddFlag(flag TileBitmask) { *f |= flag }

// HasFlag deteremine if bitmask has a flag type
func (f TileBitmask) HasFlag(flag TileBitmask) bool { return f&flag != 0 }

// ClearFlag removes a flag from the bitmask
func (f *TileBitmask) ClearFlag(flag TileBitmask) { *f &= ^flag }

// ToggleFlag toggles a flag to the opposite value
func (f *TileBitmask) ToggleFlag(flag TileBitmask) { *f ^= flag }
