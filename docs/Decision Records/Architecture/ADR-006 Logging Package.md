# ADR-006 Logging Package

## Status

Accepted

## Context

Which logging package to use.

## Decision

- use slog
- use LogAttrs()

## Why

slog because it is in the standard library. zerolog would add a dependency. 
Which is against the goal of this project. slog will be fast enough for a very
long time.
- [Which log library should I use in Go?](https://www.bytesizego.com/blog/which-log-library-go)
  - slog (built in) or zerolog (fastest but a dependency)
- [go-logging-benchmarks ](https://github.com/betterstack-community/go-logging-benchmarks)

use LogAttrs() to prevent miss matched key value pairs.

## Notes

- [Logging in Go with Slog: The Ultimate Guide](https://betterstack.com/community/guides/logging/logging-in-go/)
- [A Guide to Writing slog Handlers](https://github.com/golang/example/blob/master/slog-handler-guide/README.md)
- [go.dev blog](https://go.dev/blog/slog)
- https://pkg.go.dev/log/slog@master#example-Handler-LevelHandler
- [Go Wiki: Resources for slog](https://go.dev/wiki/Resources-for-slog)
- https://github.com/youngjun827/api-std-lib
  - logging with slog package example

## Consequences



## Other Possible Options



## Not an Option
- [Awesome Go's List of Logging](https://github.com/avelino/awesome-go?tab=readme-ov-file#logging)
  - Would add dependencies which wouldn't meet this project.
- [dozzle](https://github.com/amir20/dozzle)
  - monitors docker logs, maybe for next project
- https://www.reddit.com/r/golang/comments/1iw07rm/what_is_your_logging_monitoring_observability/
- https://www.reddit.com/r/golang/comments/1jd4ibv/adding_logging_to_a_library/
