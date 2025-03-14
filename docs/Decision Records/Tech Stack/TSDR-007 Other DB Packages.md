# TSDR-007 Other DB packages 

## Status

Accepted, Proposed, Deprecated or Superseded (list DR)

## Context



## Decision



## Why / Notes



## Consequences



## Other Options

Possibilities:
- Jet
  - Calls it's self a SQL builder. Not an ORM.
- ?squril?
- scanny
  - shortens pgx calls. 
  - [Working with PostgreSQL in Go using pgx](https://donchev.is/post/working-with-postgresql-in-go-using-pgx/) "Doing SQL in Go got a lot of hate in the past because of interface{} and manually scanning the result into a struct. But with pgx v5, this is no longer the case. I think that libraries like sqlx and scany are great but not necessary anymore."
- other pgx related packages.

Not an option:
- dat
  - Query builder. I'm not keen on query builders. I'd rather just write sql.
- [sqlx](https://github.com/jmoiron/sqlx)
  - This extends database/sql.
  - Rough thought is that this isn't necessary if using pgx.
  - Not all that compatible with pgx interface. Could be used with pgx's database/sql interface. 
