package gfx

import (
	"fmt"
	"log"
	"sort"
	"strings"
	"sync"

	"github.com/cebarks/TinGrizzly/internal/util"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
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
	window, err := pixelgl.NewWindow(pixelgl.WindowConfig{
		Title:     title,
		Bounds:    pixel.R(0, 0, 1920, 1080),
		VSync:     false,
		Resizable: false,
		// Monitor:   pixelgl.PrimaryMonitor(),
	})

	if util.DebugError(err) {
		log.Fatalf("Couldn't create window: %v", err)
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
