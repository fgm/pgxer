package pgxpool

import (
	"context"

	"github.com/fgm/pgxer"
	"github.com/fgm/pgxer/pgconn"
)

type Tx interface {
	Begin(ctx context.Context) (pgxer.Tx, error)
	BeginFunc(ctx context.Context, f func(pgxer.Tx) error) error
	Commit(ctx context.Context) error
	Conn() pgxer.Conn
	CopyFrom(ctx context.Context, tableName pgxer.Identifier, columnNames []string, rowSrc pgxer.CopyFromSource) (int64, error)
	Exec(ctx context.Context, sql string, arguments ...interface{}) (pgconn.CommandTag, error)
	LargeObjects() pgxer.LargeObjects
	Prepare(ctx context.Context, name, sql string) (pgconn.StatementDescription, error)
	Query(ctx context.Context, sql string, args ...interface{}) (pgxer.Rows, error)
	QueryFunc(ctx context.Context, sql string, args []interface{}, scans []interface{}, f func(pgxer.QueryFuncRow) error) (pgconn.CommandTag, error)
	QueryRow(ctx context.Context, sql string, args ...interface{}) pgxer.Row
	Rollback(ctx context.Context) error
	SendBatch(ctx context.Context, b pgxer.Batch) pgxer.BatchResults
}
