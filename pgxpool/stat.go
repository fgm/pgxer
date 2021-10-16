package pgxpool

import (
	"time"
)

type Stat interface {
	AcquireCount() int64
	AcquiredConns() int32
	AcquireDuration() time.Duration
	CanceledAcquireCount() int64
	ConstructingConns() int32
	EmptyAcquireCount() int64
	IdleConns() int32
	MaxConns() int32
	TotalConns() int32
}
