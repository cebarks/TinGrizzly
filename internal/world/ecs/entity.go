package ecs

import "github.com/dusk125/pixelutils"

type Entity struct {
	Id int

	Components []Component
}

var idGen *pixelutils.IDGen = pixelutils.NewIDGen()

func NewEntity(components ...Component) *Entity {
	return &Entity{
		Components: components,
		Id:         idGen.Gen(),
	}
}
