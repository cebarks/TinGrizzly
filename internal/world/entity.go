package world

import (
	"github.com/dusk125/pixelutils"
	"github.com/lrita/cmap"
)

type Entity struct {
	Components []string
	Id         int
	State      *cmap.Cmap
}

var eidGen *pixelutils.IDGen = &pixelutils.IDGen{}

func NewEntity(components ...string) *Entity {
	return &Entity{
		Components: components,
		Id:         eidGen.Gen(),
	}
}
