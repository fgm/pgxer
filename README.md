# PGXer

## Purpose

_pgxer_ is a thin wrapper around [jackc/pgx](https://github.com/jackc/pgx) that
exposes the _pgx_ 4.x API, but using interface types instead of concrete types,
to make it easier to mock DB access at any level from the `Pool` downwards.

In v4, _pgx_ provides a mix of interfaces (e.g. `Rows`) and concrete types
(e.g. `Pool`), making it essentially impossible to write tests relying on a DB
without accessing the DB, causing tests to be slow, hard to run concurrently,
and require external setup.

_pgxer_ uses identical public type names, but modifies the signatures of all
public methods to take and return interfaces to support mocking all types; and
includes a mock implementation for all of them.

Where public fields exists, each one is replaced by an equivalent methods pair,
to support mocking: a getter and a fluent setter.


## Application Use

For the simpler cases, just applying three changes should be sufficient, as the
[examples](examples/), converted from the original _pgx_ examples, demonstrate:

- Import paths:
    - `import "github.com/jackc/pgx/v4"` -> `import pgx "github.com/fgm/pgxer"`
    - `import "github.com/jackc/pgx/v4/pgxpool"` -> `import "github.com/fgm/pgxer/pgxpool"`
- Concrete types and their pointers:
    - `var conn *pgx.Conn` -> `var conn pgxer.Conn`
    - `var pool *pgxpool.Pool` -> `var pool pgxpool.Pool`
- Fields to equivalent getter/setter methods:
    - After `poolConfig, err := pgxpool.ParseConfig(os.Getenv("DATABASE_URL"))`
    - `poolConfig.ConnConfig.Logger = logger` -> `poolConfig.ConnConfig().SetLogger(logger)`

## Using the demos

- Assumes a PostgreSQL instance available without a password as `postgres://localhost/postgres`,
  with the `psql` command available
- Build the DB, the binaries, and install them as go executables:
  - `make && make install`
- After you are finished using them, remove the demo databases:
  - `make clean`
