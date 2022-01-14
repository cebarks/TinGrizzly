package world

import (
	"fmt"
	"reflect"

	"github.com/kelindar/tile"
	"github.com/lrita/cmap"
	"github.com/rs/zerolog/log"
)

var enforcedTypes map[string]reflect.Type = map[string]reflect.Type{
	"type": reflect.TypeOf(""),
	// "type":     reflect.TypeOf(&TileType{}),
	"location": reflect.TypeOf(&tile.Point{}),
	"sprite":   reflect.TypeOf(""),
}

type State struct {
	store *cmap.Cmap
}

func NewState() *State {
	return &State{
		store: &cmap.Cmap{},
	}
}

func (s *State) Get(key string) (interface{}, error) {
	if value, ok := s.store.Load(key); ok {
		return value, nil
	}
	return nil, fmt.Errorf("attempt to access invalid/unset state key")
}

// GetWithDefault returns the existing value for the key if present.
// Otherwise, it stores and returns the given value.
// The loaded result is true if the value was loaded, false if stored.
func (s *State) GetWithDefault(key string, defaultValue interface{}) (interface{}, bool) {
	return s.store.LoadOrStore(key, defaultValue)
}

func (s *State) Set(key string, value interface{}) {
	for k, t := range enforcedTypes {
		if k == key {
			if reflect.TypeOf(value) != t {
				log.Panic().Msgf("invalid type `%T` for state-key `%s`", value, key)
			}
		}
	}
	s.store.Store(key, value)
}
