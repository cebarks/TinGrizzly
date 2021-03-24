package util

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/rs/zerolog/log"
)

// SetupCloseHandler creates a 'listener' on a new goroutine which will notify the
// program if it receives an interrupt from the OS. We then handle this by calling
// our clean up procedure and exiting the program.
func SetupCloseHandler() {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		log.Info().Msg("SIGTERM Received.")
		Running = false
	}()
}
