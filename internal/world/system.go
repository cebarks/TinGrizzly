package world

import "github.com/kelindar/tile"

type SystemType int8

const (
	SystemTypeDummy SystemType = iota
	SystemTypeTile
	SystemTypeEntity
)

type System struct {
	TileUpdateFunc   func(sys System, target tile.Point, delta float64)
	EntityUpdateFunc func(sys System, target Entity, delta float64)

	UpdateView *tile.View

	Priority   int
	Id         string
	Type       int
	Concurrent bool
}

func (s *System) UpdateTile(delta float64, entity Entity, loc tile.Point) error {
	return nil
}

func (s *System) UpdateEntity(delta float64, entity Entity, loc tile.Point) error {
	return nil
}

//systems is a sortable slice of systems
type systems []*System

func (s systems) Len() int {
	return len(s)
}

func (s systems) Less(i, j int) bool {
	return s[i].Priority > s[j].Priority
}

func (s systems) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
