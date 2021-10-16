package pgxer

import (
	"context"

	"github.com/jackc/pgx/v4"
)

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

const (
	LargeObjectModeWrite = LargeObjectMode(pgx.LargeObjectModeWrite)
	LargeObjectModeRead  = LargeObjectMode(pgx.LargeObjectModeRead)
)
