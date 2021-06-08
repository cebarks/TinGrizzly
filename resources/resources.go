package resources

import (
	"embed"
	"image"
	"io"
	"io/fs"

	"github.com/cebarks/spriteplus"
	"github.com/faiface/pixel"
	"github.com/rs/zerolog/log"
)

//go:embed *
var resourceEmbed embed.FS

var sheet *spriteplus.SpriteSheet

func Setup() {
	tilesFile, _ := resourceEmbed.Open("assets/tiles.png")

	// var err error
	sheet := spriteplus.NewSpriteSheet(true)

	sheet.
}

func GetResource(path string) fs.File {
	file, err := resourceEmbed.Open(path)

	if err == fs.ErrNotExist {
		log.Panic().Msgf("Tried to load non-existant resource: %s", path)
	} else if err != nil {
		log.Panic().Err(err).Msgf("Couldn't load resource: %s", path)
	}

	return file
}

func GetResourceBytes(path string) []byte {
	bytes, err := resourceEmbed.ReadFile(path)

	if err == fs.ErrNotExist {
		log.Panic().Msgf("Tried to load non-existant resource: %s", path)
	} else if err != nil {
		log.Panic().Err(err).Msgf("Couldn't load resource: %s", path)
	}

	return bytes
}

// func GetSprite(sheet spriteplus.SpriteSheet, sprite string) *pixel.Sprite {
func GetSprite(sprite string) *pixel.Sprite {
	return Tiles.GetSprite(sprite)
}

func loadPictureFromReader(r io.Reader) pixel.Picture {
	i, fmt, err := image.Decode(r)
	if err != nil {
		log.Fatal().Err(err).Msgf("Couldn't load picture.")
	}
	log.Trace().Msgf("Loaded image; format: %s", fmt)

	return pixel.PictureDataFromImage(i)
}
