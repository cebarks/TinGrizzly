package world

import (
	"sync"

	"github.com/cebarks/TinGrizzly/internal/util"
	"github.com/kelindar/tile"
	"github.com/rs/zerolog/log"
)

type World struct {
	Lookup map[uint32]*TileData
	Grid   *tile.Grid
	State  *State
}

type tileUpdate struct {
	w     *World
	td    *TileData
	delta float64
	p     tile.Point
}

func (w *World) Update(delta float64) {
	var wg sync.WaitGroup
	var work []tileUpdate
	w.Grid.Each(func(p tile.Point, t tile.Tile) {
		td := w.TileDataLookupFromTile(t)

		if td.Header.Bitmask.HasFlag(FlagActive) {
			work = append(work, tileUpdate{
				delta: delta,
				w:     w,
				p:     p,
				td:    td,
			})
			wg.Add(1)
		}
	})

	wg.Wait()
}

//NewWorld creates an world of the given size with blank tiles.
func NewWorld(sizeX, sizeY int16) *World {
	world := &World{
		Lookup: make(map[uint32]*TileData, sizeX*sizeY),
		Grid:   tile.NewGrid(sizeX, sizeY),
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
	var header TileHeader

	header.FromTile(t)

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
	return state.(tile.Point).Integer()
}

//initTile inits TileData for all tiles in the grid and saves a header pointing to it in the Lookup map.
func initEmptyTile(world *World, p tile.Point) {
	tileData := TileData{}

	tileData.State.Set("location", p)
	tileData.State.Set("type", TileTypeEmpty)

	header := &TileHeader{
		Index: tileData.Index(),
	}

	header.Bitmask.AddFlag(FlagActive)

	header.SaveTo(world.Grid, p)

	world.Lookup[tileData.Index()] = &tileData
}

func (w *World) SetTileTo(p tile.Point, ttype TileType) *TileData {
	td := w.TileDataLookup(p.X, p.Y)

	td.State.Set("type", ttype)

	return td
}
