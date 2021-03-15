package world

import (
	"encoding/binary"
	"log"

	"github.com/kelindar/tile"
)

type TileData struct {
	Type     TileType
	Location tile.Point
	State    *TileState
}

type World struct {
	Lookup map[uint32]*TileData
	Grid   *tile.Grid
}

func (w *World) TileDataFromTile(t tile.Tile) *TileData {
	metaBytes := t[:2]
	indexBytes := t[2:]

	meta := binary.LittleEndian.Uint16(metaBytes)
	index := binary.LittleEndian.Uint32(indexBytes)

	log.Printf("index: %v | meta: %v", index, meta)

	tileData := w.Lookup[index]
	return tileData
}

func NewWorld(sizeX, sizeY int16) *World {
	world := &World{
		Lookup: make(map[uint32]*TileData, sizeX*sizeY),
		Grid:   tile.NewGrid(sizeX, sizeY),
	}

	world.Grid.Each(func(p tile.Point, t tile.Tile) {
		tileData := TileData{
			Type:     TileTypeEmpty,
			State:    &TileState{TileBitmask: 0},
			Location: p,
		}
		world.Lookup[tileData.Index()] = &tileData
	})

	return world
}

func (td TileData) Bytes() []byte {
	maskBytes := td.State.TileBitmask.Bytes()

	indexBytes := make([]byte, 4)
	binary.LittleEndian.PutUint32(indexBytes, td.Index())

	return append(maskBytes, indexBytes...)
}

func (td TileData) Index() uint32 {
	return td.Location.Integer()
}
