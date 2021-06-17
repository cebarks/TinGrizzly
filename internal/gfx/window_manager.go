package gfx

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
	"sync"

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
	title         string = "TinGrizzly"
	subtitles            = make(map[string]interface{})
	subtitleMutex        = &sync.Mutex{}
)

func BuildWindowManager() *WindowManager {
	w, h := parseResolution(util.Cfg().Graphics.Resolution)

	glfw.WindowHint(glfw.Samples, util.Cfg().Graphics.Samples)
	glfw.WindowHint(glfw.RefreshRate, 0)

	window, err := pixelgl.NewWindow(pixelgl.WindowConfig{
		Title:     title,
		Bounds:    pixel.R(0, 0, w, h),
		VSync:     util.Cfg().Graphics.Vsync,
		Resizable: false,
		// Monitor:   pixelgl.PrimaryMonitor(),
	})

	if util.Cfg().Graphics.Fullscreen {
		window.SetMonitor(pixelgl.PrimaryMonitor())
	}

	if util.DebugError(err) {
		log.Fatal().Msgf("Couldn't create window: %v", err)
	}

	winm := WindowManager{Window: window}
	return &winm
}

func (wm *WindowManager) SetSubtitle(key string, subtitle interface{}) {
	subtitleMutex.Lock()
	subtitles[key] = subtitle
	subtitleMutex.Unlock()
}

func (wm *WindowManager) UpdateSubtitles() {
	subtitleMutex.Lock()
	var subs []string
	for k, sub := range subtitles {
		subs = append(subs, fmt.Sprintf("%s=%s", k, sub))
	}

	sort.Slice(subs, func(i, j int) bool { return subs[i] < subs[j] })

	joined := strings.Join(subs, "|")

	wm.SetTitle(fmt.Sprintf("%s - %s", title, joined))
	subtitleMutex.Unlock()
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
