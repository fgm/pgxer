package pgconn

import "github.com/jackc/pgconn"

type PgError = pgconn.PgError

func SafeToRetry(err error) bool {
	return pgconn.SafeToRetry(err)
}

func Timeout(err error) bool {
	return pgconn.Timeout(err)
}
