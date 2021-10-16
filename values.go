package pgxer

import "github.com/jackc/pgx/v4"

type SerializationError = pgx.SerializationError

const (
	TextFormatCode   = pgx.TextFormatCode
	BinaryFormatCode = pgx.BinaryFormatCode
)
