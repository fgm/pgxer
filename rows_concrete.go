package pgxer

import (
	"github.com/jackc/pgx/v4"

	"github.com/fgm/pgxer/pgconn"
)

// ConcreteRow method Scan is not needed: the base implementation is directly usable.
type ConcreteRow struct {
	pgx.Row
}

// ConcreteRows methods are not needed except for CommandTag: the base implementation is directly usable.
type ConcreteRows struct {
	pgx.Rows
}

func (c ConcreteRows) CommandTag() pgconn.CommandTag {
	return c.Rows.CommandTag()
}
