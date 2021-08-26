package world

import (
	"testing"

	"github.com/kelindar/tile"
	"github.com/stretchr/testify/assert"
)

func Test_ToTile(t *testing.T) {
	header := &TileHeader{
		Bitmask: FlagActive,
		Index:   425678,
	}

	expected := tile.Tile{0x2a, 0x1, 0xce, 0x7e, 0x6, 0x0}

	assert.Equal(t, expected, header.ToTile())
}

func Test_FromTile(t *testing.T) {
	expected := &TileHeader{
		Bitmask: FlagActive,
		Index:   425678,
	}

	actual := HeaderFromTile(tile.Tile{0x2a, 0x1, 0xce, 0x7e, 0x6, 0x0})

	assert.Equal(t, expected, actual)
}

func Test_Save(t *testing.T) {
	grid := tile.NewGrid(3, 3)

	td := newTileData(tile.At(2, 2))
	td.Save(grid)

	actual, _ := grid.At(2, 2)
	expected := tile.Tile{0x2a, 0x1, 0x2, 0x0, 0x2, 0x0}

	assert.Equal(t, expected, actual)
}
