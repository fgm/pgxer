package pgconn

import (
	basePgconn "github.com/jackc/pgconn"
)

type ConcreteConfig Config

func PgconnConfigFromBasePgconnConfig(basePgconn.Config) Config {
	// FIXME implement
	return nil
}
