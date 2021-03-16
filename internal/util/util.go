package util

import (
	"log"

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
