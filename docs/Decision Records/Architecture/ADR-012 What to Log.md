# ADR-006 What to Log

## Status

Accepted

## Context

What should be logged?

## Decision

Log request, response, and error stack trace
- **Pending** Error stack trace in this example is rough. Could use alternative if using dependencies.
- https://betterstack.com/community/guides/logging/logging-in-go/#error-logging-with-slog

See ADR-007 Sensitive Information

## Why

Logs should contain enough information in order to troubleshoot a problem when reported.
For an API that is at least the request, response, and error stack trace.

## Notes

## Consequences
