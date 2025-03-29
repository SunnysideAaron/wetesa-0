# ADR-006 Logging Levels

## Status

Accepted

## Context

Which log levels to use.

## Decision

Use default slog levels, in the following manor:
  - DEBUG (-4) Only turn on for in-depth troubleshooting.
  - INFO (0) default level in production. Enough information to troubleshoot basic problems.
  - WARN (4) Create a ticket. Something is wrong and needs fixing. Properly handled errors are info not warn.
  - ERROR (8) Call someone NOW! Something is wrong and needs immediate fixing.

Allow changing log level at runtime.

Allow different log levels for different parts of the code.

## Why

Use default slog levels because they are the defaults. This project is just an example.
Most devs will expect these levels. Unless they have chosen something else on purpose.

As code base grows developers will want to be able to turn on and off logging
for different parts of the code. No need to turn on a bunch of log messages for
code they are not working on.

## Notes
- https://stackoverflow.com/questions/76970895/change-log-level-of-go-lang-slog-in-runtime

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

