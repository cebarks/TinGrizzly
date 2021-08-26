package util

import (
	"testing"

	"github.com/faiface/pixel"
	"github.com/kelindar/tile"
	"github.com/stretchr/testify/assert"
)

func Test_PointToVec(t *testing.T) {
	assert.Equal(t, pixel.V(42, 42), PointToVec(tile.At(42, 42)))
}

func Test_PointToVecScaled(t *testing.T) {
	assert.Equal(t, pixel.V(42, 42), PointToVecScaled(tile.At(21, 21), 2))
}

func Test_FileExists(t *testing.T) {
	assert.True(t, FileExists("./util_test.go"))
}
