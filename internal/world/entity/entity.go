package entity

import (
	"reflect"

	"github.com/rs/zerolog/log"

	"github.com/cebarks/TinGrizzly/internal/world/entity/component"
	"github.com/faiface/pixel"
)

type Entity struct {
	Rotation    pixel.Vec
	Active      bool
	Components  []*component.Component
	BoundingBox *pixel.Rect
}

func (entity *Entity) AddComponent(new *component.Component) {
	for _, existing := range entity.Components {
		if reflect.TypeOf(new) == reflect.TypeOf(existing) {
			log.Panic().Msgf("Attempt to add a new component with already existing type: %v", reflect.TypeOf(new))
		}
	}
	entity.Components = append(entity.Components, new)
}

func (entity *Entity) GetComponent(withType component.Component) *component.Component {
	for _, comp := range entity.Components {
		if reflect.TypeOf(comp) == reflect.TypeOf(withType) {
			return comp
		}
	}
	log.Panic().Msgf("No component with type: %v", reflect.TypeOf(withType))
	return nil
}
