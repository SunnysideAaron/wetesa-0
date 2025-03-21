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
https://www.reddit.com/r/golang/comments/1jdakzs/recommended_way_to_use_uuid_typesto_type_or_not/
https://github.com/avelino/awesome-go?tab=readme-ov-file#uuid

### Date Columns

TODO write up date columns

utc in all date columns, how to store timezone used at time?

nameing convention of date columns
created_date, created_time, created, etc? time stamp
if time column time with and without date

### All times in UTC
    Store timezone of user entering as well.

## Decision

- table_id not id
- Use UUIDs for primary keys
- PENDING DATE Column names

## Consequences










