package world

import "encoding/binary"

type TileType byte

const (
	TileTypeEmpty = iota
	TileTypeStone
)

type TileBitmask uint16

const (
	FlagActive = 1 << iota
	// FlagStateful
)

func (tb TileBitmask) Bytes() []byte {
	var bytes []byte = make([]byte, 2)
	binary.LittleEndian.PutUint16(bytes, uint16(tb))
	return bytes
}

// AddFlag add a flag to the bitmask
func (f *TileBitmask) AddFlag(flag TileBitmask) { *f |= flag }

// HasFlag deteremine if bitmask has a flag type
func (f TileBitmask) HasFlag(flag TileBitmask) bool {
	return f&flag != 0
}

func (f *TileBitmask) RemoveFlag(flag TileBitmask) {
	*f &^= flag
}
