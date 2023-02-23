package selfupdate

import (
	"io"
	stdlog "log"
	"os"
)

var log = stdlog.New(io.Discard, "", 0)
var logEnabled = false

// EnableLog enables to output logging messages in library.
func EnableLog() {
	if logEnabled {
		return
	}

	logEnabled = true

	log.SetOutput(os.Stderr)

	log.SetFlags(stdlog.Ltime)
}

// DisableLog disables to output logging messages in library.
func DisableLog() {
	if !logEnabled {
		return
	}

	logEnabled = false

	log.SetOutput(io.Discard)

	log.SetFlags(0)
}
