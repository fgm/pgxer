package pgxer

import (
	basePgconn "github.com/jackc/pgconn"
	baseStmtCache "github.com/jackc/pgconn/stmtcache"
	basePgx "github.com/jackc/pgx/v4"

	"github.com/fgm/pgxer/pgconn"
	"github.com/fgm/pgxer/stmtcache"
)

type ConcreteConnConfig struct {
	*basePgx.ConnConfig
}

func (cc *ConcreteConnConfig) Config() pgconn.Config {
	// FIXME implement
	return nil
}

func (cc *ConcreteConnConfig) SetConfig(config pgconn.Config) ConnConfig {
	// FIXME implement
	return nil
}

func (cc *ConcreteConnConfig) BuildStatementCache() BuildStatementCacheFunc {
	// FIXME implement
	return nil
}

func (cc *ConcreteConnConfig) Logger() Logger {
	// FIXME implement
	return nil
}
func (cc *ConcreteConnConfig) LogLevel() LogLevel {
	// FIXME implement
	return LogLevelNone
}

func (cc *ConcreteConnConfig) PreferSimpleProtocol() bool {
	// FIXME implement
	return false
}

func (cc *ConcreteConnConfig) SetBuildStatementCache(BuildStatementCacheFunc) ConnConfig {
	// FIXME implement
	return nil
}

func (cc *ConcreteConnConfig) SetLogger(Logger) ConnConfig {
	// FIXME implement
	return nil
}

func (cc *ConcreteConnConfig) SetLogLevel(LogLevel) ConnConfig {
	// FIXME implement
	return nil
}

func (cc *ConcreteConnConfig) SetPreferSimpleProtocol(bool) ConnConfig {
	// FIXME implement
	return nil
}

func (cc *ConcreteConnConfig) ConnString() string {
	// FIXME implement
	return ""
}

func (cc *ConcreteConnConfig) Copy() ConnConfig {
	// FIXME implement
	return nil
}

func BuildStatementCacheFromBaseBuildStatementCache(func(conn *basePgconn.PgConn) baseStmtCache.Cache) func(conn pgconn.PgConn) stmtcache.Cache {
	// FIXME implement
	return nil
}

func baseBuildFrontendFromBuildFrontend(frontend pgconn.BuildFrontendFunc) basePgconn.BuildFrontendFunc {
	return nil
}

func baseFallbacksFromFallbacks([]pgconn.FallbackConfig) []*basePgconn.FallbackConfig {
	return nil
}

func baseValidateConnectFromValidateConnect(connect pgconn.ValidateConnectFunc) basePgconn.ValidateConnectFunc {
	return nil
}

func baseAfterConnectFromAfterConnect(connect pgconn.AfterConnectFunc) basePgconn.AfterConnectFunc {
	return nil
}

func baseOnNoticeFromOnNotice(notice pgconn.NoticeHandler) basePgconn.NoticeHandler {
	return nil
}

func baseOnNotificationFromOnNotification(notification pgconn.NotificationHandler) basePgconn.NotificationHandler {
	return nil
}
