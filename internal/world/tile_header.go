package world

import (
	"bytes"
	"encoding/binary"
	"log"

	"github.com/cebarks/TinGrizzly/internal/util"
	"github.com/kelindar/tile"
)

type TileHeader struct {
	Bitmask TileBitmask
	_       byte
	Index   uint32
}

//Tile returns a tile representation of this TileHeader
// func (th *TileHeader) Tile() tile.Tile {
// 	buf := new(bytes.Buffer)
// 	binary.Write(buf, binary.LittleEndian, th)
// 	var t tile.Tile
// 	copy(t[:], buf.Bytes())
// 	return t
// }

//Save saves this TileHeader to the given point and grid.
func (header *TileHeader) Save(grid *tile.Grid, p tile.Point) {
	buf := new(bytes.Buffer)
	if err := binary.Write(buf, binary.LittleEndian, header); util.DebugError(err) {
		log.Fatal(err) //TODO log.Fatal call
	}

	var t tile.Tile
	copy(t[:], buf.Bytes())
	grid.WriteAt(p.X, p.Y, t)
}

//Load loads the tile from the given point and grid into this TileHeader.
func (header *TileHeader) Load(grid *tile.Grid, p tile.Point) {
	t, _ := grid.At(p.X, p.Y)
	header.FromTile(t)
}

//FromTile loads the given tile into this TileHeader.
func (header *TileHeader) FromTile(t tile.Tile) {
	buf := new(bytes.Buffer)
	for _, b := range t {
		buf.WriteByte(b)
	}

	if err := binary.Read(buf, binary.LittleEndian, header); util.DebugError(err) {
		log.Fatal(err) //TODO log.Fatal call
	}
}
