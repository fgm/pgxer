package examples

import (
	"context"
	"github.com/jackc/pgx/v4"
)

type Pool interface {
	Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error)
}
