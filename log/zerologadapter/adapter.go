package zerologadapter

import (
	"context"

	"github.com/jackc/pgx/v4/log/zerologadapter"
	"github.com/rs/zerolog"

	"github.com/fgm/pgxer"
)

type logger struct {
	internal *zerologadapter.Logger
}

func NewLogger(l zerolog.Logger) pgxer.Logger {
	return &logger{internal: zerologadapter.NewLogger(l.With().Str("module", "pgx").Logger())}
}

func (l *logger) Log(ctx context.Context, level pgxer.LogLevel, msg string, data map[string]interface{}) {
	l.internal.Log(ctx, level, msg, data)
}
