package pgxpool

import (
	"context"

	basePgxPool "github.com/jackc/pgx/v4/pgxpool"

	"github.com/fgm/pgxer"
	"github.com/fgm/pgxer/pgconn"
)

type ConcreteConn basePgxPool.Conn

func (cc *ConcreteConn) Begin(ctx context.Context) (pgxer.Tx, error) {
	return nil, errUnimplemented
}

func (cc *ConcreteConn) BeginFunc(ctx context.Context, f func(pgxer.Tx) error) error {
	return errUnimplemented
}

func (cc *ConcreteConn) BeginTx(ctx context.Context, txOptions pgxer.TxOptions) (pgxer.Tx, error) {
	return nil, errUnimplemented
}

func (cc *ConcreteConn) BeginTxFunc(ctx context.Context, txOptions pgxer.TxOptions, f func(pgxer.Tx) error) error {
	return errUnimplemented
}
func (cc *ConcreteConn) Conn() pgxer.Conn {
	return (*pgxer.ConcreteConn)((*basePgxPool.Conn)(cc).Conn())
}

func (cc *ConcreteConn) CopyFrom(ctx context.Context, tableName pgxer.Identifier, columnNames []string, rowSrc pgxer.CopyFromSource) (int64, error) {
	return 0, errUnimplemented
}

func (cc *ConcreteConn) Exec(ctx context.Context, sql string, arguments ...interface{}) (pgconn.CommandTag, error) {
	ct, err := (*basePgxPool.Conn)(cc).Exec(ctx, sql, arguments...)
	return pgconn.CommandTag(ct), err
}

func (cc *ConcreteConn) Ping(ctx context.Context) error {
	return errUnimplemented
}

func (cc *ConcreteConn) Query(ctx context.Context, sql string, args ...interface{}) (pgxer.Rows, error) {
	return nil, errUnimplemented
}

func (cc *ConcreteConn) QueryFunc(ctx context.Context, sql string, args []interface{}, scans []interface{}, f func(row pgxer.QueryFuncRow) error) (pgconn.CommandTag, error) {
	return nil, errUnimplemented
}

func (cc *ConcreteConn) QueryRow(ctx context.Context, sql string, args ...interface{}) pgxer.Row {
	return nil
}

func (cc *ConcreteConn) Release() {
}

func (cc *ConcreteConn) SendBatch(ctx context.Context, b pgxer.Batch) pgxer.BatchResults {
	return nil
}
