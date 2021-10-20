package pgxer

import (
	basePgconn "github.com/jackc/pgconn"
	baseStmtCache "github.com/jackc/pgconn/stmtcache"

	"github.com/fgm/pgxer/pgconn"
	"github.com/fgm/pgxer/stmtcache"
)

type ConcretConnConfig struct{}

func (cc *ConcretConnConfig) Config() pgconn.Config {
	// FIXME implement
	return nil
}

func (cc *ConcretConnConfig) SetConfig(config pgconn.Config) ConnConfig {
	// FIXME implement
	return nil
}

func (cc *ConcretConnConfig) BuildStatementCache() BuildStatementCacheFunc {
	// FIXME implement
	return nil
}

func (cc *ConcretConnConfig) Logger() Logger {
	// FIXME implement
	return nil
}
func (cc *ConcretConnConfig) LogLevel() LogLevel {
	// FIXME implement
	return LogLevelNone
}

func (cc *ConcretConnConfig) PreferSimpleProtocol() bool {
	// FIXME implement
	return false
}

func (cc *ConcretConnConfig) SetBuildStatementCache(BuildStatementCacheFunc) ConnConfig {
	// FIXME implement
	return nil
}

func (cc *ConcretConnConfig) SetLogger(Logger) ConnConfig {
	// FIXME implement
	return nil
}

func (cc *ConcretConnConfig) SetLogLevel(LogLevel) ConnConfig {
	// FIXME implement
	return nil
}

func (cc *ConcretConnConfig) SetPreferSimpleProtocol(bool) ConnConfig {
	// FIXME implement
	return nil
}

func (cc *ConcretConnConfig) ConnString() string {
	// FIXME implement
	return ""
}

func (cc *ConcretConnConfig) Copy() ConnConfig {
	// FIXME implement
	return nil
}

func BuildStatementCacheFromBaseBuildStatementCache(func(conn *basePgconn.PgConn) baseStmtCache.Cache) func(conn pgconn.PgConn) stmtcache.Cache {
	// FIXME implement
	return nil
}
