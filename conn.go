package pgxer

import (
	"context"

	"github.com/jackc/pgproto3/v2"
	"github.com/jackc/pgx/v4"

	"github.com/fgm/pgxer/pgconn"
	"github.com/fgm/pgxer/stmtcache"
)

type BuildStatementCacheFunc func(conn pgconn.PgConn) stmtcache.Cache

type Conn interface {
	Begin(ctx context.Context) (Tx, error)
	BeginFunc(ctx context.Context, f func(Tx) error) (err error)
	BeginTx(ctx context.Context, txOptions TxOptions) (Tx, error)
	BeginTxFunc(ctx context.Context, txOptions TxOptions, f func(Tx) error) (err error)
	Close(ctx context.Context) error
	Config() ConnConfig
	CopyFrom(ctx context.Context, tableName Identifier, columnNames []string, rowSrc CopyFromSource) (int64, error)
	Deallocate(ctx context.Context, name string) error
	Exec(ctx context.Context, sql string, arguments ...interface{}) (pgconn.CommandTag, error)
	IsClosed() bool
	PgConn() pgconn.PgConn
	Ping(ctx context.Context) error
	Prepare(ctx context.Context, name, sql string) (sd pgconn.StatementDescription, err error)
	Query(ctx context.Context, sql string, args ...interface{}) (Rows, error)
	QueryFunc(ctx context.Context, sql string, args []interface{}, scans []interface{}, f func(QueryFuncRow) error) (pgconn.CommandTag, error)
	QueryRow(ctx context.Context, sql string, args ...interface{}) Row
	SendBatch(ctx context.Context, b Batch) BatchResults
	StatementCache() stmtcache.Cache
	WaitForNotification(ctx context.Context) (pgconn.Notification, error)
}

type ConnConfig interface {
	Config() pgconn.Config
	SetConfig(config pgconn.Config) ConnConfig

	BuildStatementCache() BuildStatementCacheFunc
	Logger() Logger
	LogLevel() LogLevel
	PreferSimpleProtocol() bool

	SetBuildStatementCache(BuildStatementCacheFunc) ConnConfig
	SetLogger(Logger) ConnConfig
	SetLogLevel(LogLevel) ConnConfig
	SetPreferSimpleProtocol(bool) ConnConfig

	ConnString() string
	Copy() ConnConfig
}

type Identifier = pgx.Identifier

type QueryFuncRow interface {
	FieldDescriptions() []pgproto3.FieldDescription
	RawValues() [][]byte
}

type QueryResultFormats = pgx.QueryResultFormats
type QueryResultFormatsByOID = pgx.QueryResultFormatsByOID
type QuerySimpleProtocol = pgx.QuerySimpleProtocol

func Connect(ctx context.Context, connString string) (Conn, error) {
	return nil, nil
}

func ConnectConfig(ctx context.Context, connConfig ConnConfig) (Conn, error) {
	return nil, nil
}

func ParseConfig(connString string) (ConnConfig, error) {
	return nil, nil
}

var ErrNoRows = pgx.ErrNoRows
var ErrInvalidLogLevel = pgx.ErrInvalidLogLevel
