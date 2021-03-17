package util

import (
	"image"
	_ "image/png"
	"os"

	"github.com/rs/zerolog/log"

	"github.com/faiface/pixel"
	"github.com/kelindar/tile"
	"github.com/snwfdhmp/errlog"
)

var (
	// Running - global variable for app state
	Running     bool
	errorLogger errlog.Logger
)

func init() {
	SetupLogging()
}

//SetupLogging sets up errlog and zerolog and sets errlog to use zerolog to
func SetupLogging() {
	errorLogger = errlog.NewLogger(&errlog.Config{
		PrintFunc:          log.Printf,
		LinesBefore:        6,
		LinesAfter:         4,
		PrintError:         true,
		PrintSource:        true,
		PrintStack:         true,
		ExitOnDebugSuccess: false,
	})
}

// DebugError handles an error with errlog
func DebugError(err error) bool {
	return errorLogger.Debug(err)
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
