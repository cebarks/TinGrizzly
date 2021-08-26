package world

type SystemType int8

const (
	SystemTypeDummy SystemType = 1 << iota
	SystemTypeTile
	SystemTypeEntity
)

type System struct {
	TileUpdateFunc   func(sys *System, target *TileData, delta float64) error
	EntityUpdateFunc func(sys *System, target *Entity, delta float64) error

	Priority   int
	Id         string
	Type       SystemType
	Concurrent bool
}

func (s *System) UpdateTile(delta float64, tile *TileData) error {
	return s.TileUpdateFunc(s, tile, delta)
}

func (s *System) UpdateEntity(delta float64, entity *Entity) error {
	return s.EntityUpdateFunc(s, entity, delta)
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
