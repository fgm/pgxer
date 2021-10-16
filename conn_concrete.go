package pgxer

import (
	"github.com/fgm/pgxer/pgconn"
)

type connConfig struct{}

func (cc *connConfig) Config() pgconn.Config {
	return nil
}

func (cc *connConfig) SetConfig(config pgconn.Config) ConnConfig {
	return nil
}

func (cc *connConfig) BuildStatementCache() BuildStatementCacheFunc {
	return nil
}

func (cc *connConfig) Logger() Logger {
	return nil
}
func (cc *connConfig) LogLevel() LogLevel {
	return LogLevelNone
}

func (cc *connConfig) PreferSimpleProtocol() bool {
	return false
}

func (cc *connConfig) SetBuildStatementCache(BuildStatementCacheFunc) ConnConfig {
	return nil
}

func (cc *connConfig) SetLogger(Logger) ConnConfig {
	return nil
}

func (cc *connConfig) SetLogLevel(LogLevel) ConnConfig {
	return nil
}

func (cc *connConfig) SetPreferSimpleProtocol(bool) ConnConfig {
	return nil
}

func (cc *connConfig) ConnString() string {
	return ""
}

func (cc *connConfig) Copy() ConnConfig {
	return nil
}
