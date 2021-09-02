package gfx

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/rs/zerolog/log"

	"github.com/cebarks/TinGrizzly/internal/util"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/go-gl/glfw/v3.3/glfw"
)

type WindowManager struct {
	*pixelgl.Window
}

var (
	title string = "TinGrizzly"
)

func BuildWindowManager() *WindowManager {
	glfw.WindowHint(glfw.Samples, util.Cfg().Graphics.Samples)
	glfw.WindowHint(glfw.RefreshRate, 0)

	w, h := parseResolution(util.Cfg().Graphics.Resolution)

	window, err := pixelgl.NewWindow(pixelgl.WindowConfig{
		Title:     title,
		Bounds:    pixel.R(0, 0, w, h),
		VSync:     util.Cfg().Graphics.Vsync,
		Resizable: false,
	})

	if util.DebugError(err) {
		log.Fatal().Msgf("Couldn't create window: %v", err)
	}

	if util.Cfg().Graphics.Fullscreen {
		window.SetMonitor(pixelgl.PrimaryMonitor())
	}

	winm := WindowManager{Window: window}
	return &winm
}

func (wm *WindowManager) SetSubtitle(subtitle string) {
	wm.SetTitle(fmt.Sprintf("%s - %s", title, subtitle))
}

func parseResolution(res string) (float64, float64) {
	split := strings.Split(res, "x")

	w, err := strconv.ParseInt(split[0], 10, 0)
	if err != nil {
		log.Fatal().Err(err).Msgf("Invalid resolution defined: %s", res)
	}

	h, err := strconv.ParseInt(split[1], 10, 0)
	if err != nil {
		log.Fatal().Err(err).Msgf("Invalid resolution defined: %s", res)
	}

	return float64(w), float64(h)
}
