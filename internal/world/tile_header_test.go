package world

import (
	"testing"

	"github.com/kelindar/tile"
	"github.com/stretchr/testify/assert"
)

// func Test_Tile(t *testing.T) {
// 	header := TileHeader{
// 		Bitmask: FlagActive,
// 		Index:   425678,
// 	}

// 	expected := tile.Tile{0x1, 0x0, 0xce, 0x7e, 0x6, 0x0}

// 	assert.Equal(t, expected, header.Tile())
// }

func Test_Save(t *testing.T) {
	grid := tile.NewGrid(3, 3)

	header := TileHeader{
		Bitmask: FlagActive,
		Index:   425678,
	}

	header.Save(grid, tile.At(2, 2))

	actual, _ := grid.At(2, 2)
	expected := tile.Tile{0x0, 0x1, 0xce, 0x7e, 0x6, 0x0}

	assert.Equal(t, expected, actual)
}

func Test_FromTile(t *testing.T) {
	grid := tile.NewGrid(3, 3)

	expected := TileHeader{
		Bitmask: FlagActive,
		Index:   425678,
	}

	rawHeader := tile.Tile{0x0, 0x1, 0xce, 0x7e, 0x6, 0x0}

	grid.WriteAt(2, 2, rawHeader)

	var header TileHeader
	header.FromTile(rawHeader)

	assert.Equal(t, expected, header)
}

func Test_Load(t *testing.T) {
	grid := tile.NewGrid(3, 3)

	expected := TileHeader{
		Bitmask: FlagActive,
		Index:   425678,
	}

	grid.WriteAt(2, 2, tile.Tile{0x0, 0x1, 0xce, 0x7e, 0x6, 0x0})

	var header TileHeader
	header.Load(grid, tile.At(2, 2))

	assert.Equal(t, expected, header)
}
