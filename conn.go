package pgxer

import (
	"context"

	basePgconn "github.com/jackc/pgconn"
	"github.com/jackc/pgproto3/v2"
	basePgx "github.com/jackc/pgx/v4"

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

type Identifier = basePgx.Identifier

type QueryFuncRow interface {
	FieldDescriptions() []pgproto3.FieldDescription
	RawValues() [][]byte
}

type QueryResultFormats = basePgx.QueryResultFormats
type QueryResultFormatsByOID = basePgx.QueryResultFormatsByOID
type QuerySimpleProtocol = basePgx.QuerySimpleProtocol

func Connect(ctx context.Context, connString string) (Conn, error) {
	bc, err := basePgx.Connect(ctx, connString)
	return (*ConcreteConn)(bc), err
}

func ConnectConfig(ctx context.Context, connConfig ConnConfig) (Conn, error) {
	pgcc := connConfig.Config()
	cc := &basePgx.ConnConfig{
		Config: basePgconn.Config{
			Host:            pgcc.Host(),
			Port:            pgcc.Port(),
			Database:        pgcc.Database(),
			User:            pgcc.User(),
			Password:        pgcc.Password(),
			TLSConfig:       pgcc.TLSConfig(),
			ConnectTimeout:  pgcc.ConnectTimeout(),
			DialFunc:        pgcc.DialFunc(),
			LookupFunc:      pgcc.LookupFunc(),
			BuildFrontend:   baseBuildFrontendFromBuildFrontend(pgcc.BuildFrontend()),
			RuntimeParams:   pgcc.RuntimeParams(),
			Fallbacks:       baseFallbacksFromFallbacks(pgcc.Fallbacks()),
			ValidateConnect: baseValidateConnectFromValidateConnect(pgcc.ValidateConnect()),
			AfterConnect:    baseAfterConnectFromAfterConnect(pgcc.AfterConnect()),
			OnNotice:        baseOnNoticeFromOnNotice(pgcc.OnNotice()),
			OnNotification:  baseOnNotificationFromOnNotification(pgcc.OnNotification()),
		},
		Logger:               connConfig.Logger(),
		LogLevel:             connConfig.LogLevel(),
		BuildStatementCache:  nil,
		PreferSimpleProtocol: connConfig.PreferSimpleProtocol(),
	}
	bc, err := basePgx.ConnectConfig(ctx, cc)
	return (*ConcreteConn)(bc), err

}

func ParseConfig(connString string) (ConnConfig, error) {
	baseCC, err := basePgx.ParseConfig(connString)
	cc := &ConcreteConnConfig{ConnConfig: baseCC}
	return cc, err
}

var ErrNoRows = basePgx.ErrNoRows
var ErrInvalidLogLevel = basePgx.ErrInvalidLogLevel
