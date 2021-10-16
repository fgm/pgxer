package pgxer

import (
	"context"

	"github.com/jackc/pgproto3/v2"
	"github.com/jackc/pgx/v4"

	"github.com/fgm/pgxer/pgconn"
	"github.com/fgm/pgxer/stmtcache"
)

type Batch interface {
	Len() int
	Queue(query string, arguments ...interface{})
}

type BatchResults interface {
	Close() error
	Exec() (pgconn.CommandTag, error)
	Query() (Rows, error)
	QueryRow() Row
}

type BuildStatementCacheFunc func(conn *pgconn.PgConn) stmtcache.Cache

type Conn interface {
	Begin(ctx context.Context) (Tx, error)
	BeginFunc(ctx context.Context, f func(Tx) error) (err error)
	BeginTx(ctx context.Context, txOptions TxOptions) (Tx, error)
	BeginTxFunc(ctx context.Context, txOptions TxOptions, f func(Tx) error) (err error)
	Close(ctx context.Context) error
	Config() *ConnConfig
	CopyFrom(ctx context.Context, tableName Identifier, columnNames []string, rowSrc CopyFromSource) (int64, error)
	Deallocate(ctx context.Context, name string) error
	Exec(ctx context.Context, sql string, arguments ...interface{}) (pgconn.CommandTag, error)
	IsClosed() bool
	PgConn() *pgconn.PgConn
	Ping(ctx context.Context) error
	Prepare(ctx context.Context, name, sql string) (sd *pgconn.StatementDescription, err error)
	Query(ctx context.Context, sql string, args ...interface{}) (Rows, error)
	QueryFunc(ctx context.Context, sql string, args []interface{}, scans []interface{}, f func(QueryFuncRow) error) (pgconn.CommandTag, error)
	QueryRow(ctx context.Context, sql string, args ...interface{}) Row
	SendBatch(ctx context.Context, b *Batch) BatchResults
	StatementCache() stmtcache.Cache
	WaitForNotification(ctx context.Context) (*pgconn.Notification, error)
}

type ConnConfig interface {
	pgconn.Config

	BuildStatementCache() BuildStatementCacheFunc
	Logger() Logger
	LogLevel() LogLevel
	PreferSimpleProtocol() bool

	SetBuildStatementCache(BuildStatementCacheFunc)
	SetLogger(Logger)
	SetLogLevel(LogLevel)
	SetPreferSimpleProtocol(bool)

	ConnString() string
	Copy() *ConnConfig
}

type CopyFromSource interface {
	Err() error
	Next() bool
	Values() ([]interface{}, error)
}

type Identifier interface {
	Sanitize() string
}

type LargeObject interface {
	Close() error
	Read(p []byte) (int, error)
	Seek(offset int64, whence int) (n int64, err error)
	Tell() (n int64, err error)
	Truncate(size int64) (err error)
	Write(p []byte) (int, error)
}

type LargeObjectMode = pgx.LargeObjectMode

type LargeObjects interface {
	Create(ctx context.Context, oid uint32) (uint32, error)
	Open(ctx context.Context, oid uint32, mode LargeObjectMode) (*LargeObject, error)
	Unlink(ctx context.Context, oid uint32) error
}

type Logger interface {
	Log(ctx context.Context, level LogLevel, msg string, data map[string]interface{})
}

type LogLevel = pgx.LogLevel

type QueryFuncRow interface {
	FieldDescriptions() []pgproto3.FieldDescription
	RawValues() [][]byte
}

type QueryResultFormats = pgx.QueryResultFormats
type QueryResultFormatsByOID = pgx.QueryResultFormatsByOID
type QuerySimpleProtocol = pgx.QuerySimpleProtocol

type Row interface {
	Scan(dest ...interface{}) error
}

type Rows interface {
	Close()
	CommandTag() pgconn.CommandTag
	Err() error
	FieldDescriptions() []pgproto3.FieldDescription
	Next() bool
	RawValues() [][]byte
	Scan(dest ...interface{}) error
	Values() ([]interface{}, error)
}

type ScanArgError interface {
	ColumnIndex() int
	Err() error

	SetColumnIndex(int)
	SetErr(error)

	Error() string
	Unwrap() error
}

type SerializationError = pgx.SerializationError

type Tx interface {
	Begin(ctx context.Context) (Tx, error)
	BeginFunc(ctx context.Context, f func(Tx) error) (err error)
	Commit(ctx context.Context) error
	Conn() *Conn
	CopyFrom(ctx context.Context, tableName Identifier, columnNames []string, rowSrc CopyFromSource) (int64, error)
	Exec(ctx context.Context, sql string, arguments ...interface{}) (commandTag pgconn.CommandTag, err error)
	LargeObjects() LargeObjects
	Prepare(ctx context.Context, name, sql string) (*pgconn.StatementDescription, error)
	Query(ctx context.Context, sql string, args ...interface{}) (Rows, error)
	QueryFunc(ctx context.Context, sql string, args []interface{}, scans []interface{}, f func(QueryFuncRow) error) (pgconn.CommandTag, error)
	QueryRow(ctx context.Context, sql string, args ...interface{}) Row
	Rollback(ctx context.Context) error
	SendBatch(ctx context.Context, b *Batch) BatchResults
}

type TxAccessMode = pgx.TxAccessMode
type TxDeferrableMode = pgx.TxDeferrableMode
type TxIsoLevel = pgx.TxIsoLevel

type TxOptions interface {
	AccessMode() TxAccessMode
	DeferrableMode() TxDeferrableMode
	IsoLevel() TxIsoLevel

	SetAccessMode(TxAccessMode)
	SetDeferrableMode(TxDeferrableMode)
	SetIsoLevel(TxIsoLevel)
}

const (
	LargeObjectModeWrite = LargeObjectMode(pgx.LargeObjectModeWrite)
	LargeObjectModeRead  = LargeObjectMode(pgx.LargeObjectModeRead)
)

const (
	LogLevelTrace = LogLevel(pgx.LogLevelTrace)
	LogLevelDebug = LogLevel(pgx.LogLevelDebug)
	LogLevelInfo  = LogLevel(pgx.LogLevelInfo)
	LogLevelWarn  = LogLevel(pgx.LogLevelWarn)
	LogLevelError = LogLevel(pgx.LogLevelError)
	LogLevelNone  = LogLevel(pgx.LogLevelNone)
)

const (
	ReadWrite = TxAccessMode(pgx.ReadWrite)
	ReadOnly  = TxAccessMode(pgx.ReadOnly)
)

const (
	Deferrable    = TxDeferrableMode(pgx.Deferrable)
	NotDeferrable = TxDeferrableMode(pgx.NotDeferrable)
)

const (
	Serializable    = TxIsoLevel(pgx.Serializable)
	RepeatableRead  = TxIsoLevel(pgx.RepeatableRead)
	ReadCommitted   = TxIsoLevel(pgx.ReadCommitted)
	ReadUncommitted = TxIsoLevel(pgx.ReadUncommitted)
)
