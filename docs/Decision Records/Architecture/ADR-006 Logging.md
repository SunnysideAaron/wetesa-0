# ADR-006 Logging

## Status

Accepted, Proposed, Deprecated or Superseded (list DR)

## Context

How exactly to handle logging.

## Decision

use slog

use LogAttrs()

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

Use default slog levels, for these reasons:
  - DEBUG (-4) Only turn on for in depth troubleshooting.
  - INFO (0) default level in production. Enough information to troubleshoot basic problems.
  - WARN (4) Create a ticket. Something is wrong and needs fixing. Properly handled errors are info not warn.
  - ERROR (8) Call someone NOW! Something is wrong and needs immediate fixing.

**Pending** what to log
logs should contain enough information in order to troubleshoot a problem when reported. For an API that is at least the request, response, and error stack trace.

**Pending** What to do if user accidentally puts password in user name field?

**flesh out**
Needs to be turned on and off with out code changes (in prod) and limited to file /  function / client

TODO error stack traces

TODO hiding sensitive information

TODO what to log for api

## Why

slog because it is in the standard library. zerolog would add a dependency. 
Which is against the goal of this project. slog will be fast enough for a very
long time.
- [Which log library should I use in Go?](https://www.bytesizego.com/blog/which-log-library-go)
  - slog (built in) or zerolog (fastest but a dependency)
- [go-logging-benchmarks ](https://github.com/betterstack-community/go-logging-benchmarks)

use LogAttrs() to prevent miss matched key value pairs.

We will log to STDOUT not to a file. A different tool will send the logs to where
ever they need to go. Docker has means to see the last logs if needed in case of
needing to debug server crashes.

Is slog asynchronous? Not a problem since we are logging to STDOUT. If we were
logging to a file we would need to find out and code a solution if not.

Why not log sensitive information? Because we should assume they will fall into
the wrong hands. Going for the logs is basic steps for any hacker. At the very
least we may be sending logs to a third party where they
are stored. Even if temporarily. With right to be deleted laws any personal info
logged also has to be able to be deleted. Easier to just not log it.

Use default slog levels because they are the defaults. This is just an example.
Most devs will expect these levels. Unless they have chosen something else on purpose.

## Notes

- [Logging in Go with Slog: The Ultimate Guide](https://betterstack.com/community/guides/logging/logging-in-go/)
- [A Guide to Writing slog Handlers](https://github.com/golang/example/blob/master/slog-handler-guide/README.md)
- [go.dev blog](https://go.dev/blog/slog)
- https://stackoverflow.com/questions/76970895/change-log-level-of-go-lang-slog-in-runtime
- https://pkg.go.dev/log/slog@master#example-Handler-LevelHandler
- [Go Wiki: Resources for slog](https://go.dev/wiki/Resources-for-slog)

### which levels to use and their meanings

- log/slog package provides four log levels by default, with each one associated with an integer value:
  - DEBUG (-4)
  - INFO (0)
  - WARN (4)
  - ERROR (8)
- [Let’s talk about logging](https://dave.cheney.net/2015/11/05/lets-talk-about-logging)
  - DEBUG
    - Things that developers care about when they are developing or debugging software.
  - INFO
    - Things that users care about when using your software.
- [Google Cloud Logging API v2](https://cloud.google.com/logging/docs/reference/v2/rest/v2/LogEntry)
  - DEFAULT 	(0) The log entry has no assigned severity level.
  - DEBUG 	(100) Debug or trace information.
  - INFO 	(200) Routine information, such as ongoing status or performance.
  - NOTICE 	(300) Normal but significant events, such as start up, shut down, or a configuration change.
  - WARNING 	(400) Warning events might cause problems.
  - ERROR 	(500) Error events are likely to cause problems.
  - CRITICAL 	(600) Critical events cause more severe problems or outages.
  - ALERT 	(700) A person must take an action immediately.
  - EMERGENCY 	(800) One or more systems are unusable.
- [when to use log levels](https://www.reddit.com/r/golang/comments/1ctaz7n/when_to_use_slog_levels/)
  - Revolutionary_Ad7262
    - DEBUG: when I disable these logs I am fine with potential debugging. So there should not be any new information, which is impossible to extract from other log entries + from the code examination
    - INFO: everything, which is neccessary to examine issue on production. For example you cannot debug why your JSON request is rejected, if you don't log it
    - WARN something is not working properly, but it does not affect the business. Examples: cache operation failed (so you have lower performance, but it works anyway), HTTP request failed but you have a retries (so you log it as WARN, but if the final try fails, then ERROR, if it is necessary)
    - ERROR: something is not working and it affects the business
    - How often you should check log levels:
      - INFO/ DEBUG: never, only if needed
      - WARN: once a while, if other metrics are not alerting
      - ERROR: asap

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
