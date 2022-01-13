package world

import (
	"github.com/dusk125/pixelutils"
)

type Entity struct {
	Components []string
	Id         int
	State      *State
}

var eidGen *pixelutils.IDGen = &pixelutils.IDGen{}

func NewEntity(components ...string) *Entity {
	return &Entity{
		Components: components,
		Id:         eidGen.Gen(),
	}
}
