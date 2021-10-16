# PGXer

## Purpose

_pgxer_ is a thin wrapper around [jackc/pgx](https://github.com/jackc/pgx) that
exposes the _pgx_ 4.x API, but using interface types instead of concrete types, 
to make it easier to mock DB access at any level from the `Pool` downwards.

In v4, _pgx_ provides a mix of interfaces (e.g. `Rows`) and concrete types
(e.g. `Pool`), making it essentially impossible to write tests relying on a DB
without accessing the DB, causing tests to be slow and need external setup.

_pgxer_ uses identical public type names, but modifies the signatures of all public
methods to take and return interfaces to support mocking all types; and includes
a mock implementation for all of them.

Where public fields exists, these are replaced by equivalent methods to support
mocking. The mocks are based on 

## Use

For the simpler cases, just modifying the import paths from should be sufficient:

- `import "github.com/jackc/pgx/v4"` -> `import pgx "github.com/fgm/pgxer"`
- `import "github.com/jackc/pgx/v4/pgxpool"` -> `import "github.com/fgm/pgxer/pgxpool"`
