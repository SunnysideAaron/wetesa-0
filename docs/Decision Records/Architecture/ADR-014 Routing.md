# ADR-00
# TSDR-00

## Status

Accepted, Proposed, Deprecated or Superseded (list DR)

## Context

Needed:
- middleware
- groupging routes
- middleware grouping
- $ wildcard feature
- route default to home
- route not found error as well?

## Decision

Use the standard library for routing.

There appears to be 3 ways to pass information between middleware / handlers:
- setting middleware as a method of a type (possibly a struct)
- context
  - good for globlal key value pairs. ie user id, session id, request id, etc.
- method parameter
  - breaks the (ResponseWriter, *Request) method call, but there are ways around that.
  - Going to avoid this pattern for now.

Will use the "http.MethodGet" etc constants. [pkg.go.dev](https://pkg.go.dev/net/http#pkg-constants)
This will allow the compiler to catch mistypes.

## Why / Notes
Since Go 1.22 (2024-FEB) added ServeMux, better routing, and path parameters the
standard library is recommended by the community for routing.
This is further supported by the fact that many routers are now abandoned.

### information

- [The standard library now has all you need for advanced routing in Go.](https://www.youtube.com/watch?v=H7tbjKFSg58&t=8s) (2024-MAR)
  - [summary](https://gist.ly/youtube-summarizer/advanced-http-routing-with-go-122-path-parameters-http-methods-middleware-and-more)
  - example loging middleware, wrapping middleware, v1 api, of middleware to
    specific routes (uses sub routers), context to pass values
- [Confused by http.HandlerFunc? This post makes it click](https://www.willem.dev/articles/http-handler-func/) (2023-APR)
- https://go.dev/blog/routing-enhancements
- https://douglasmakey.medium.com/go-1-22s-http-package-updates-42aca70ceb9b
  - shows how to do {$} wildcard feature 
- https://www.reddit.com/r/golang/comments/1aoxlsr/middleware_in_go_1220/
  - drannoc-dono has example of passing types to middleware
- https://drstearns.github.io/tutorials/gomiddleware/
  - https://drstearns.github.io/tutorials/gohandlerctx/
  -  has part about using methods on types (structs) to pass values to middleware
- https://vishnubharathi.codes/blog/exploring-middlewares-in-go/
  - Section "Enter http.Handler" has part about using methods on types (structs) to pass values to middleware

### simple examples

- [How I write HTTP services in Go after 13 years](https://grafana.com/blog/2024/02/09/how-i-write-http-services-in-go-after-13-years/)
  - [Maker funcs return the handler](https://grafana.com/blog/2024/02/09/how-i-write-http-services-in-go-after-13-years/#maker-funcs-return-the-handler)
- [Building REST APIs With Go 1.22 http.ServeMux](https://shijuvar.medium.com/building-rest-apis-with-go-1-22-http-servemux-2115f242f02b) (2024-FEB)
  - also has using methods on types (structs) to pass values to handlers
- CHI features in standard library
    - [Middleware and grouping with stdlib](https://gist.github.com/alexaandru/747f9d7bdfb1fa35140b359bf23fa820)
    - [reddit post on why still chi](https://www.reddit.com/r/golang/comments/1avn6ih/is_chi_relevant_anymore/)
- https://codewithflash.com/advanced-routing-with-go-122
- [RapidGo](https://github.com/rwiteshbera/rapidgo)

### complex examples

- [ardanlabs](https://github.com/ardanlabs/service/blob/master/app/domain/homeapp/route.go)

## Consequences

Don't have support of other developers on framework. Don't get any of the other goodies the framework has baked in.

## Other Possible Options
- chi
- https://github.com/ngamux/ngamux


## Not an Option

- [Awesome Go's List of Routers](https://github.com/avelino/awesome-go?tab=readme-ov-file#routers)
  - Almost all of these are pre Go 1.22. Most are abandoned.
- [AutoVerse: A Modular Go Framework for RESTful APIs](https://github.com/Muga20/Go-Modular-Application)
  - didn't like this
- [Let’s say you want to build a Go REST API. Should you use the standard library, a router, or a full-blown framework?](https://www.reddit.com/r/golang/comments/15y5wiq/lets_say_you_want_to_build_a_go_rest_api_should/)
  - (2023-?MAR?) Pre Go 1.22
- [DIY Golang Web Server: No Dependencies Needed!](https://www.youtube.com/watch?v=eqvDSkuBihs) (Video)
  - Too simple
- [Rahjoo: lightweight zero dependency http router library](https://www.reddit.com/r/golang/comments/1jdmzw5/lightweight_zero_dependency_http_router_library/)
  - not quite standard library
- [Golang REST API Example [Without Framework]](https://golang.cafe/blog/golang-rest-api-example.html)
  - old pre go 1.22 example. does show using constants for http methods