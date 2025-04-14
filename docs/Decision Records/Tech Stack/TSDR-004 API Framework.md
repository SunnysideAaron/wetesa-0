# TSDR-004 API framework (or not!)

## Status

Accepted

## Context

frameworks are a way to quickly package a set of solutions to common problems.
Allowing development to get going quickly.

## Decision

Use the standard library.

## Why / Notes

It is repeatedly recommended to not use a framework. Using the standard library
instead. Most frameworks in Go were developed before Go 1.22 (2024-FEB) added
better routing. Using them will lock code into their middleware and way of doing
things.

## Consequences

Having to build things from scratch takes longer.

## Other Possible Options

- Example tech stack from [reddit post](https://www.reddit.com/r/golang/comments/15y5wiq/lets_say_you_want_to_build_a_go_rest_api_should/) Chi, connectrpc, sqlc, squirrel, 3rd party auth
Chi for routing and mixing, for JSON req-res, I'll use connectrpc. In case I need cookie auth or file upload I use Chi I'll go with connectrpc.com for the transport and application layer. You get the protobuf as API spec and you can also generate the client SDK. It works like twirp but complies with grpc. For db access I'll use SQLC, query is validated and faster than ORM layer. You can use query builder like Squirel for complex dynamic query Use 3rd party auth, so you don't spent too much time working on authentication

- [Awesome Go's List of Web Frameworks](https://github.com/avelino/awesome-go?tab=readme-ov-file#web-frameworks)

### Web Frameworks Listing

- [Atreugo](https://github.com/savsgio/atreugo) - High performance and extensible micro web framework with zero memory allocations in hot paths.
- [Beego](https://github.com/beego/beego) - beego is an open-source, high-performance web framework for the Go programming language.
- [Chi]()
  - **SHORT LIST TO TRY** should complement standard library
  - "Chi's middleware, httplog and render packages are fantastic."
- [Chimera](https://github.com/matt1484/chimera)
  - Chi based rest api framework
- [Confetti Framework](https://confetti-framework.github.io/docs/) - Confetti is a Go web application framework with an expressive, elegant syntax. Confetti combines the elegance of Laravel and the simplicity of Go.
- [Don](https://github.com/abemedia/go-don) - A highly performant and simple to use API framework.
- [Echo](https://github.com/labstack/echo) - High performance, minimalist Go web framework.
  - **SHORTISH LIST TO TRY**
- [Fastschema](https://github.com/fastschema/fastschema) - A flexible Go web framework and Headless CMS.
- [Fiber](https://github.com/gofiber/fiber) - An Express.js inspired web framework build on Fasthttp.
  - **SHORTISH LIST TO TRY**
- [Flamingo](https://github.com/i-love-flamingo/flamingo) - Framework for pluggable web projects. Including a concept for modules and offering features for DI, Configareas, i18n, template engines, graphql, observability, security, events, routing & reverse routing etc.
- [Flamingo Commerce](https://github.com/i-love-flamingo/flamingo-commerce) - Providing e-commerce features using clean architecture like DDD and ports and adapters, that you can use to build flexible e-commerce applications.
- [Fuego](https://github.com/go-fuego/fuego) - The framework for busy Go developers! Web framework generating OpenAPI 3 spec from source code.
- [Gin](https://github.com/gin-gonic/gin) - Gin is a web framework written in Go! It features a martini-like API with much better performance, up to 40 times faster. If you need performance and good productivity.
  - **SHORT LIST TO TRY**
- [Ginrpc](https://github.com/xxjwxc/ginrpc) - Gin parameter automatic binding tool,gin rpc tools.
- [GoFr](https://github.com/gofr-dev/gofr) - Gofr is an opinionated microservice development framework.
- [GoFrame](https://github.com/gogf/gf) - GoFrame is a modular, powerful, high-performance and enterprise-class application development framework of Golang.
- [golamb](https://github.com/twharmon/golamb) - Golamb makes it easier to write API endpoints for use with AWS Lambda and API Gateway.
- [Gone](https://github.com/gone-io/gone) - A lightweight dependency injection and web framework inspired by Spring.
- [goravel](https://github.com/goravel/goravel) - A Laravel-inspired web framework with ORM, authentication, queue, task scheduling, and more built-in features.
- [gorilla/mux]()
  - was abandoned and then restarted. old
- [Goyave](https://github.com/go-goyave/goyave) - Feature-complete REST API framework aimed at clean code and fast development, with powerful built-in functionalities.
- [Hertz](https://github.com/cloudwego/hertz) - A high-performance and strong-extensibility Go HTTP framework that helps developers build microservices.
- [hiboot](https://github.com/hidevopsio/hiboot) - hiboot is a high performance web application framework with auto configuration and dependency injection support.
- [iWF](https://github.com/indeedeng/iwf) - iWF is an all-in-one platform for developing long-running business processes. It offers a convenient abstraction for utilizing databases, ElasticSearch, message queues, durable timers, and more, with a clean, simple, and user-friendly interface.
- [Lit](https://github.com/jvcoutinho/lit) - Highly performant declarative web framework for Golang, aiming for simplicity and quality of life.
- [Microservice](https://github.com/claygod/microservice) - The framework for the creation of microservices, written in Golang.
- [patron](https://github.com/beatlabs/patron) - Patron is a microservice framework following best cloud practices with a focus on productivity.
- [Pnutmux](https://gitlab.com/fruitygo/pnutmux) - Pnutmux is a powerful Go web framework that uses regex for matching and handling HTTP requests. It offers features such as CORS handling, structured logging, URL parameters extraction, middlewares, and concurrency limiting.
- [Revel](https://github.com/revel/revel) - High-productivity web framework for the Go language.
- [rk-boot](https://github.com/rookie-ninja/rk-boot) - A bootstrapper library for building enterprise go microservice with Gin and gRPC quickly and easily.
- [rux](https://github.com/gookit/rux) - Simple and fast web framework for build golang HTTP applications.
- [templui](https://github.com/axzilla/templui) - Modern UI Components for Go & Templ.
- [uAdmin](https://github.com/uadmin/uadmin) - Fully featured web framework for Golang, inspired by Django.
- [WebGo](https://github.com/naughtygopher/webgo) - A micro-framework to build web apps with handler chaining, middleware, and context injection. With standard library-compliant HTTP handlers (i.e., `http.HandlerFunc`)..
- [xun](https://github.com/yaitoo/xun)
  - uses go 1.22 mux
- [Yokai](https://github.com/ankorstore/yokai) - Simple, modular, and observable Go framework for backend applications.
  - (built on echo)
  - [demo](https://ankorstore.github.io/yokai/demos/http-application/)

## Not an Option

- [Goa](https://github.com/goadesign/goa) - Goa provides a holistic approach for developing remote APIs and microservices in Go.
  - Looked into goa. new user documentation / steps kept being painful. Everything just slightly not working and requiring troubleshooting. If new user experience is like that than can't trust code. May be willing to try again.
- [Huma](https://github.com/danielgtaylor/huma/) - Framework for modern REST/GraphQL APIs with built-in OpenAPI 3, generated documentation, and a CLI.
  - Same issue as goa. Tutorial was working well until it didn't, couldn't troubleshoot and fix without time / pain.
- [ogen]()
  - Similar problem as goa. Intro docs are light and not complete. Had limited patience for documentation that wasn't working.
- [swag](https://github.com/swaggo/swag)
  - Seems to just be special comments in code. Ie writing the api spec inline with the code. Not bad but doesn't save typing. That's just typing in the same place.
- [go-rest-api-service-template ](https://github.com/p2p-b2b/go-rest-api-service-template)
  - includes dependencies
  - readme has a bunch of good good links worth following.
- [ardanlabs](https://github.com/ardanlabs/service)
  - Seems a bit overly complex to use starting point for a basic example, but may be a good study example
