package ecs

import (
	"github.com/dusk125/pixelutils"
	"github.com/lrita/cmap"
)

type Entity struct {
	Components []Component
	Id         int
	State      *cmap.Cmap
}

var idGen *pixelutils.IDGen = pixelutils.NewIDGen()

func NewEntity(components ...Component) *Entity {
	return &Entity{
		Components: components,
		Id:         idGen.Gen(),
	}
}
