package world

import (
	"github.com/cebarks/TinGrizzly/internal/util"
	"github.com/cebarks/TinGrizzly/internal/world/tiles"
	"github.com/kelindar/tile"
	"github.com/rs/zerolog/log"
)

type World struct {
	Lookup   map[uint32]*TileData
	Grid     *tile.Grid
	State    *State
	Entities []*Entity
}

type tileUpdate struct {
	w     *World
	td    *TileData
	delta float64
	p     tile.Point
}

func (w *World) Update(delta float64) error {
	return nil
}

//NewWorld creates an world of the given size with blank tiles. World size must be a multiple of 3.
func NewWorld(sizeX, sizeY int16) *World { //TODO: support bigger maps backed by multiple grids?
	world := &World{
		Lookup:   make(map[uint32]*TileData, sizeX*sizeY),
		Entities: make([]*Entity, 42),
		Grid:     tile.NewGrid(sizeX, sizeY),
	}

	world.Grid.Each(func(p tile.Point, t tile.Tile) {
		initEmptyTile(world, p)
	})

	return world
}

//TileDataLookup returns the TileData associated with the coordinates
func (w *World) TileDataLookup(x, y int16) *TileData {
	t, _ := w.Grid.At(x, y)
	return w.TileDataLookupFromTile(t)
}

//TileDataLookupFromPoint returns the TileData associated with the coordinates
func (w *World) TileDataLookupFromPoint(p tile.Point) *TileData {
	return w.TileDataLookup(p.X, p.Y)
}

//TileDataLookupFromTile returns the TileData associated with the given tile
func (w *World) TileDataLookupFromTile(t tile.Tile) *TileData {
	header := HeaderFromTile(t)

	tileData := w.Lookup[header.Index]
	tileData.Header = header

	return tileData
}

//Index returns the key used for (*world.World).Lookup
func (td *TileData) Index() uint32 {
	state, err := td.State.Get("location")

	if util.DebugError(err) {
		log.Error().Err(err).Msg("couldn't generate index for tiledata")
	}

	return state.(*tile.Point).Integer()
}

//initTile inits TileData for all tiles in the grid and saves a header pointing to it in the Lookup map.
func initEmptyTile(world *World, p tile.Point) {
	td := newTileData(p)

	if td2 := world.Lookup[td.Index()]; td2 != nil {
		log.Panic().Msgf("Duplicate tiledata index (%v): %+v, %+v", td.Index(), td, td2)
	}

	td.State.Set("type", "empty")

	td.Save(world.Grid)

	world.Lookup[td.Index()] = td
}

func newTileData(p tile.Point) *TileData {
	td := &TileData{
		State: NewState(),
		Header: &TileHeader{
			Index:   p.Integer(),
			Bitmask: FlagActive,
		},
	}

	td.State.Set("location", &p)
	td.State.Set("type", "empty")

	return td
}

func (w *World) SetTileTo(p tile.Point, typeId string) *TileData {
	td := w.TileDataLookup(p.X, p.Y)

	td.State.Set("type", tiles.TileTypes[typeId])

	return td
}
