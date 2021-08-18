package world

import (
	"fmt"
	"reflect"

	"github.com/kelindar/tile"
	"github.com/lrita/cmap"
	"github.com/rs/zerolog/log"
)

var enforcedTypes map[string]reflect.Type = map[string]reflect.Type{
	"tile_type": reflect.TypeOf(TileTypeEmpty),
	"location":  reflect.TypeOf(tile.Point{}),
}

type State struct {
	store *cmap.Cmap
}

func (s *State) Get(key string) (interface{}, error) {
	if value, ok := s.store.Load(key); ok {
		return value, nil
	}
	return nil, fmt.Errorf("attempt to access invalid state key")
}

func (s *State) GetWithDefault(key string, defaultValue interface{}) (interface{}, bool) {
	return s.store.LoadOrStore(key, defaultValue)
}

func (s *State) Set(key string, value interface{}) {
	for k, t := range enforcedTypes {
		if k == key {
			if reflect.TypeOf(value) == t {
				log.Error().Msgf("invalid type `%T` for state-key `%s`", value, key)
				log.Trace().Msgf("state dump: %+v", s)
			}
		}
	}
	s.store.Store(key, value)
}
