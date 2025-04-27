# ADR-015 Middleware

## Status

Accepted

## Context



## Decision

Installed middleware:
- AllowQuerySemicolons (Standard Library)
- CORS
- IP
- Logging
- MaxBytesHandler (Standard Library)
- Recover
- RequestID
- TimeoutHandler (Standard Library)

### Pending

- [Alice](https://github.com/justinas/alice)
  - Middleware chaining
- [nosurf](https://github.com/justinas/nosurf)
  - prevents Cross-Site Request Forgery attacks. Is this related to CORS?
- Rate limiting seems a smidge complex with saving IPs. I'm punting that ball for now.
  - [api-std-lib](https://github.com/youngjun827/api-std-lib/blob/main/cmd/api/middleware.go)
    - has a simple rate limiter.

## Sources of Middlewares

- [chi](https://github.com/go-chi/chi?tab=readme-ov-file#middlewares)
- [awesome-go](https://github.com/avelino/awesome-go?tab=readme-ov-file#middlewares)
- [grafana](https://github.com/grafana/grafana/tree/main/pkg/middleware)
- [exposure-notifications-server](https://github.com/google/exposure-notifications-server/tree/main/internal/middleware)
- [echo](https://echo.labstack.com/docs/category/middleware)
  - will they work with standard library? example code if nothing else
- gin
  - [rk-gin](https://github.com/rookie-ninja/rk-gin)
  - [gin-contrib](https://github.com/gin-contrib)
- [rye](https://github.com/InVisionApp/rye)
- [For slog](https://go.dev/wiki/Resources-for-slog#logging-middleware)

## Examples for specific problems
- RequestID
  - https://github.com/google/exposure-notifications-server/blob/main/internal/middleware/request_id.go
  - https://github.com/gin-contrib/requestid/blob/master/requestid.go
  - https://github.com/go-chi/chi/blob/v1.5.5/middleware/request_id.go#L67
  - https://github.com/labstack/echo/blob/master/middleware/request_id.go
    - https://echo.labstack.com/docs/middleware/request-id
- Rate Limiter
  - [api-std-lib](https://github.com/youngjun827/api-std-lib/blob/main/cmd/api/middleware.go)
    - has a simple rate limiter.
  - [Tollboth](https://github.com/didip/tollbooth)
    - strong contender but would require a dependency.
  - [echo](https://github.com/labstack/echo/blob/master/middleware/rate_limiter.go)
    - https://echo.labstack.com/docs/middleware/rate-limiter
    - per IP
  - https://github.com/ulule/limiter
    - per IP
  - https://github.com/rookie-ninja/rk-gin/blob/master/middleware/ratelimit/middleware.go
    - per global or path
  - [chi](https://pkg.go.dev/github.com/go-chi/chi/middleware#Throttle)
    - per global
- Recovery
  - [exposure notificaiton server](https://github.com/google/exposure-notifications-server/blob/main/internal/middleware/recovery.go)
    - there is a unit test for this.
  - [chi](https://github.com/go-chi/chi/blob/v1.5.5/middleware/recoverer.go#L21)
  - [echo](https://echo.labstack.com/docs/middleware/recover)

## Not an Option
- [AutoVerse: A Modular Go Framework for RESTful APIs](https://github.com/Muga20/Go-Modular-Application)
  - didn't like this
- [Timeout Middleware in Go: Simple in Theory, Complex in Practice ](https://www.reddit.com/r/golang/comments/1jf1inr/timeout_middleware_in_go_simple_in_theory_complex/)
  - I could be wrong but default timeout should work OK.
- https://github.com/ngamux/ngamux?tab=readme-ov-file#provided-middlewares
  - There are other examples of everything listed here. Check those first.
- [api-std-lib](https://github.com/youngjun827/api-std-lib/blob/main/cmd/api/middleware.go)
  - writen for go 1.21 but does have some interesting things like rate limiting middleware. logging with slog package and data validation
