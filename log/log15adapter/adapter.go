package log15adapter

import (
	"context"

	"github.com/jackc/pgx/v4/log/log15adapter"

	"github.com/fgm/pgxer"
)

type Log15Logger interface {
	Debug(msg string, ctx ...interface{})
	Info(msg string, ctx ...interface{})
	Warn(msg string, ctx ...interface{})
	Error(msg string, ctx ...interface{})
	Crit(msg string, ctx ...interface{})
}

type logger struct {
	internal *log15adapter.Logger
}

func NewLogger(l Log15Logger) pgxer.Logger {
	return &logger{internal: log15adapter.NewLogger(l)}
}

func (l *logger) Log(ctx context.Context, level pgxer.LogLevel, msg string, data map[string]interface{}) {
	l.internal.Log(ctx, level, msg, data)
}
