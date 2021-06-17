package resources

import (
	"embed"
	"image"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/cebarks/TinGrizzly/internal/util"
	"github.com/cebarks/spriteplus"
	"github.com/faiface/pixel"
	"github.com/rs/zerolog/log"
)

//go:embed *
var resourceEmbed embed.FS

var Sheet *spriteplus.SpriteSheet

func Setup() {
	Sheet = spriteplus.NewSpriteSheet(true)

	err := Sheet.AddSprite(pixel.PictureDataFromImage(loadImageFromReader(GetResource("assets/ball.png"))), "ball")
	if util.DebugError(err) {
		log.Panic().Err(err).Msg("could not load ball.png")
	}

	importTileSet("assets/tiles.png")
}

func importTileSet(defPath string) {
	// importedTileCount, err := Sheet.AddTileset(loadImageFromReader(GetResource(path)), ids, 16, 16, 1, 1)
	// if util.DebugError(err) {
	// 	log.Panic().Err(err).Msg("could not load tileset: %s", defPath)
	// }
	// log.Debug().Msgf("Imported %d tiles from tileset: %s", importedTileCount, "assets/tiles.png")
}

func GetResource(path string) fs.File {
	var file fs.File
	var src string
	var err error

	if util.Cfg().Graphics.Resources.Embedded {
		file, err = resourceEmbed.Open(path)
		src = "embedded"
	} else {
		file, err = os.Open(filepath.Join("resources", path))
		src = "filesystem"
	}

	if err == fs.ErrNotExist {
		log.Panic().Msgf("Tried to load non-existant resource (%s) from %s", path, src)
	} else if err != nil {
		log.Panic().Err(err).Msgf("Couldn't load resource (%s) from %s", path, src)
	}

	log.Trace().Msgf("Loaded resource (%s) from %s", path, src)
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

func GetResourceString(path string) string {
	bytes := GetResourceBytes(path)
	b := strings.Builder{}

	for _, byt := range bytes {
		b.WriteByte(byt)
	}

	return b.String()
}

//GetResourceImage returns a resource as an image.Image
func GetResourceImage(path string) (image.Image, string) {
	r := GetResource(path)

	img, fmt, err := image.Decode(r)
	if util.DebugError(err) {
		log.Panic().Err(err).Msgf("Failed to get resource as image: %s", path)
	}

	return img, fmt
}

func GetSprite(sprite string) *pixel.Sprite {
	return Sheet.GetSprite(sprite)
}

func loadImageFromReader(r io.Reader) image.Image {
	i, fmt, err := image.Decode(r)

	if err != nil {
		log.Fatal().Err(err).Msgf("Couldn't load picture.")
	}

	log.Trace().Msgf("Loaded image; format: %s", fmt)
	return i
}
