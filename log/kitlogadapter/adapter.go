package kitlogadapter

import (
	"context"

	"github.com/go-kit/log"
	"github.com/jackc/pgx/v4/log/kitlogadapter"

	"github.com/fgm/pgxer"
)

type logger struct {
	internal *kitlogadapter.Logger
}

func NewLogger(l log.Logger) pgxer.Logger {
	return &logger{internal: kitlogadapter.NewLogger(l)}
}

func (l *logger) Log(ctx context.Context, level pgxer.LogLevel, msg string, data map[string]interface{}) {
	l.internal.Log(ctx, level, msg, data)
}
