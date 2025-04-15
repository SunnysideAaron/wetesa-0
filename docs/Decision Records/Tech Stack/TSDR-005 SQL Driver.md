# TSDR-005 SQL driver

## Status

Accepted

## Context

Go does not include sql drivers in the standard library. [Driver listing](https://go.dev/wiki/SQLDrivers)

## Decision

- [pgx](https://github.com/jackc/pgx)
- Use the pgx interface. Not pgx's database/sql interface.

## Why / Notes

- Often recommended package. 
- Is more than just a sql driver. But not an all db problems in one like Bun.
- Support for approximately 70 different PostgreSQL types
- [Choosing Between the pgx and database/sql Interfaces](https://github.com/jackc/pgx#choosing-between-the-pgx-and-databasesql-interfaces)
- We are 100% committed to PostgreSQL. Use the pgx interface.
- Don't bother with pgx's database/sql interface. pgx interface is more performative.
- [PGX Top to Bottom](https://www.youtube.com/watch?v=sXMSWhcHCf8)
- [How to work with Postgres in Go](https://medium.com/avitotech/how-to-work-with-postgres-in-go-bad2dabd13e4)
- [Working with PostgreSQL in Go using pgx](https://donchev.is/post/working-with-postgresql-in-go-using-pgx/)
- [pkg.go.dev pgx](https://pkg.go.dev/github.com/jackc/pgx#section-documentation)

## Consequences



## Other Options

Possibilities:
- [Bun](https://bun.uptrace.dev/)
  - includes SQL Driver, ORM, Fixtures, Migrations, Multiple DBs, query builder, and models. Con: probably not a good fit with sqlc.
- [pq](https://github.com/lib/pq)
  - Last release was 2023-04

Not an option:
- [gopgsqldriver](https://github.com/jbarham/gopgsqldriver)
  - Last code edit was 2011
- [pg](https://github.com/go-pg/pg)
  - Replaced by Bun.
