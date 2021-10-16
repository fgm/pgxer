package pgxer

import (
	"context"

	"github.com/jackc/pgx/v4"
)

type Logger interface {
	Log(ctx context.Context, level LogLevel, msg string, data map[string]interface{})
}

type LogLevel = pgx.LogLevel

const (
	LogLevelTrace = LogLevel(pgx.LogLevelTrace)
	LogLevelDebug = LogLevel(pgx.LogLevelDebug)
	LogLevelInfo  = LogLevel(pgx.LogLevelInfo)
	LogLevelWarn  = LogLevel(pgx.LogLevelWarn)
	LogLevelError = LogLevel(pgx.LogLevelError)
	LogLevelNone  = LogLevel(pgx.LogLevelNone)
)

func LogLevelFromString(s string) (LogLevel, error) {
	return pgx.LogLevelFromString(s)
}
