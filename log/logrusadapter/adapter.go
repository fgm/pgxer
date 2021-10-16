package logrusadapter

import (
	"context"

	"github.com/jackc/pgx/v4/log/logrusadapter"
	"github.com/sirupsen/logrus"

	"github.com/fgm/pgxer"
)

type logger struct {
	internal *logrusadapter.Logger
}

func NewLogger(l logrus.FieldLogger) pgxer.Logger {
	return &logger{internal: logrusadapter.NewLogger(l)}
}

func (l *logger) Log(ctx context.Context, level pgxer.LogLevel, msg string, data map[string]interface{}) {
	l.internal.Log(ctx, level, msg, data)
}
