package logger

import (
	"fmt"
	"io"

	"github.com/rs/zerolog"
)

type Logger struct {
	log *zerolog.Logger
}

func New(level string, w io.Writer) *Logger {
	switch level {
	case "INFO":
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	case "ERROR":
		zerolog.SetGlobalLevel(zerolog.ErrorLevel)
	case "WARN":
		zerolog.SetGlobalLevel(zerolog.WarnLevel)
	case "DEBUG":
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	default:
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}
	log := zerolog.New(w).With().Timestamp().Logger()
	return &Logger{&log}
}

func (l Logger) Info(msg string) {
	l.log.Info().Msg(msg)
}

func (l Logger) Error(msg string) {
	l.log.Error().Msg(msg)
}

func (l Logger) Debug(msg string) {
	l.log.Debug().Msg(msg)
}

func (l Logger) Warn(msg string) {
	l.log.Warn().Msg(msg)
}

func (l Logger) Errorf(format string, args ...interface{}) {
	l.Error(fmt.Sprintf(format, args...))
}
