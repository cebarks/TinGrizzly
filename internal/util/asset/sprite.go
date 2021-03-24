package asset

import (
	"fmt"

	"github.com/cebarks/TinGrizzly/internal/util"
	"github.com/dusk125/pixelutils"
	"github.com/faiface/pixel"
	"github.com/rs/zerolog/log"
)

type SpriteSheet struct {
	
}

var Sprites map[string]*pixel.Sprite = make(map[string]*pixel.Sprite)

func init() {
	// LoadTileSet()
}

func LoadTileSet() {
	pic, err := pixelutils.LoadPictureData("assets/tileset.png")
	if util.DebugError(err) {
		log.Fatal().Err(err).Msg("Couldn't load tileset")
	}

	for x := 0; x < 12; x++ {
		for y := 0; y < 43; y++ {
			startX := float64(x * 16)
			startY := float64(y * 16)
			endX := float64((x + 1) * 16)
			endY := float64((y + 1) * 16)
			Sprites[fmt.Sprintf("%v:%v", x, y)] = pixel.NewSprite(pic, pixel.R(startX, startY, endX, endY))
		}
	}

	log.Debug().Msgf("Loaded %v Sprites from tileset.", len(Sprites))
}

func GetSprite(sprite string) *pixel.Sprite {
	lookup := Sprites[sprite]

	if lookup == nil {
		lookup, err := pixelutils.LoadSprite(sprite)

		if util.DebugError(err) {
			log.Panic().Err(err).Msg("Couldn't load sprite!")
		}
		Sprites[sprite] = lookup
	}

	return lookup
}
