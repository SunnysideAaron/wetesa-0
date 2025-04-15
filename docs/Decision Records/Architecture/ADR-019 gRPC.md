# ADR-016 gRPC

## Status

Accepted

## Context



## Decision

Not using gRPC at this time.

## Why / Notes

gRPC is probably overkill for a simple API. gRPC is for server to server
 communication. ie microservices.

"My advice: don't use gRPC if you need to make calls from your browser. It should
 only be used for server-to-server communication in my opinion. The browser
  support is not there for the trailers or use of HTTP/2. "
- https://www.reddit.com/r/golang/comments/1c1hwbf/is_grpc_a_good_alternative_for_rest_when_building/

## Consequences

May require some serious refactoring if we decide to use gRPC later.

## Other Options

Possibilities:

Not an option:

