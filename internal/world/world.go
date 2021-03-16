package world

import (
	"log"

	"github.com/kelindar/tile"
)

type TileData struct {
	Type     TileType
	Location tile.Point
	State    *TileState
	Header   TileHeader
}

type World struct {
	Lookup map[uint32]*TileData
	Grid   *tile.Grid
}

func NewWorld(sizeX, sizeY int16) *World {
	world := &World{
		Lookup: make(map[uint32]*TileData, sizeX*sizeY),
		Grid:   tile.NewGrid(sizeX, sizeY),
	}

	world.Grid.Each(func(p tile.Point, t tile.Tile) {
		initTile(world, p)
	})

	return world
}

//TileDataLookupFromTile returns the TileData associated with the coordinates
func (w *World) TileDataLookup(x, y int16) *TileData {
	t, _ := w.Grid.At(x, y)
	return w.TileDataLookupFromTile(t)
}

//TileDataLookupFromTile returns the TileData associated with the given tile
func (w *World) TileDataLookupFromTile(t tile.Tile) *TileData {
	var header TileHeader

	header.FromTile(t)
	log.Printf("header: %+v", header)

	tileData := w.Lookup[header.Index]
	tileData.Header = header

	return tileData
}

//Index returns the key used for (*world.World).Lookup
func (td TileData) Index() uint32 {
	return td.Location.Integer()
}

//initTile inits TileData for all tiles in the grid and saves a header pointing to it in the Lookup map.
func initTile(world *World, p tile.Point) {
	tileData := TileData{
		Type:     TileTypeEmpty,
		State:    &TileState{},
		Location: p,
	}

	header := &TileHeader{
		Bitmask: FlagActive,
		Index:   tileData.Index(),
	}

	header.Save(world.Grid, p)

	world.Lookup[tileData.Index()] = &tileData
}
