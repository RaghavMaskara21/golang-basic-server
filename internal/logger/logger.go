package logger

import (
	"github.com/rs/zerolog"
	zerologger "hayday/server/internal/zerolog"
)

var Log *Logger

type Logger struct {
	Log zerolog.Logger
}

type LoggerEvent struct {
	Event zerolog.Logger
}

func InitiateLogger() *Logger {
	Log = &Logger{
		Log: zerologger.CreateZeroLogger(),
	}
	return Log
}

func (logger *Logger) Infof(format string, args ...interface{}) {
	logger.Log.Info().Msgf(format, args...)
}

func (logger *Logger) Debugf(format string, args ...interface{}) {
	logger.Log.Debug().Msgf(format, args...)
}

func (logger *Logger) Warnf(format string, args ...interface{}) {
	logger.Log.Warn().Msgf(format, args...)
}

func (logger *Logger) Errorf(format string, args ...interface{}) {
	logger.Log.Error().Msgf(format, args...)
}

func (logger *Logger) Fatalf(format string, args ...interface{}) {
	logger.Log.Fatal().Msgf(format, args...)
}

func (logger *Logger) WithFields(fields map[string]interface{}) *LoggerEvent {
	return &LoggerEvent{
		Event: logger.Log.With().Fields(fields).Logger(),
	}
}

func (logger *LoggerEvent) Infof(format string, args ...interface{}) {
	logger.Event.Info().Msgf(format, args...)
}

func (logger *LoggerEvent) Debugf(format string, args ...interface{}) {
	logger.Event.Debug().Msgf(format, args...)
}

func (logger *LoggerEvent) Warnf(format string, args ...interface{}) {
	logger.Event.Warn().Msgf(format, args...)
}

func (logger *LoggerEvent) Errorf(format string, args ...interface{}) {
	logger.Event.Error().Msgf(format, args...)
}

func (logger *LoggerEvent) Fatalf(format string, args ...interface{}) {
	logger.Event.Fatal().Msgf(format, args...)
}
