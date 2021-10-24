package pgxpool

import (
	"context"
	"time"

	basePgx "github.com/jackc/pgx/v4"
	basePgxPool "github.com/jackc/pgx/v4/pgxpool"

	"github.com/fgm/pgxer"
	"github.com/fgm/pgxer/internal"
)

type ConcreteConfig struct {
	basePgxPool.Config
}

func (c *ConcreteConfig) ConnConfig() pgxer.ConnConfig {
	// FIXME implement
	return &pgxer.ConcreteConnConfig{}
}

func (c *ConcreteConfig) SetConnConfig(config pgxer.ConnConfig) Config {
	// FIXME implement
	return c
}

func (c *ConcreteConfig) AfterConnect(context.Context, pgxer.Conn) error {
	// FIXME implement
	return internal.ErrUnimplemented
}
func (c *ConcreteConfig) AfterRelease(pgxer.Conn) bool {
	// FIXME implement
	return false
}
func (c *ConcreteConfig) BeforeAcquire(context.Context, pgxer.Conn) bool {
	// FIXME implement
	return false
}
func (c *ConcreteConfig) BeforeConnect(context.Context, pgxer.ConnConfig) error {
	// FIXME implement
	return internal.ErrUnimplemented
}

func (c *ConcreteConfig) SetAfterConnect(func(context.Context, pgxer.Conn) error) Config {
	// FIXME implement
	return c
}
func (c *ConcreteConfig) SetAfterRelease(func(pgxer.Conn) bool) Config {
	// FIXME implement
	return c
}

func (c *ConcreteConfig) SetBeforeAcquire(func(context.Context, pgxer.Conn) bool) Config {
	// FIXME implement
	return c
}

func (c *ConcreteConfig) SetBeforeConnect(func(context.Context, pgxer.ConnConfig) error) Config {
	// FIXME implement
	return c
}

func (c *ConcreteConfig) HealthCheckPeriod() time.Duration {
	// FIXME implement
	return 0
}
func (c *ConcreteConfig) LazyConnect() bool {
	// FIXME implement
	return false
}
func (c *ConcreteConfig) MaxConnIdleTime() time.Duration {
	// FIXME implement
	return 0
}
func (c *ConcreteConfig) MaxConnLifetime() time.Duration {
	// FIXME implement
	return 0
}
func (c *ConcreteConfig) MaxConns() int32 {
	// FIXME implement
	return 0
}
func (c *ConcreteConfig) MinConns() int32 {
	// FIXME implement
	return 0
}

func (c *ConcreteConfig) SetHealthCheckPeriod(time.Duration) Config {
	// FIXME implement
	return c
}

func (c *ConcreteConfig) SetLazyConnect(bool) Config {
	// FIXME implement
	return c
}

func (c *ConcreteConfig) SetMaxConnIdleTime(time.Duration) Config {
	// FIXME implement
	return c
}

func (c *ConcreteConfig) SetMaxConnLifetime(time.Duration) Config {
	// FIXME implement
	return c
}

func (c *ConcreteConfig) SetMaxConns(int32) Config {
	// FIXME implement
	return c
}

func (c *ConcreteConfig) SetMinConns(int32) Config {
	// FIXME implement
	return c
}

func BaseConnConfigFromConnConfig(c pgxer.ConnConfig) *basePgx.ConnConfig {
	// FIXME implement
	cc, _ := basePgx.ParseConfig("")
	return cc
}

func BaseAfterConnectFromAfterConnect(func(context.Context, pgxer.Conn) error) func(context.Context, *basePgx.Conn) error {
	// FIXME implement
	return nil
}

func BaseAfterReleaseFromAfterRelease(func(conn pgxer.Conn) bool) func(conn *basePgx.Conn) bool {
	// FIXME implement
	return nil
}

func BaseBeforeAcquireFromBeforeAcquire(func(ctx context.Context, conn pgxer.Conn) bool) func(ctx context.Context, conn *basePgx.Conn) bool {
	// FIXME implement
	return nil
}

func BaseBeforeConnectFromBeforeConnect(func(ctx context.Context, connConfig pgxer.ConnConfig) error) func(context.Context, *basePgx.ConnConfig) error {
	// FIXME implement
	return nil
}

func AfterConnectFromBaseAfterConnect(func(ctx context.Context, conn *basePgx.Conn) error) func(context.Context, pgxer.Conn) error {
	// FIXME implement
	return nil
}

func AfterReleaseFromBaseAfterRelease(func(conn *basePgx.Conn) bool) func(pgxer.Conn) bool {
	// FIXME implement
	return nil
}

func BeforeAcquireFromBaseBeforeAcquire(func(ctx context.Context, conn *basePgx.Conn) bool) func(ctx context.Context, conn pgxer.Conn) bool {
	// FIXME implement
	return nil
}

func BeforeConnectFromBaseBeforeConnect(func(ctx context.Context, config *basePgx.ConnConfig) error) func(ctx context.Context, config pgxer.ConnConfig) error {
	// FIXME implement
	return nil
}

func ConnConfigFromBaseConnConfig(config *basePgx.ConnConfig) pgxer.ConnConfig {
	// FIXME implement
	return nil
}
