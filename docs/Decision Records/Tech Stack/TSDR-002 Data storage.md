# TSDR-002 Data storage  

## Status

Accepted

## Context



## Decision

PostgreSQL

## Why / Notes

  - Pros. Enterprise scale.
  - ACID
    - At one time in the past using triggers was not recommended. Perhaps that is no longer the case.
  - Other reasons 
    - (Can't remember off hand why it's better than MariaDB. Scale? ACID? I remember there were good reasons just can't remember them right now. I'm sure there are articles on topic.)

## Consequences

Data store lock in. Changing databases is something we like to talk about but in practice it is way too costly down the road.

## Other Options

Possibilities:
- MariaDB
  - Strong option.
- MySQL
  - Oracle dependency.

Not an option:
- SQLite
  - Limited features. Not appropriate for enterprise scale.