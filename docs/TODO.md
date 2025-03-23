# TODO

- [x] [The NewServer constructor](https://grafana.com/blog/2024/02/09/how-i-write-http-services-in-go-after-13-years/#the-newserver-constructor)
  - [x] [Long argument lists](https://grafana.com/blog/2024/02/09/how-i-write-http-services-in-go-after-13-years/#long-argument-lists)
- [x] [Map the entire API surface in routes.go](https://grafana.com/blog/2024/02/09/how-i-write-http-services-in-go-after-13-years/#map-the-entire-api-surface-in-routesgo)
- [x] [func main() only calls run()](https://grafana.com/blog/2024/02/09/how-i-write-http-services-in-go-after-13-years/#func-main-only-calls-run)
  - [x] [Gracefully shutting down](https://grafana.com/blog/2024/02/09/how-i-write-http-services-in-go-after-13-years/#gracefully-shutting-down)
  - [x] [Controlling the environment](https://grafana.com/blog/2024/02/09/how-i-write-http-services-in-go-after-13-years/#controlling-the-environment)
- [x] [Maker funcs return the handler](https://grafana.com/blog/2024/02/09/how-i-write-http-services-in-go-after-13-years/#maker-funcs-return-the-handler)
- [x] [Handle decoding/encoding in one place](https://grafana.com/blog/2024/02/09/how-i-write-http-services-in-go-after-13-years/#handle-decodingencoding-in-one-place)
- [x] [Validating data](https://grafana.com/blog/2024/02/09/how-i-write-http-services-in-go-after-13-years/#validating-data)

end point that sits and waits so can verify that context cancels
404 not found vs default home url

- [ ] [The adapter pattern for middleware](https://grafana.com/blog/2024/02/09/how-i-write-http-services-in-go-after-13-years/#the-adapter-pattern-for-middleware)
- [ ] [Sometimes I return the middleware](https://grafana.com/blog/2024/02/09/how-i-write-http-services-in-go-after-13-years/#sometimes-i-return-the-middleware)
  - [ ] [An opportunity to hide the request/response types away](https://grafana.com/blog/2024/02/09/how-i-write-http-services-in-go-after-13-years/#an-opportunity-to-hide-the-requestresponse-types-away)

## Broader go as api todo later
- https://eblog.fly.dev/backendbasics.html
  - https://www.reddit.com/r/golang/comments/1jehpba/starting_systems_programming_pt_1_programmers/
- [go-rest-api-service-template ](https://github.com/p2p-b2b/go-rest-api-service-template)  
- Alex Edwards’ book “Let’s Go” which features a CRUD app
- [ardanlabs](https://github.com/ardanlabs/service) Seems a bit overly complex but may be a good study example
  - seems to store vendors??? TODO

## Todo later after implementing testing

- [ ] [Use inline request/response types for additional storytelling in tests](https://grafana.com/blog/2024/02/09/how-i-write-http-services-in-go-after-13-years/#use-inline-requestresponse-types-for-additional-storytelling-in-tests)
- [ ] [sync.Once to defer setup](https://grafana.com/blog/2024/02/09/how-i-write-http-services-in-go-after-13-years/#synconce-to-defer-setup)
- [ ] [Designing for testability](https://grafana.com/blog/2024/02/09/how-i-write-http-services-in-go-after-13-years/#designing-for-testability)
  - [ ] [What is the unit when unit testing?](https://grafana.com/blog/2024/02/09/how-i-write-http-services-in-go-after-13-years/#what-is-the-unit-when-unit-testing)
  - [ ] [Testing with the run function](https://grafana.com/blog/2024/02/09/how-i-write-http-services-in-go-after-13-years/#testing-with-the-run-function)
  - [ ] [Waiting for readiness](https://grafana.com/blog/2024/02/09/how-i-write-http-services-in-go-after-13-years/#waiting-for-readiness)

## TODO wetesa 1
- https://github.com/spf13/viper
  - allows for hot loading of config.
  - what other features would we actually use?
- [Awesome Go - Database](https://github.com/avelino/awesome-go?tab=readme-ov-file#database)
https://github.com/avelino/awesome-go?tab=readme-ov-file#date-and-time
https://github.com/avelino/awesome-go?tab=readme-ov-file#email
https://github.com/avelino/awesome-go?tab=readme-ov-file#file-handling
https://github.com/avelino/awesome-go?tab=readme-ov-file#financial
  - currency is different in go because of floats
https://github.com/avelino/awesome-go?tab=readme-ov-file#forms
https://github.com/avelino/awesome-go?tab=readme-ov-file#geographic

https://github.com/avelino/awesome-go?tab=readme-ov-file#goroutines
https://github.com/avelino/awesome-go?tab=readme-ov-file#images
https://github.com/avelino/awesome-go?tab=readme-ov-file#job-scheduler
https://github.com/avelino/awesome-go?tab=readme-ov-file#json
https://github.com/avelino/awesome-go?tab=readme-ov-file#messaging
https://github.com/avelino/awesome-go?tab=readme-ov-file#microsoft-office
https://github.com/avelino/awesome-go?tab=readme-ov-file#strings
https://github.com/avelino/awesome-go?tab=readme-ov-file#uncategorized
https://github.com/avelino/awesome-go?tab=readme-ov-file#http-clients 
  - maybe for testing?
https://github.com/avelino/awesome-go?tab=readme-ov-file#package-management
  - might have something for backing up dependencies
https://github.com/avelino/awesome-go?tab=readme-ov-file#performance
https://github.com/avelino/awesome-go?tab=readme-ov-file#query-language
https://github.com/avelino/awesome-go?tab=readme-ov-file#science-and-data-analysis
https://github.com/avelino/awesome-go?tab=readme-ov-file#security
https://github.com/avelino/awesome-go?tab=readme-ov-file#text-processing
  - see package "address" and "sq", others?
https://github.com/avelino/awesome-go?tab=readme-ov-file#regular-expressions
https://github.com/avelino/awesome-go?tab=readme-ov-file#third-party-apis
https://github.com/avelino/awesome-go?tab=readme-ov-file#utilities
https://github.com/avelino/awesome-go?tab=readme-ov-file#go-tools
https://github.com/avelino/awesome-go?tab=readme-ov-file#devops-tools
https://github.com/avelino/awesome-go?tab=readme-ov-file#other-software

## TRAINING
https://github.com/avelino/awesome-go?tab=readme-ov-file#conferences
https://github.com/avelino/awesome-go?tab=readme-ov-file#e-books
https://github.com/avelino/awesome-go?tab=readme-ov-file#meetups
https://github.com/avelino/awesome-go?tab=readme-ov-file#tutorials
https://github.com/avelino/awesome-go?tab=readme-ov-file#guided-learning






