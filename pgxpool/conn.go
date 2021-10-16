package pgxpool

import (
	"context"

	"github.com/fgm/pgxer"
	"github.com/fgm/pgxer/pgconn"
)

// Conn is an acquired pgxer.Conn from a pgxpool.Pool.
type Conn interface {
	Begin(ctx context.Context) (pgxer.Tx, error)
	BeginFunc(ctx context.Context, f func(pgxer.Tx) error) error
	BeginTx(ctx context.Context, txOptions pgxer.TxOptions) (pgxer.Tx, error)
	BeginTxFunc(ctx context.Context, txOptions pgxer.TxOptions, f func(pgxer.Tx) error) error
	Conn() pgxer.Conn
	CopyFrom(ctx context.Context, tableName pgxer.Identifier, columnNames []string, rowSrc pgxer.CopyFromSource) (int64, error)
	Exec(ctx context.Context, sql string, arguments ...interface{}) (pgconn.CommandTag, error)
	Ping(ctx context.Context) error
	Query(ctx context.Context, sql string, args ...interface{}) (pgxer.Rows, error)
	QueryFunc(ctx context.Context, sql string, args []interface{}, scans []interface{}, f func(row pgxer.QueryFuncRow) error) (pgconn.CommandTag, error)
	QueryRow(ctx context.Context, sql string, args ...interface{}) pgxer.Row
	Release()
	SendBatch(ctx context.Context, b pgxer.Batch) pgxer.BatchResults
}
