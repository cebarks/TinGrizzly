package util

import (
	"fmt"
	"sort"
	"strings"
	"time"

	"github.com/lrita/cmap"
)

var DebugMap *cmap.Cmap

func debugInit() {
	DebugMap = &cmap.Cmap{}
	DebugMap.Store("commit", GitCommit)
	DebugMap.Store("start_time", time.Now())
}

func BuildDebugSubtitle() string {
	var subs []string
	DebugMap.Range(func(key, value interface{}) bool {
		subs = append(subs, fmt.Sprintf("%s=%v", key, value))
		return true
	})

	sort.Slice(subs, func(i, j int) bool { return subs[i] < subs[j] })

	return strings.Join(subs, "|")
}
