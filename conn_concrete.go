package pgxer

import (
	"context"

	basePgx "github.com/jackc/pgx/v4"

	"github.com/fgm/pgxer/pgconn"
	"github.com/fgm/pgxer/stmtcache"
)

type ConcreteConn basePgx.Conn

func (c *ConcreteConn) Begin(ctx context.Context) (Tx, error) {
	// TODO implement me
	panic("implement me")
}

func (c *ConcreteConn) BeginFunc(ctx context.Context, f func(Tx) error) (err error) {
	// TODO implement me
	panic("implement me")
}

func (c *ConcreteConn) BeginTx(ctx context.Context, txOptions TxOptions) (Tx, error) {
	// TODO implement me
	panic("implement me")
}

func (c *ConcreteConn) BeginTxFunc(ctx context.Context, txOptions TxOptions, f func(Tx) error) (err error) {
	// TODO implement me
	panic("implement me")
}

func (c *ConcreteConn) Close(ctx context.Context) error {
	// TODO implement me
	panic("implement me")
}

func (c *ConcreteConn) Config() ConnConfig {
	// TODO implement me
	panic("implement me")
}

func (c *ConcreteConn) CopyFrom(ctx context.Context, tableName Identifier, columnNames []string, rowSrc CopyFromSource) (int64, error) {
	// TODO implement me
	panic("implement me")
}

func (c *ConcreteConn) Deallocate(ctx context.Context, name string) error {
	// TODO implement me
	panic("implement me")
}

func (c *ConcreteConn) Exec(ctx context.Context, sql string, arguments ...interface{}) (pgconn.CommandTag, error) {
	// TODO implement me
	panic("implement me")
}

func (c *ConcreteConn) IsClosed() bool {
	// TODO implement me
	panic("implement me")
}

func (c *ConcreteConn) PgConn() pgconn.PgConn {
	// TODO implement me
	panic("implement me")
}

func (c *ConcreteConn) Ping(ctx context.Context) error {
	// TODO implement me
	panic("implement me")
}

func (c *ConcreteConn) Prepare(ctx context.Context, name, sql string) (sd pgconn.StatementDescription, err error) {
	// TODO implement me
	panic("implement me")
}

func (c *ConcreteConn) Query(ctx context.Context, sql string, args ...interface{}) (Rows, error) {
	// TODO implement me
	panic("implement me")
}

func (c *ConcreteConn) QueryFunc(ctx context.Context, sql string, args []interface{}, scans []interface{}, f func(QueryFuncRow) error) (pgconn.CommandTag, error) {
	// TODO implement me
	panic("implement me")
}

func (c *ConcreteConn) QueryRow(ctx context.Context, sql string, args ...interface{}) Row {
	// TODO implement me
	panic("implement me")
}

func (c *ConcreteConn) SendBatch(ctx context.Context, b Batch) BatchResults {
	// TODO implement me
	panic("implement me")
}

func (c *ConcreteConn) StatementCache() stmtcache.Cache {
	// TODO implement me
	panic("implement me")
}

func (c *ConcreteConn) WaitForNotification(ctx context.Context) (pgconn.Notification, error) {
	// WaitForNotification is a pointer method, so it needs an addressable.
	n, err := (*basePgx.Conn)(c).WaitForNotification(ctx)
	return &pgconn.ConcreteNotification{Notification: n}, err
}
