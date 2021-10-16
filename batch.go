package pgxer

import "github.com/fgm/pgxer/pgconn"

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
