package pgxer

import (
	"context"

	"github.com/jackc/pgx/v4"

	"github.com/fgm/pgxer/pgconn"
)

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

	SetAccessMode(TxAccessMode) TxOptions
	SetDeferrableMode(TxDeferrableMode) TxOptions
	SetIsoLevel(TxIsoLevel) TxOptions
}

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

var ErrTxClosed = pgx.ErrTxClosed
var ErrTxCommitRollback = pgx.ErrTxCommitRollback
