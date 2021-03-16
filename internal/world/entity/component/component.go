package component

import "github.com/faiface/pixel/pixelgl"

type Component interface {
	onUpdate(dt float64) error
	onDraw(*pixelgl.Window) error
}
