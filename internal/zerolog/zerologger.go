package zerologger

import (
	"os"
	"time"

	"github.com/rs/zerolog"
)

func CreateZeroLogger() zerolog.Logger {
	consoleWriter := zerolog.ConsoleWriter{
		Out:        os.Stdout,
		TimeFormat: time.RFC3339,
	}
	return zerolog.New(consoleWriter).With().Timestamp().Caller().Logger()
}
