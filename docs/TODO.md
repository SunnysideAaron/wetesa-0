# TODO

- [x] [The NewServer constructor](https://grafana.com/blog/2024/02/09/how-i-write-http-services-in-go-after-13-years/#the-newserver-constructor)
  - [x] [Long argument lists](https://grafana.com/blog/2024/02/09/how-i-write-http-services-in-go-after-13-years/#long-argument-lists)
- [x] [Map the entire API surface in routes.go](https://grafana.com/blog/2024/02/09/how-i-write-http-services-in-go-after-13-years/#map-the-entire-api-surface-in-routesgo)
- [x] [func main() only calls run()](https://grafana.com/blog/2024/02/09/how-i-write-http-services-in-go-after-13-years/#func-main-only-calls-run)
  - [x] [Gracefully shutting down](https://grafana.com/blog/2024/02/09/how-i-write-http-services-in-go-after-13-years/#gracefully-shutting-down)
  - [x] [Controlling the environment](https://grafana.com/blog/2024/02/09/how-i-write-http-services-in-go-after-13-years/#controlling-the-environment)
- [x] [Maker funcs return the handler](https://grafana.com/blog/2024/02/09/how-i-write-http-services-in-go-after-13-years/#maker-funcs-return-the-handler)
- [x] [Handle decoding/encoding in one place](https://grafana.com/blog/2024/02/09/how-i-write-http-services-in-go-after-13-years/#handle-decodingencoding-in-one-place)
- [ ] [Validating data](https://grafana.com/blog/2024/02/09/how-i-write-http-services-in-go-after-13-years/#validating-data)
- [ ] [The adapter pattern for middleware](https://grafana.com/blog/2024/02/09/how-i-write-http-services-in-go-after-13-years/#the-adapter-pattern-for-middleware)
- [ ] [Sometimes I return the middleware](https://grafana.com/blog/2024/02/09/how-i-write-http-services-in-go-after-13-years/#sometimes-i-return-the-middleware)
  - [ ] [An opportunity to hide the request/response types away](https://grafana.com/blog/2024/02/09/how-i-write-http-services-in-go-after-13-years/#an-opportunity-to-hide-the-requestresponse-types-away)

## Todo later after implementing testing

- [ ] [Use inline request/response types for additional storytelling in tests](https://grafana.com/blog/2024/02/09/how-i-write-http-services-in-go-after-13-years/#use-inline-requestresponse-types-for-additional-storytelling-in-tests)
- [ ] [sync.Once to defer setup](https://grafana.com/blog/2024/02/09/how-i-write-http-services-in-go-after-13-years/#synconce-to-defer-setup)
- [ ] [Designing for testability](https://grafana.com/blog/2024/02/09/how-i-write-http-services-in-go-after-13-years/#designing-for-testability)
  - [ ] [What is the unit when unit testing?](https://grafana.com/blog/2024/02/09/how-i-write-http-services-in-go-after-13-years/#what-is-the-unit-when-unit-testing)
  - [ ] [Testing with the run function](https://grafana.com/blog/2024/02/09/how-i-write-http-services-in-go-after-13-years/#testing-with-the-run-function)
  - [ ] [Waiting for readiness](https://grafana.com/blog/2024/02/09/how-i-write-http-services-in-go-after-13-years/#waiting-for-readiness)
