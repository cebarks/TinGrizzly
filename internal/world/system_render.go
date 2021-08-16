package world

import (
	"github.com/cebarks/TinGrizzly/resources"
	"github.com/faiface/pixel"
)

type RenderSystem struct {
	tileBatch *pixel.Batch
	// Canvas    *pixelgl.Canvas
}

func BuildRenderSystem() *RenderSystem {
	rs := &RenderSystem{
		tileBatch: pixel.NewBatch(&pixel.TrianglesData{}, resources.Sheet.SourcePic()),
	}

	// rs.Canvas =

	return rs
}

func (rs *RenderSystem) Update(delta float64, entity Entity) {

}
