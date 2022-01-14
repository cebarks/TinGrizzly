package gfx

import (
	"github.com/cebarks/TinGrizzly/internal/world"
	"github.com/cebarks/TinGrizzly/resources"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/kelindar/tile"
	"github.com/rs/zerolog/log"
)

type WorldCamera struct {
	View   *tile.View
	Canvas *pixelgl.Canvas
}

func NewCamera(view *tile.View) *WorldCamera {
	if view == nil {
		log.Panic().Msg("can't create camera with nil view")
	}
	return &WorldCamera{
		View:   view,
		Canvas: pixelgl.NewCanvas(pixel.R(0, 0, 21*16, 21*16)),
	}
}

func (wc *WorldCamera) Update(w *world.World) {
	// batch := pixel.NewBatch(pixel.MakeTrianglesData(0), resources.Sheet.SourcePic())
	log.Trace().Msg("Updating camera canvas")
	wc.View.Each(func(p tile.Point, t tile.Tile) {
		td := w.TileDataLookupFromTile(t)

		tts, err := td.State.Get("type")
		if err != nil {
			log.Panic().Err(err).Msg("could not get tile_type for tile in view")
			return
		}

		spriteName, err := td.State.Get("sprite")
		log.Printf("id=%s spriteName1=%s", tts, spriteName)
		if err != nil {
			// log.Warn().Msgf("Tile '%s' does not have a sprite set!", tiles.TileTypes[tts.(string)])
			spriteName = "error"
		}

		var sprite *pixel.Sprite = resources.GetSprite(spriteName.(string))
		log.Printf("spriteName2: %s", spriteName)
		// if typ := tt.(world.TileType); typ == world.TileTypeStone {
		// 	sprite = resources.GetSprite("stone")
		// } else if typ == world.TileTypeEmpty {
		// 	sprite = resources.GetSprite("grass")
		// } else {
		// 	sprite = resources.GetSprite("error")
		// }

		// sprite.Draw(batch, pixel.IM.Moved(pixel.V(8, 8)).Moved(pixel.V(float64(p.X*16), float64(p.Y*16))))
		sprite.Draw(wc.Canvas, pixel.IM.Moved(pixel.V(8, 8)).Moved(pixel.V(float64(p.X*16), float64(p.Y*16))))
	})
	// batch.Draw(wc.Canvas)
}
