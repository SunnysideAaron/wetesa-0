# TSDR-009 Non-dependencies

## Status

Accepted

## Context

For different reasons this example project will not even attempt certain things.
A primary goal of this example is to include a minimum of dependencies.

## Decision

Do not attempt Authentication, Authorization, or Cryptography.

Don't use any of the database packages that are not in the standard library.

## Why / Notes

Authentication, Authorization, or Cryptography are all easy to get wrong. With 
catastrophic results. They should not be "home-brewed" in a production setting.
Find libraries that do those things right.

There are several packages that would help with database work. But they all add
dependencies.

## Consequences



## Other Possible Options
- Jet
  - Calls it's self a SQL builder. Not an ORM.
- ?squril?
- scanny
  - shortens pgx calls. 
  - [Working with PostgreSQL in Go using pgx](https://donchev.is/post/working-with-postgresql-in-go-using-pgx/) "Doing SQL in Go got a lot of hate in the past because of interface{} and manually scanning the result into a struct. But with pgx v5, this is no longer the case. I think that libraries like sqlx and scany are great but not necessary anymore."
- other pgx related packages.

## Not an Option
- dat
  - Query builder. I'm not keen on query builders. I'd rather just write sql.
- [sqlx](https://github.com/jmoiron/sqlx)
  - This extends database/sql.
  - Rough thought is that this isn't necessary if using pgx.
  - Not all that compatible with pgx interface. Could be used with pgx's database/sql interface. 
