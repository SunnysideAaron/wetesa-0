# ADR-004 Database Columns

## Status

Accepted, Proposed, 

## Context

### table_id not id

When joining multiple tables in complex queries it is far less confusing to have joining column names match. TODO provide examples.

tables ids are "table_id" not "id" so that when doing joins the developer knows the id's match and links are not on mis matched ids.

example should show aliased tables. and why on complex joins having column names match maters.


### UUIDs for all id type primary keys

Simple incrementing integers make replication difficult. Use of UUIDs makes replication far easier.

TODO RESEARCH:
https://www.postgresql.org/docs/current/datatype-uuid.html
https://ntietz.com/blog/til-uses-for-the-different-uuid-versions/
    - "For example, consider using v7 if you are using UUIDs as database keys."
    - Would we ever want v7 of uuid? does postgress care? seems v4 is default in postgress. This needs deeper research.
https://neon.tech/postgresql/postgresql-tutorial/postgresql-uuid
https://www.reddit.com/r/golang/comments/1jdakzs/recommended_way_to_use_uuid_typesto_type_or_not/
https://github.com/avelino/awesome-go?tab=readme-ov-file#uuid

### Date Columns

TODO write up date columns

utc in all date columns, how to store timezone used at time?

nameing convention of date columns
created_date, created_time, created, etc? time stamp
if time column time with and without date


TODO
use column type timestamptz

https://community.spiceworks.com/t/zone-of-misunderstanding/928839
"you just SET TIMEZONE in the user’s connection to the database, and timestamps will automatically come back in the appropriate time zone. Beats the heck out of messes like the PHPBB time zone code."

### All times in UTC
    Store timezone of user entering as well.

## Decision

- table_id not id
- Use UUIDs for primary keys
- PENDING DATE Column names

## Consequences










