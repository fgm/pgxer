package pgxpool

import (
	"context"
	"errors"

	basePgxPool "github.com/jackc/pgx/v4/pgxpool"

	"github.com/fgm/pgxer"
	"github.com/fgm/pgxer/pgconn"
)

type ConcretePool basePgxPool.Pool

var errUnimplemented = errors.New("Unimplemented")

func (p *ConcretePool) Acquire(ctx context.Context) (Conn, error) {
	// FIXME implement
	return nil, errUnimplemented
}
func (p *ConcretePool) AcquireFunc(ctx context.Context, f func(conn Conn) error) error {
	// FIXME implement
	return errUnimplemented
}
func (p *ConcretePool) AcquireAllIdle(ctx context.Context) []Conn {
	// FIXME implement
	return nil
}
func (p *ConcretePool) ConnConfig() pgxer.ConnConfig {
	// FIXME implement
	return nil
}
func (p *ConcretePool) Stat() Stat {
	// FIXME implement
	return nil
}
func (p *ConcretePool) Exec(ctx context.Context, sql string, arguments ...interface{}) (pgconn.CommandTag, error) {
	// FIXME implement
	return nil, errUnimplemented
}
func (p *ConcretePool) Query(ctx context.Context, sql string, args ...interface{}) (pgxer.Rows, error) {
	// FIXME implement
	return nil, errUnimplemented
}
func (p *ConcretePool) QueryRow(ctx context.Context, sql string, args ...interface{}) pgxer.Row {
	// FIXME implement
	return nil
}
func (p *ConcretePool) QueryFunc(ctx context.Context, sql string, args []interface{}, scans []interface{}, f func(pgxer.QueryFuncRow) error) (pgconn.CommandTag, error) {
	// FIXME implement
	return nil, errUnimplemented
}
func (p *ConcretePool) SendBatch(ctx context.Context, b pgxer.Batch) pgxer.BatchResults {
	// FIXME implement
	return nil
}
func (p *ConcretePool) Begin(ctx context.Context) (pgxer.Tx, error) {
	// FIXME implement
	return nil, errUnimplemented
}
func (p *ConcretePool) BeginTx(ctx context.Context, txOptions pgxer.TxOptions) (pgxer.Tx, error) {
	// FIXME implement
	return nil, errUnimplemented
}
func (p *ConcretePool) BeginFunc(ctx context.Context, f func(pgxer.Tx) error) error {
	// FIXME implement
	return errUnimplemented
}
func (p *ConcretePool) BeginTxFunc(ctx context.Context, txOptions pgxer.TxOptions, f func(pgxer.Tx) error) error {
	// FIXME implement
	return errUnimplemented
}
func (p *ConcretePool) CopyFrom(ctx context.Context, tableName pgxer.Identifier, columnNames []string, rowSrc pgxer.CopyFromSource) (int64, error) {
	// FIXME implement
	return 0, errUnimplemented
}
func (p *ConcretePool) Ping(ctx context.Context) error {
	// FIXME implement
	return errUnimplemented
}
