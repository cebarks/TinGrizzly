package asset

import (
	"github.com/cebarks/TinGrizzly/internal/util"
	"github.com/dusk125/pixelutils"
	"github.com/faiface/pixel"
	"github.com/rs/zerolog/log"
)

var sprites map[string]*pixel.Sprite = make(map[string]*pixel.Sprite)

func GetSprite(sprite string) *pixel.Sprite {
	lookup := sprites[sprite]

	if lookup == nil {
		lookup, err := pixelutils.LoadSprite(sprite)

		if util.DebugError(err) {
			log.Panic().Err(err).Msg("Couldn't load sprite!")
		}
		sprites[sprite] = lookup
	}

	return lookup
}

