# TSDR-004 API framework (or not!)

## Status

Pending

## Context

API with open auth login, return 1 table
	You might want something like SCS for sessions, maybe Goth for Google/Apple authentication/authorization if you need it.
	
		
	research gRPC vrs rest API
		https://www.reddit.com/r/golang/comments/1c1hwbf/is_grpc_a_good_alternative_for_rest_when_building/
		"My advice: don't use gRPC if you need to make calls from your browser. It should only be used for server-to-server communication in my opinion. The browser support is not there for the trailers or use of HTTP/2. "
	
	
Alex Edwards’ book “Let’s Go” which features a CRUD app

Primeagen’s 2+ hour video on building an HTMX app.


We want auto code generation for (should include basic validation):
- API End points
- API to DB crud
- DB Schema
- OpenAPI 3 (Swagger 1 and 2 are previous versions of OpenAPI)
- Web Client use of API
- Mobile Client use of API

[Awesome Go's List of Web Frameworks](https://github.com/avelino/awesome-go?tab=readme-ov-file#web-frameworks)

Known options (started from post at [reddit](https://www.reddit.com/r/golang/comments/1avsog1/go_openapi_codegen/)):
- Code → Spec
  - swaggo/swag (OAS3 Beta) **TRIED**
    supported frameworks
    - gin **SHORT LIST TO TRY**
    - echo
    - buffalo
    - net/http (standard library) **SHORT LIST TO TRY**
    - gorilla/mux
      - was abandoned and then restarted. old
    - go-chi/chi  **SHORT LIST TO TRY** should complement standard library
    - flamingo
    - fiber
    - atreugo
    - hertz
  - Huma by
  - Fuego (built with go 1.22 uses generics and standard http started end of 2023)
  - Tonic
  - Astra by
  - (Gin-only, Echo & Fiber WIP as of 2024-02-22)
- Spec → SDK
  - https://packagemain.tech/p/practical-openapi-in-golang **SHORT LIST TO TRY**
  - oapi-codegen
    doesn't do OpenAPI 3?
  - ogen **TRIED**
  - openapi-generator
  - swagger-codegen
  - microsoft/kiota
- DSL → Spec + Code
  - goa  **TRIED**
- Unknown if Spec or SDK first
  - [swaggest - rest](https://github.com/swaggest/rest)
  - [swaggest - openapi-go](https://github.com/swaggest/openapi-go)
- Just Code ie (Web Framework)
  -  Yokai (built on echo)
      - [demo](https://ankorstore.github.io/yokai/demos/http-application/)
- OpenAPI Implementation (these appear to be more about validation testing spec? look into later)
  - libopenapi
  - kin-openapi (seems semi-abandoned as of 2024-02-22)

- go restful
  - generates open api spec from code.

### Other example tech stacks

From [reddit post](https://www.reddit.com/r/golang/comments/15y5wiq/lets_say_you_want_to_build_a_go_rest_api_should/) Chi, connectrpc, sqlc, squirrel, 3rd party auth
Chi for routing and mixing, for JSON req-res, I'll use connectrpc. In case I need cookie auth or file upload I use Chi I'll go with connectrpc.com for the transport and application layer. You get the protobuf as API spec and you can also generate the client SDK. It works like twirp but complies with grpc. For db access I'll use SQLC, query is validated and faster than ORM layer. You can use query builder like Squirel for complex dynamic query Use 3rd party auth, so you don't spent too much time working on authentication

### goa

Looked into goa. new user documentation / steps kept being painful. Everything just slightly not working and requiring troubleshooting. If new user experience is like that than can't trust code.

### Huma

Same issue as goa. Tutorial was working well until it didn't, couldn't troubleshoot and fix without time / pain.

### net/http (Go standard library)

- [oto](https://github.com/pacedotdev/oto/tree/main/otohttp) by Mat Ryer **liked this** can generate anything I want from templates.
- [how-i-write-http-services-in-go-after-13-years](https://grafana.com/blog/2024/02/09/how-i-write-http-services-in-go-after-13-years/) by Mat Ryer
- [The standard library now has all you need for advanced routing in Go.](https://www.youtube.com/watch?v=H7tbjKFSg58&t=8s) (video)
- [Go Blueprint Code](https://github.com/Melkeydev/go-blueprint) - [Go Blueprint Web](https://go-blueprint.dev/)
- [ardanlabs](https://github.com/ardanlabs/service) Seems a bit overly complex but may be a good study example
  - seems to store vendors??? TODO
- [golang-rest-api-example](https://golang.cafe/blog/golang-rest-api-example.html)
- [Building REST APIs With Go 1.22 http.ServeMux](https://shijuvar.medium.com/building-rest-apis-with-go-1-22-http-servemux-2115f242f02b)
- CHI features
    - [Middleware and grouping with stdlib](https://gist.github.com/alexaandru/747f9d7bdfb1fa35140b359bf23fa820)
    - [reddit post on why still chi](https://www.reddit.com/r/golang/comments/1avn6ih/is_chi_relevant_anymore/)

#### new api examples 2025-03-12

- [go-rest-api-service-template](https://github.com/p2p-b2b/go-rest-api-service-template/tree/main)
- [go-zero](https://github.com/zeromicro/go-zero)
https://www.reddit.com/r/golang/comments/15y5wiq/lets_say_you_want_to_build_a_go_rest_api_should/
https://github.com/youngjun827/api-std-lib




### ogen

Similar problem as goa. Intro docs are light and not complete. Had limited patience for documentation that wasn't working.

###  [swag](https://github.com/swaggo/swag)

Seems to just be special comments in code. Ie writing the api spec inline with the code. Not bad but doesn't save typing. That's just typing in the same place.

### Additional urls

- [reddit go_openapi_codegen](https://www.reddit.com/r/golang/comments/1avsog1/go_openapi_codegen/)
- [reddit confused_by_the_openapi_options](https://www.reddit.com/r/golang/comments/1gmhz08/confused_by_the_openapi_options_for_go/)
- [Laurence de Jong](https://ldej.nl/post/generating-go-from-openapi-3/)
- [speakeasy](https://www.speakeasy.com/docs/languages/golang/oss-comparison-go) trying to sell speakeasy

## Decision


## Why / Notes


## Consequences

## Other Options

Possibilities: