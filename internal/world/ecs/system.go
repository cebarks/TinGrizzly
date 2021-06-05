package ecs

type System interface {
	Update(delta float64, entity Entity)
	Remove(entity Entity)

	IsConcurrent() bool
	Priority() int
}

type Component interface {
	Update(delta float64, entity Entity)
}

//systems is a sortable slice of systems
type systems []System

func (s systems) Len() int {
	return len(s)
}

func (s systems) Less(i, j int) bool {
	return s[i].Priority() > s[j].Priority()
}

func (s systems) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
