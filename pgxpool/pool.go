package pgxpool

import (
	"context"
	"time"

	"github.com/fgm/pgxer"
	"github.com/fgm/pgxer/pgconn"
)

type Config interface {
	ConnConfig() pgxer.ConnConfig
	SetConnConfig(config pgxer.ConnConfig) Config

	AfterConnect(context.Context, pgxer.Conn) error
	AfterRelease(pgxer.Conn) bool
	BeforeAcquire(context.Context, pgxer.Conn) bool
	BeforeConnect(context.Context, pgxer.ConnConfig) error

	SetAfterConnect(func(context.Context, pgxer.Conn) error) Config
	SetAfterRelease(func(pgxer.Conn) bool) Config
	SetBeforeAcquire(func(context.Context, pgxer.Conn) bool) Config
	SetBeforeConnect(func(context.Context, pgxer.ConnConfig) error) Config

	HealthCheckPeriod() time.Duration
	LazyConnect() bool
	MaxConnIdleTime() time.Duration
	MaxConnLifetime() time.Duration
	MaxConns() int32
	MinConns() int32

	SetHealthCheckPeriod(time.Duration) Config
	SetLazyConnect(bool) Config
	SetMaxConnIdleTime(time.Duration) Config
	SetMaxConnLifetime(time.Duration) Config
	SetMaxConns(int32) Config
	SetMinConns(int32) Config
}

type Pool interface {
	Acquire(ctx context.Context) (Conn, error)
	AcquireFunc(ctx context.Context, f func(Conn) error) error
	AcquireAllIdle(ctx context.Context) []Conn
	ConnConfig() pgxer.ConnConfig
	Stat() Stat
	Exec(ctx context.Context, sql string, arguments ...interface{}) (pgconn.CommandTag, error)
	Query(ctx context.Context, sql string, args ...interface{}) (pgxer.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...interface{}) pgxer.Row
	QueryFunc(ctx context.Context, sql string, args []interface{}, scans []interface{}, f func(pgxer.QueryFuncRow) error) (pgconn.CommandTag, error)
	SendBatch(ctx context.Context, b pgxer.Batch) pgxer.BatchResults
	Begin(ctx context.Context) (pgxer.Tx, error)
	BeginTx(ctx context.Context, txOptions pgxer.TxOptions) (pgxer.Tx, error)
	BeginFunc(ctx context.Context, f func(pgxer.Tx) error) error
	BeginTxFunc(ctx context.Context, txOptions pgxer.TxOptions, f func(pgxer.Tx) error) error
	CopyFrom(ctx context.Context, tableName pgxer.Identifier, columnNames []string, rowSrc pgxer.CopyFromSource) (int64, error)
	Ping(ctx context.Context) error
}

func Connect(ctx context.Context, connString string) (Pool, error) {
	return nil, nil
}

func ConnectConfig(ctx context.Context, config Config) (Pool, error) {
	return nil, nil
}

func ParseConfig(connString string) (Config, error) {
	return nil, nil
}
