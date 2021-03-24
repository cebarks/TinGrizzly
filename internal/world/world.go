package world

import (
	"sync"

	"github.com/rs/zerolog/log"

	"github.com/cebarks/TinGrizzly/internal/util"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/kelindar/tile"
	"golang.org/x/image/colornames"
)

type TileData struct {
	Type     TileType
	Location tile.Point
	State    TileState
	Header   TileHeader
}

type World struct {
	Lookup     map[uint32]*TileData
	Grid       *tile.Grid
	ballSprite *pixel.Sprite
	Canvas     *pixelgl.Canvas
	tileBatch  *pixel.Batch
}

func (w *World) Render(win *pixelgl.Window) {
	w.Canvas.Clear(colornames.Whitesmoke)

	w.tileBatch.Clear()

	w.Grid.Each(func(p tile.Point, t tile.Tile) {
		td := w.TileDataLookupFromTile(t)

		mat := pixel.IM
		mat = pixel.IM.Moved(pixel.V(128, 128))

		var sprite *pixel.Sprite

		switch td.Type {
		default:
		case TileTypeStone:
			sprite = w.ballSprite
		}

		if sprite == nil {
			return
		}

		mat = mat.Moved(util.PointToVecScaled(p, 256))
		mat = mat.Scaled(pixel.ZV, .1)
		sprite.Draw(w.tileBatch, mat) //draw each tile to a batch.Draw instead of individually
	})

	w.tileBatch.Draw(w.Canvas) //draw tiles to world canvas

	w.Canvas.Draw(win, pixel.IM.Moved(win.Bounds().Center())) //draw the world canvas to the center of the window
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

	for _, work := range work {
		go func(tu tileUpdate) {
			tu.td.State.Update(tu.w, tu.p, tu.delta)
			defer wg.Done()
		}(work)
	}

	wg.Wait()
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
		log.Fatal().Err(err).Msg("Couldn't load asset: assets/ball.png")
	}

	world.ballSprite = pixel.NewSprite(pic, pic.Bounds())
	world.tileBatch = pixel.NewBatch(&pixel.TrianglesData{}, pic)
	world.Canvas = pixelgl.NewCanvas(pixel.R(1, 1, 1024, 1024))

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
func (td TileData) Index() uint32 {
	return td.Location.Integer()
}

//initTile inits TileData for all tiles in the grid and saves a header pointing to it in the Lookup map.
func initTile(world *World, p tile.Point) {
	tileData := TileData{
		Type:     TileTypeEmpty,
		Location: p,
	}

	tileData.State = &TileStateEmpty{}

	header := &TileHeader{
		Index: tileData.Index(),
	}

	header.Bitmask.AddFlag(FlagActive)

	header.Save(world.Grid, p)

	world.Lookup[tileData.Index()] = &tileData
}

func (w *World) SetTileTo(p tile.Point, typ TileType) *TileData {
	td := w.TileDataLookup(p.X, p.Y)

	td.Type = typ
	td.State = newTileStateForType(typ)
	return td
}
