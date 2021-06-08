package world

import (
	"github.com/cebarks/TinGrizzly/internal/world/ecs"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

type RenderSystem struct {
	tileBatch *pixel.Batch
	Canvas    *pixelgl.Canvas
}

func BuildRenderSystem() *RenderSystem {
	rs := &RenderSystem{
		tileBatch: pixel.NewBatch(&pixel.TrianglesData{}, resources.),
	}

	rs.tileBatch = 
	rs.Canvas = pixelgl.NewCanvas(pixel.R(1, 1, 1024, 1024))

	return rs
}

func (rs *RenderSystem) Update(delta float64, entity ecs.Entity) {

}
