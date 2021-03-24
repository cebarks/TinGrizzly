package util

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/snwfdhmp/errlog"
)

var (
	errorLogger errlog.Logger
	// DebugError handles an error with errlog
	DebugError func(error) bool
)

//SetupLogging sets up errlog and zerolog and sets errlog to use zerolog to
func SetupLogging() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout})

	consoleWriter := zerolog.ConsoleWriter{
		Out:        os.Stdout,
		NoColor:    false,
		TimeFormat: time.RFC1123Z,
		FormatCaller: func(i interface{}) string {
			path := i.(string)
			idx := strings.Index(path, "/TinGrizzly/") + 12
			//trim build path from output
			return path[idx:]
		},
	}

	wd, _ := os.Getwd()
	path := filepath.Join(wd, "log", fmt.Sprintf("log-%s.json", strings.ReplaceAll(time.Now().Format(time.Stamp), " ", "-")))

	os.Mkdir("log", 0755)

	logFile, err := os.Create(path)
	if err != nil {
		log.Error().Err(err).Msg("Couldn't create log file.")
	} else {
		log.Info().Msgf("Saving log file to: %s", path)
	}

	// fileWriter := bufio.NewWriter(logFile)
	fileWriter := zerolog.New(logFile).With().Logger()

	// go func() {
	// 	defer fileWriter.Flush()
	// 	for {
	// 		fileWriter.Flush()
	// 		time.Sleep(25 * time.Millisecond)
	// 	}
	// }()

	log.Logger = log.Output(zerolog.MultiLevelWriter(consoleWriter, fileWriter))

	errorLogger = errlog.NewLogger(&errlog.Config{
		LinesBefore:        3,
		LinesAfter:         4,
		PrintError:         true,
		PrintSource:        true,
		PrintStack:         true,
		ExitOnDebugSuccess: false,
		PrintFunc: func(format string, data ...interface{}) {
			log.Error().Msgf(format, data...)
		},
	})

	if zerolog.GlobalLevel() <= zerolog.TraceLevel {
		//adds file and line number to log
		log.Logger = log.With().Caller().Logger()
	} else {
		errorLogger.Disable(true)
	}

	DebugError = errorLogger.Debug
}
