package util

import (
	"os"

	"github.com/rs/zerolog/log"

	"github.com/faiface/pixel"
	"github.com/kelindar/tile"
)

var (
	// Running - global variable for app state
	Running bool
	//GitCommit is set to the most recent git commit hash at build time
	GitCommit string

	TargetFPS int64 = 144
	TargetUPS int64 = 50
)

func Startup() {
	SetupLogging()
	log.Info().Str("version", GitCommit).Msg("Launching...")
	ReloadCfgFromDisk()
	Running = true

	debugInit()
}

func PointToVec(p tile.Point) pixel.Vec {
	return pixel.V(float64(p.X), float64(p.Y))
}

func PointToVecScaled(p tile.Point, scale float64) pixel.Vec {
	return pixel.V(float64(p.X)*scale, float64(p.Y)*scale)
}

//FileExists returns true if the given path exists and isn't a directory
func FileExists(filename string) bool {
	s, err := os.Stat(filename)

	if err != nil {
		if os.IsNotExist(err) {
			return false
		}
		log.Error().Err(err).Msgf("helpers.FileExists(%s) errored out:", err) // If the err wasn't expected, something really went wrong
	} else if s == nil { // If no s is returned there is a different issue.
		return false
	}
	return !s.IsDir()
}
