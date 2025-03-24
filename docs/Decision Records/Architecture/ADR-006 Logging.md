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





## Decision



## Why / Notes

- [Logging in Go with Slog: The Ultimate Guide](https://betterstack.com/community/guides/logging/logging-in-go/)
- [Complete Guide to Logging in Golang with slog](https://signoz.io/guides/golang-slog/)

- [go.dev blog](https://go.dev/blog/slog)
- [Which log library should I use in Go?](https://www.bytesizego.com/blog/which-log-library-go)
  - slog, or zerolog (fastest but a dependency)

- https://stackoverflow.com/questions/76970895/change-log-level-of-go-lang-slog-in-runtime


### which levels to use and their meanings
- [Let’s talk about logging](https://dave.cheney.net/2015/11/05/lets-talk-about-logging)
- [when to use log levels](https://www.reddit.com/r/golang/comments/1ctaz7n/when_to_use_slog_levels/)
- [Google Cloud Logging API v2](https://cloud.google.com/logging/docs/reference/v2/rest/v2/LogEntry)


## Consequences


## Other Possible Options

- https://github.com/youngjun827/api-std-lib
  - logging with slog package example

## Not an Option
- [Awesome Go's List of Logging](https://github.com/avelino/awesome-go?tab=readme-ov-file#logging)
  - Would add dependencies which wouldn't meet this project.
- [dozzle](https://github.com/amir20/dozzle)
  - monitors docker logs, maybe for next project
- https://www.reddit.com/r/golang/comments/1iw07rm/what_is_your_logging_monitoring_observability/
- https://www.reddit.com/r/golang/comments/1jd4ibv/adding_logging_to_a_library/
