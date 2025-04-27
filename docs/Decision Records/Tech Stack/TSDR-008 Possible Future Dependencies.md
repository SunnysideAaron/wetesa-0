# TSDR-009 Possible Future Dependencies

## Status

Accepted

## Context

A primary goal of this example is to include a minimum of dependencies. Demonstrating
how to use the standard library. Projects derived from this example may
want to add some dependencies. Here are a few ideas.

## Notes

### Authentication, Authorization, or Cryptography.

Do not attempt Authentication, Authorization, or Cryptography. They are all easy
to get wrong. With catastrophic results. They should not be "home-brewed" in a
production setting. Find libraries that do those things right or at least have a
good reputation and lots of other eyes on the code.

### Code Generation

For larger projects this could save a lot of time.

### Database

There are a lot of ways to skin the database cat. Some options worth mentioning:
- Jet
  - Calls it's self a SQL builder. Not an ORM.
- [squirrel](https://github.com/Masterminds/squirrel)
- dat
  - Query builder. I'm not keen on query builders. I'd rather just write sql.
- [sqlx](https://github.com/jmoiron/sqlx)
  - This extends database/sql.
  - Rough thought is that this isn't necessary if using pgx.
  - Not all that compatible with pgx interface. Could be used with pgx's database/sql interface.

### Middleware

See "ADR-015 Middleware.md"

### pgx related packages

- scanny
  - shortens pgx calls. 
  - [Working with PostgreSQL in Go using pgx](https://donchev.is/post/working-with-postgresql-in-go-using-pgx/) "Doing SQL in Go got a lot of hate in the past because of interface{} and manually scanning the result into a struct. But with pgx v5, this is no longer the case. I think that libraries like sqlx and scany are great but not necessary anymore."
- search for other pgx related packages.

### Slog pretty handlers

- check awesome go's list
- https://pkg.go.dev/log/slog@go1.24.1#hdr-Writing_a_handler
- https://pkg.go.dev/testing/slogtest@go1.24.1
