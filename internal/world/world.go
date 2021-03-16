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

func (w *World) TileDataLookup(t tile.Tile) *TileData {
	var header TileHeader

	header.FromTile(t)
	log.Printf("header: %+v", header)

	tileData := w.Lookup[header.Index]
	tileData.Header = header
	return tileData
}

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

func (td TileData) Index() uint32 {
	return td.Location.Integer()
}
