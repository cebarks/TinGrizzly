package world

import (
	"testing"

	"github.com/kelindar/tile"
	"github.com/stretchr/testify/assert"
)

func Test_StateMap(t *testing.T) {
	state := NewState()

	state.Set("test", "hello, world!")

	actual, err := state.Get("test")

	assert.Nil(t, err)
	assert.Equal(t, "hello, world!", actual)

	actual, err = state.Get("test2")
	assert.NotNil(t, err)
	assert.Nil(t, actual)

	actual, flag := state.GetWithDefault("test3", "Hello, World!")
	actual2, err2 := state.Get("test3")

	assert.False(t, flag)
	assert.Equal(t, "Hello, World!", actual)

	assert.Nil(t, err2)
	assert.Equal(t, "Hello, World!", actual2)
}

func Test_TypeWhitelist(t *testing.T) {
	state := NewState()

	assert.NotPanics(t, func() {
		state.Set("location", tile.At(42, 42))
		state.Set("tile_type", "empty")
	})

	assert.Panics(t, func() {
		state.Set("location", "test")
		state.Set("location", 42)
		state.Set("tile_type", 1)
	})
}
