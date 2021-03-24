package util

import (
	"image"
	_ "image/png"
	"os"

	"github.com/rs/zerolog/log"

	"github.com/faiface/pixel"
	"github.com/kelindar/tile"
)

var (
	// Running - global variable for app state
	Running   bool
	GitCommit string

	TargetFPS int64 = 144
	TargetUPS int64 = 50
)

func Startup() {
	SetupLogging()
	log.Info().Str("version", GitCommit).Msg("Launching...")
	ReloadCfgFromDisk()
	Running = true
}

func LoadPicture(path string) (pixel.Picture, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}
	return pixel.PictureDataFromImage(img), nil
}

func PointToVec(p tile.Point) pixel.Vec {
	return pixel.V(float64(p.X), float64(p.Y))
}

func PointToVecScaled(p tile.Point, scale float64) pixel.Vec {
	return pixel.V(float64(p.X)*scale, float64(p.Y)*scale)
}
