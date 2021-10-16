package pgconn

import (
	"context"
	"crypto/tls"
	"time"

	"github.com/jackc/pgconn"
)

type AfterConnectFunc func(ctx context.Context, pgconn PgConn) error

type Config interface {
	Host() string
	Port() uint16
	Database() string
	User() string
	Password() string
	TLSConfig() *tls.Config // nil disables TLS
	ConnectTimeout() time.Duration
	DialFunc() DialFunc     // e.g. net.Dialer.DialContext
	LookupFunc() LookupFunc // e.g. net.Resolver.LookupHost
	BuildFrontend() BuildFrontendFunc
	RuntimeParams() map[string]string // Run-time parameters to set on connection as session default values (e.g. search_path or application_name)

	Fallbacks() []FallbackConfig

	// ValidateConnect is called during a connection attempt after a successful authentication with the PostgreSQL server.
	// It can be used to validate that the server is acceptable. If this returns an error the connection is closed and the next
	// fallback config is tried. This allows implementing high availability behavior such as libpq does with target_session_attrs.
	ValidateConnect() ValidateConnectFunc

	// AfterConnect is called after ValidateConnect. It can be used to set up the connection (e.g. Set session variables
	// or prepare statements). If this returns an error the connection attempt fails.
	AfterConnect() AfterConnectFunc

	// OnNotice is a callback function called when a notice response is received.
	OnNotice() NoticeHandler

	// OnNotification is a callback function called when a notification from the LISTEN/NOTIFY system is received.
	OnNotification() NotificationHandler

	Copy() Config
}

type FallbackConfig interface {
	Host() string
	Port() uint16
	TLSConfig() *tls.Config

	SetHost(string) FallbackConfig
	SetPort(uint16) FallbackConfig
	SetTLSConfig(config *tls.Config) FallbackConfig
}

type ValidateConnectFunc func(ctx context.Context, pgconn PgConn) error

func NetworkAddress(host string, port uint16) (network, address string) {
	return pgconn.NetworkAddress(host, port)
}

func ParseConfig(connString string) (Config, error) {
	return nil, nil
}
