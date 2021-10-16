package stmtcache

import (
	"context"

	"github.com/fgm/pgxer/pgconn"
)

type Cache interface {
	Get(ctx context.Context, sql string) (*pgconn.StatementDescription, error)
	Clear(ctx context.Context) error
	StatementErrored(sql string, err error)
	Len() int
	Cap() int
	Mode() int
}
