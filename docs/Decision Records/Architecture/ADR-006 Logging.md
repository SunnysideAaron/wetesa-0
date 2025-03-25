# ADR-006 Logging

## Status

Accepted, Proposed, Deprecated or Superseded (list DR)

## Context

How exactly to handle logging.

## Decision

slog

log to STDOUT

Do not log sensitive information, including:
- Passwords
- email addresses
- Phone numbers
- DOB
- Age
- Addresses
- Bank Accounts
- Credit card numbers
- Social security numbers
- Personal identification numbers (PINs)
- Health information
- Financial information
- Any other sensitive data. (add to this list)

**Pending** what log levels to use

**Pending** what to log

**Pending** What to do if user accidentally puts password in user name field?

**flesh out**
Needs to be turned on and off with out code changes (in prod) and limited to file /  function / client



## Why

slog because it is in the standard library. zerolog would add a dependency. 
Which is against the goal of this project. slog will be fast enough for a very
long time.
- [Which log library should I use in Go?](https://www.bytesizego.com/blog/which-log-library-go)
  - slog (built in) or zerolog (fastest but a dependency)
- [go-logging-benchmarks ](https://github.com/betterstack-community/go-logging-benchmarks)

We will log to STDOUT not to a file. A different tool will send the logs to where
ever they need to go. Docker has means to see the last logs if needed in case of
needing to debug server crashes.

Is slog asynchronous? Not a problem since we are logging to STDOUT. If we were
logging to a file we would need to find out and code a solution if not.

Why not log sensitive information? Because we should assume they will fall into
the wrong hands. Going for the logs is basic steps for any hacker. At the very
least we will be sending logs to a third party where they
are stored. Even if temporarily.

## Notes

- [Logging in Go with Slog: The Ultimate Guide](https://betterstack.com/community/guides/logging/logging-in-go/)
- [A Guide to Writing slog Handlers](https://github.com/golang/example/blob/master/slog-handler-guide/README.md)
- [go.dev blog](https://go.dev/blog/slog)
- https://stackoverflow.com/questions/76970895/change-log-level-of-go-lang-slog-in-runtime
- https://pkg.go.dev/log/slog@master#example-Handler-LevelHandler
- [Go Wiki: Resources for slog](https://go.dev/wiki/Resources-for-slog)

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
