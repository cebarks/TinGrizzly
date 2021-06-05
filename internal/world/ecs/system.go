package ecs

type System interface {
	Update(delta float64, entity Entity)
	Remove(entity Entity)
	IsConcurrent() bool
}

type Component interface {
	Update(delta float64, entity Entity)
}
