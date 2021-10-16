package pgxer

import (
	"github.com/jackc/pgproto3/v2"
	"github.com/jackc/pgtype"

	"github.com/fgm/pgxer/pgconn"
)

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

	SetColumnIndex(int) ScanArgError
	SetErr(error) ScanArgError

	Error() string
	Unwrap() error
}

func ScanRow(connInfo *pgtype.ConnInfo, fieldDescriptions []pgproto3.FieldDescription, values [][]byte, dest ...interface{}) error {
	return nil
}
