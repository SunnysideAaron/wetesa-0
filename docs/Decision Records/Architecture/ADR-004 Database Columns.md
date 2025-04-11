# ADR-004 Database Columns

## Status

Accepted

## Context

There are some basic rules for database columns that if you don't choose correctly
from the beginning will F*** your team / project hard later.

## Decision

- table_id not id
- UUIDs for primary keys
- Date / Time Columns: ACTION_WORD_date, ACTION_WORD_datetime, ACTION_WORD_time (times without date)
- Time columns will be of type timestamptz and UTC. User timezone in separate column if needed.

## Why / Notes

### table_id not id

tables id columns are named are "table_id" not "id" so that when doing joins the
developer knows the id's match and links are not on mis matched ids.

When joining multiple tables in complex queries it is far less confusing to have
joining column names match. TODO provide examples.

example should show aliased tables. and why on complex joins having column names
match maters.

### UUIDs for all id type primary keys

Simple incrementing integers make replication (scaling) difficult. Use of UUIDs
makes replication far easier.

https://www.postgresql.org/docs/current/datatype-uuid.html
https://ntietz.com/blog/til-uses-for-the-different-uuid-versions/
    - "For example, consider using v7 if you are using UUIDs as database keys."
    - Would we ever want v7 of uuid? does postgress care? seems v4 is default in postgress. This needs deeper research.
https://neon.tech/postgresql/postgresql-tutorial/postgresql-uuid
https://www.reddit.com/r/golang/comments/1jdakzs/recommended_way_to_use_uuid_typesto_type_or_not/
https://github.com/avelino/awesome-go?tab=readme-ov-file#uuid

### Date / Time Columns

ACTION_WORD_date, ACTION_WORD_datetime, ACTION_WORD_time (times without date.)

Examples:
created_date, created_time, created_datetime, effective_start_date, effective_end_datetime

I've seen plenty of examples including _at, _on, _date, _datestamp. Any of which
in practice ends up having any kind of date or time. Lets just say what we mean.
There is certainly an argument against including column types in column names but
in this case we make the exception.

- don't use timetz for ACTION_WORD_time columns see:
  [Postreges Wiki: Don't Do This](https://wiki.postgresql.org/wiki/Don't_Do_This)

### timestamptz and UTC

- [Postreges Wiki: Don't Do This](https://wiki.postgresql.org/wiki/Don't_Do_This)
- https://community.spiceworks.com/t/zone-of-misunderstanding/928839
"you just SET TIMEZONE in the user’s connection to the database, and timestamps
 will automatically come back in the appropriate time zone. Beats the heck out of
  messes like the PHPBB time zone code."

## Consequences



## Other Possible Options



## Not an Option
- How about xyz_at for a timestamp and xyz_on for a date field - eg start_at or start_on?
  - Not too bad but if going that far then just use _date and _time.







