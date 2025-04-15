# ADR-006 Logging Output

## Status

Accepted

## Context

Where to send logs

## Decision

log to STDOUT

## Why

We will log to STDOUT not to a file. A different tool will send the logs to where
ever they need to go. Docker has means to see the last logs if needed in case of
needing to debug server crashes.

Is slog asynchronous? Not a problem since we are logging to STDOUT. If we were
logging to a file we would need to find out and code a solution if not.

## Notes

## Consequences
