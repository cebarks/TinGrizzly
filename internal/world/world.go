package world

import (
	"github.com/cebarks/TinGrizzly/internal/util"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
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
	sprite *pixel.Sprite
}

func (w *World) Render(win *pixelgl.Window) {
	w.Grid.Each(func(p tile.Point, t tile.Tile) {
		td := w.TileDataLookupFromTile(t)

		mat := pixel.IM
		//mat = pixel.IM.Moved(win.Bounds().Center())
		mat = pixel.IM.Moved(pixel.V(128, 128))

		switch td.Type {
		case TileTypeEmpty:
			return
		case TileTypeStone:
			mat = mat.Moved(util.PointToVecScaled(p, 256))
			mat = mat.Scaled(pixel.ZV, .1)
			w.sprite.Draw(win, mat)
		}
	})
}

func NewWorld(sizeX, sizeY int16) *World {
	world := &World{
		Lookup: make(map[uint32]*TileData, sizeX*sizeY),
		Grid:   tile.NewGrid(sizeX, sizeY),
	}

	world.Grid.Each(func(p tile.Point, t tile.Tile) {
		initTile(world, p)
	})

	pic, err := util.LoadPicture("assets/ball.png")
	if err != nil {
		panic(err)
	}

	world.sprite = pixel.NewSprite(pic, pic.Bounds())

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
