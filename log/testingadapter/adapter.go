package testingadapter

import (
	"context"

	"github.com/jackc/pgx/v4/log/testingadapter"

	"github.com/fgm/pgxer"
)

type TestingLogger interface {
	Log(args ...interface{})
}

type logger struct {
	internal *testingadapter.Logger
}

func NewLogger(l TestingLogger) pgxer.Logger {
	return &logger{internal: testingadapter.NewLogger(l)}
}

func (l *logger) Log(ctx context.Context, level pgxer.LogLevel, msg string, data map[string]interface{}) {
	l.internal.Log(ctx, level, msg, data)
}
