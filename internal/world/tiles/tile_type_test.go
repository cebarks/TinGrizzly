package tiles_test

import (
	"testing"

	"github.com/cebarks/TinGrizzly/internal/world/tiles"
	"github.com/cebarks/TinGrizzly/resources"
	"github.com/stretchr/testify/assert"
)

var testTile = []byte(`{
    "id":"test",
    "sprite": "error",
    "extra" : {
        "test1":"test123",
        "test2":"321test"
    }
}
`)

func Test_LoadTileType(t *testing.T) {
	assert.NotNil(t, tiles.TileTypes)
	assert.Equal(t, 0, len(tiles.TileTypes))

	ttyp, err := resources.LoadTileTypeFromBytes(testTile)

	assert.Nil(t, err)
	assert.NotNil(t, ttyp)

	assert.Equal(t, 1, len(tiles.TileTypes))

	assert.Equal(t, "test", ttyp.Id)
	assert.Equal(t, "error", ttyp.Sprite)
	assert.Equal(t, "test123", ttyp.Data["test1"])
	assert.Equal(t, "321test", ttyp.Data["test2"])
}
