package logger

import (
	"context"
	"os"

	"github.com/rs/zerolog"
)

type Logger interface {
	Info(msg string, data ...interface{})
	Error(msg string, data ...interface{})
	SetRequestID(requestId string) Logger
}

type logger struct {
	log *zerolog.Logger
}

type LOG_LEVEL int8

const (
	LogKey LoggerKey = "LOGGER"
)

type LoggerKey string

const (
	TRACE LOG_LEVEL = iota - 1
	DEBUG
	INFO
	WARN
	ERROR
	FATAL
	PANIC
)

func GetLoggerCtx(ctx context.Context) Logger {
	logger, ok := ctx.Value(LogKey).(*logger)
	if !ok {
		return nil
	}

	return logger
}

func NewLogger(level LOG_LEVEL) Logger {
	log := zerolog.New(os.Stdout).Level(zerolog.Level(level))
	return &logger{
		log: &log,
	}
}

func (l *logger) SetRequestID(requestId string) Logger {
	log := l.log.With().Str("requestId", requestId).Logger()
	return &logger{
		log: &log,
	}
}

func (l *logger) Error(msg string, data ...interface{}) {
	lg := l.log.Error()
	if len(data) > 0 {
		lg.Any("data", data)
	}
	lg.Msg(msg)
}

func (l *logger) Info(msg string, data ...interface{}) {
	lg := l.log.Info()
	if len(data) > 0 {
		lg.Any("data", data)
	}
	lg.Msg(msg)
}
