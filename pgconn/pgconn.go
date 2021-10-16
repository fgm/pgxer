package pgconn

import (
	"context"
	"io"
	"net"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgproto3/v2"
)

type Batch interface {
	ExecParams(sql string, paramValues [][]byte, paramOIDs []uint32, paramFormats []int16, resultFormats []int16)
	ExecPrepared(stmtName string, paramValues [][]byte, paramFormats []int16, resultFormats []int16)
}

type BuildFrontendFunc func(r io.Reader, w io.Writer) Frontend

type CommandTag interface {
	Delete() bool
	Insert() bool
	RowsAffected() int64
	Select() bool
	String() string
	Update() bool
}

type DialFunc = pgconn.DialFunc

type Frontend = pgconn.Frontend

type HijackedConn interface {
	Config() Config
	Conn() net.Conn
	Frontend() Frontend
	ParameterStatuses() map[string]string
	PID() uint32
	SecretKey() uint32
	TxStatus() byte

	SetConfig(Config) HijackedConn
	SetConn(net.Conn) HijackedConn
	SetFrontend(Frontend) HijackedConn
	SetParameterStatuses(map[string]string) HijackedConn
	SetPID(uint32) HijackedConn
	SetSecretKey(uint32) HijackedConn
	SetTxStatus(byte) HijackedConn
}

type LookupFunc = pgconn.LookupFunc

type MultiResultReader interface {
	Close() error
	NextResult() bool
	ReadAll() ([]*Result, error)
	ResultReader() *ResultReader
}

type Notice = pgconn.Notice

type NoticeHandler func(PgConn, Notice)

type Notification interface {
	Channel() string
	Payload() string
	PID() uint32

	SetChannel(string) Notification
	SetPayload(string) Notification
	SetPID(uint32) Notification
}

type NotificationHandler func(PgConn, Notification)

type PgConn interface {
	CancelRequest(ctx context.Context) error
	CleanupDone() chan (struct{})
	Close(ctx context.Context) error
	Conn() net.Conn
	CopyFrom(ctx context.Context, r io.Reader, sql string) (CommandTag, error)
	CopyTo(ctx context.Context, w io.Writer, sql string) (CommandTag, error)
	EscapeString(s string) (string, error)
	Exec(ctx context.Context, sql string) *MultiResultReader
	ExecBatch(ctx context.Context, batch *Batch) *MultiResultReader
	ExecParams(ctx context.Context, sql string, paramValues [][]byte, paramOIDs []uint32, paramFormats []int16, resultFormats []int16) *ResultReader
	ExecPrepared(ctx context.Context, stmtName string, paramValues [][]byte, paramFormats []int16, resultFormats []int16) *ResultReader
	Hijack() (*HijackedConn, error)
	IsBusy() bool
	IsClosed() bool
	ParameterStatus(key string) string
	PID() uint32
	Prepare(ctx context.Context, name, sql string, paramOIDs []uint32) (*StatementDescription, error)
	ReceiveMessage(ctx context.Context) (pgproto3.BackendMessage, error)
	ReceiveResults(ctx context.Context) *MultiResultReader
	SecretKey() uint32
	SendBytes(ctx context.Context, buf []byte) error
	TxStatus() byte
	WaitForNotification(ctx context.Context) error
}

type Result interface {
	CommandTag() CommandTag
	Err() error
	FieldDescriptions() []pgproto3.FieldDescription
	Rows() [][][]byte

	SetCommandTag(CommandTag) Result
	SetErr(error) Result
	SetFieldDescriptions([]pgproto3.FieldDescription) Result
	SetRows([][][]byte) Result
}

type ResultReader interface {
	Close() (CommandTag, error)
	FieldDescriptions() []pgproto3.FieldDescription
	NextRow() bool
	Read() Result
	Values() [][]byte
}
type StatementDescription interface {
	Fields() []pgproto3.FieldDescription
	Name() string
	ParamOIDs() []uint32
	SQL() string

	SetFields([]pgproto3.FieldDescription) StatementDescription
	SetName(string) StatementDescription
	SetParamOIDs([]uint32) StatementDescription
	SetSQL(string) StatementDescription
}

func Connect(ctx context.Context, connString string) (PgConn, error) {
	return nil, nil
}

func ConnectConfig(ctx context.Context, config Config) (pgConn PgConn, err error) {
	return nil, nil
}

func Construct(hc HijackedConn) (PgConn, error) {
	return nil, nil
}

func ErrorResponseToPgError(msg *pgproto3.ErrorResponse) *PgError {
	return nil
}
