package world

import (
	"encoding/binary"

	"github.com/rs/zerolog/log"

	"github.com/kelindar/tile"
)

type TileData struct {
	State  *State
	Header *TileHeader
}

type TileHeader struct {
	_       byte
	Bitmask TileBitmask
	Index   uint32
}

//FromTile loads the given tile into this TileHeader.
func HeaderFromTile(t tile.Tile) *TileHeader {
	magic := t[0]
	rawBitmask := t[1]
	rawIndex := t[2:6]

	if magic != 42 {
		log.Panic().Msg("CORRUPTION: BUT WHERE'S THE ANSWER TO EVERYTHING")
	}

	return &TileHeader{
		Bitmask: TileBitmask(rawBitmask),
		Index:   binary.LittleEndian.Uint32(rawIndex),
	}
}

//ToTile returns a tile version of this header
func (header *TileHeader) ToTile() tile.Tile {
	var index []byte = make([]byte, 4)
	binary.LittleEndian.PutUint32(index, header.Index)
	return tile.Tile{42, byte(header.Bitmask), index[0], index[1], index[2], index[3]}
}

func (td *TileData) Save(grid *tile.Grid) error {
	l, err := td.State.Get("location")
	if err != nil {
		return err
	}

	loc := l.(tile.Point)

	grid.WriteAt(loc.X, loc.Y, td.Header.ToTile())
	return nil
}
