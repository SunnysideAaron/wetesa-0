# ADR-00
# TSDR-00

## Status

Accepted, Proposed, Deprecated or Superseded (list DR)

## Context



## Decision



## Why / Notes



## Consequences



## Other Options

Example code:
- https://github.com/avelino/awesome-go?tab=readme-ov-file#middlewares
- [How I write HTTP services in Go after 13 years](https://grafana.com/blog/2024/02/09/how-i-write-http-services-in-go-after-13-years/)
  - [The adapter pattern for middleware](https://grafana.com/blog/2024/02/09/how-i-write-http-services-in-go-after-13-years/#the-adapter-pattern-for-middleware)
  - [Sometimes I return the middleware](https://grafana.com/blog/2024/02/09/how-i-write-http-services-in-go-after-13-years/#sometimes-i-return-the-middleware)
    - [An opportunity to hide the request/response types away](https://grafana.com/blog/2024/02/09/how-i-write-http-services-in-go-after-13-years/#an-opportunity-to-hide-the-requestresponse-types-away)
- [grafana](https://github.com/grafana/grafana/tree/main/pkg/middleware)
  - lots of different middlewares
- [Timeout Middleware in Go: Simple in Theory, Complex in Practice ](https://www.reddit.com/r/golang/comments/1jf1inr/timeout_middleware_in_go_simple_in_theory_complex/)
- [exposure-notifications-server](https://github.com/google/exposure-notifications-server/tree/main/internal/middleware)
- https://github.com/youngjun827/api-std-lib
  - writen for go 1.21 but does have some interesting things like rate limiting middleware. logging with slog package and data validation
- https://github.com/ngamux/ngamux?tab=readme-ov-file#provided-middlewares
  
Possibilities:

Not an option:
- [AutoVerse: A Modular Go Framework for RESTful APIs](https://github.com/Muga20/Go-Modular-Application)
  - didn't like this

