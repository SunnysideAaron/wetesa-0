# ADR-006 Logging

## Status

Accepted, Proposed, Deprecated or Superseded (list DR)

## Context

How exactly to handle logging.

### How long to keep logs. EVALUATE

Podcast was recommending as little as 4 days (long weekend). 7 days, 14 days or 30 days. Don't need forever. Start small first. Only lengthen for business need.

----------

Needs to be turned on and off with out code changes (in prod) and limited to file /  function / client

Advice was to send all logs to standard out (standard error?) and then use a different tool to send standard out on to where ever to use logs. Research why that advice was given


https://www.reddit.com/r/golang/comments/1iw07rm/what_is_your_logging_monitoring_observability/

https://www.reddit.com/r/golang/comments/1jd4ibv/adding_logging_to_a_library/

## Decision



## Why / Notes



## Consequences



## Other Options

Possibilities:
- https://caddyserver.com/docs/logging
  - example of structuring logs

Not an option:

