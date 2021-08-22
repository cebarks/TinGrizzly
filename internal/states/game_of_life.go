package states

import (
	"image"
	"image/color"
	"math/rand"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/kelindar/tile"
	"golang.org/x/image/colornames"
)

// StateConway is a testing state for any dev work
type StateConway struct {
	Grid *tile.Grid
}

func (s StateConway) Update(sc *StateContext, dt float64) error {
	s.Grid.Each(func(p tile.Point, t tile.Tile) {
		alive := countNeighbors(s.Grid, p)

		newTile := Dead

		if tile, _ := s.Grid.At(p.X, p.Y); tile == Alive {
			if alive < 2 {
				newTile = Dead
			}
			if alive == 2 || alive == 3 {
				newTile = Alive
			}
			if alive > 3 {
				newTile = Dead
			}
		} else {
			if alive == 3 {
				newTile = Alive
			}
		}

		s.Grid.WriteAt(p.X, p.Y, newTile)
	})
	// time.Sleep(200 * time.Millisecond)
	return nil
}

func countNeighbors(grid *tile.Grid, p tile.Point) int {
	var count int
	grid.Neighbors(p.X, p.Y, func(p tile.Point, t tile.Tile) {
		if t == Alive {
			count++
		}
	})
	return count
}

func (s *StateConway) Render(win *pixelgl.Window) error {
	win.Clear(colornames.Mediumpurple)

	img := image.NewGray16(image.Rect(0, 0, 500, 500))

	var x, y int16
	for x = 0; x < 500; x++ {
		for y = 0; y < 500; y++ {
			if tile, _ := s.Grid.At(x, y); tile == Alive {
				img.SetGray16(int(x), int(y), color.Black)
			} else {
				img.SetGray16(int(x), int(y), color.White)
			}
		}
	}

	pixel.NewSprite(pixel.PictureDataFromImage(img), pixel.R(0, 0, 500, 500)).Draw(win, pixel.IM.Scaled(pixel.ZV, 2).Moved(win.Bounds().Center()))

	return nil
}

func (s *StateConway) Start() {
	s.Grid = tile.NewGrid(500, 500)

	var x, y int16
	for x = 0; x < 500; x++ {
		for y = 0; y < 500; y++ {
			if rand.Intn(100) < 10 {
				s.Grid.WriteAt(x, y, Dead)
			} else {
				s.Grid.WriteAt(x, y, Alive)
			}
		}
	}
}

func (s StateConway) Stop() {
}

var (
	Dead  tile.Tile = tile.Tile{0}
	Alive tile.Tile = tile.Tile{1}
)
