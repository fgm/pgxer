package zapadapter

import (
	"context"

	"github.com/jackc/pgx/v4/log/zapadapter"
	"go.uber.org/zap"

	"github.com/fgm/pgxer"
)

type logger struct {
	internal *zapadapter.Logger
}

func NewLogger(l *zap.Logger) pgxer.Logger {
	return &logger{internal: zapadapter.NewLogger(l.WithOptions(zap.AddCallerSkip(1)))}
}

func (l *logger) Log(ctx context.Context, level pgxer.LogLevel, msg string, data map[string]interface{}) {
	l.internal.Log(ctx, level, msg, data)
}
