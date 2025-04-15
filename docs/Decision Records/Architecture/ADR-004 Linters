# ADR-012 Linters  

## Status

Accepted

## Context

Linters help catch code smells and enforce coding standards.

## Decision

Use these linters:
- [golangci-lint](https://golangci-lint.run/)
- [govulncheck](https://pkg.go.dev/golang.org/x/vuln/cmd/govulncheck) 

## Why
- [golangci-lint](https://golangci-lint.run/)
    - Runs many other linters. Including [Staticcheck](https://staticcheck.dev/) and [revive](https://revive.run/)
    - 2025 APR AaronS Did a first pass on which which linters and rules to use. Will need to live with them for a bit
      and adjust as needed. Someday in imaginary land of infinite time I'd like to do a full deep audit and document rule by rule.
    - [golangci-lint install](https://golangci-lint.run/welcome/install/)
      - Using "go install" isn't recommended. Since we are running it in a docker image I'm hoping it will
        be OK. Especially just for example code. Seems to be working at the moment. Adjust as needed.
    - [Welcome to golangci-lint v2](https://ldez.github.io/blog/2025/03/23/golangci-lint-v2/)
    - https://www.google.com/search?q=+golangci-lint+docker+image+install+guide&client=firefox-b-1-d&sca_esv=8bdb6913f2ceb8bc&sxsrf=AHTn8zpLWYmWxoJC2nRp2ONW8XEzw-rA8g%3A1743292561024&ei=kYjoZ_OeAZaB0PEPzZuPmQk&ved=0ahUKEwizq9vuvrCMAxWWADQIHc3NI5MQ4dUDCBA&uact=5&oq=+golangci-lint+docker+image+install+guide&gs_lp=Egxnd3Mtd2l6LXNlcnAiKSBnb2xhbmdjaS1saW50IGRvY2tlciBpbWFnZSBpbnN0YWxsIGd1aWRlMgUQIRigATIFECEYoAEyBRAhGKABMgUQIRigATIFECEYoAFI60JQ-AdY7kBwAXgAkAEAmAFuoAGIE6oBBDI3LjO4AQPIAQD4AQGYAh6gApMUwgIKEAAYsAMY1gQYR8ICChAjGIAEGCcYigXCAgQQIxgnwgIKEAAYgAQYQxiKBcICBhAAGAcYHsICBRAAGIAEwgILEAAYgAQYkQIYigXCAgoQABiABBgUGIcCwgIGEAAYFhgewgIIEAAYFhgKGB7CAgsQABiABBiGAxiKBcICBRAAGO8FwgIIEAAYgAQYogTCAggQABiiBBiJBcICBRAhGKsCmAMAiAYBkAYIkgcEMjQuNqAH5H4&sclient=gws-wiz-serp
    - https://hub.docker.com/r/golangci/golangci-lint/tags
    - https://www.reddit.com/r/golang/comments/1jepzes/alternatives_to_golangcilint_that_are_fast/
        - Be careful with versions of golangci compatible with your go compiler. If they mismatch it will be extra slow and take 100% CPU for minutes. 
        - What we did in my team was to only lint the changed files on push and lint all files inside CI/CD. And use the generated cache! 
        - there are other notes on how to only lint changed files.
        - Make sure NOT to install golangci via "go install"
    - [maratori Golden Config](https://gist.github.com/maratori/47a4d00457a92aa426dbd48a18776322)
    - [Oleg Kovalov](https://olegk.dev/go-linters-configuration-the-right-version)
    - https://www.reddit.com/r/golang/comments/1jjit5g/golangcilint_which_linters_do_you_enable_which/
- [govulncheck](https://pkg.go.dev/golang.org/x/vuln/cmd/govulncheck) 
  - Reports known vulnerabilities by checking the vulnerability database at https://vuln.go.dev
  - not a pure linter.  does not cache.
  - will never be part of golangci-lint. 
  - once a day is excessive. maybe just before release? Will need to be part of ci but not every build.

## Notes

- "Learning Go. An Idiomatic Approach. Real-world Go Programming 2ed 2024" by Bodner J. recommends the following linters.
  - [Staticcheck](https://staticcheck.dev/)
  - [revive](https://revive.run/)
  - [golangci-lint](https://golangci-lint.run/)
  - [govulncheck](https://pkg.go.dev/golang.org/x/vuln/cmd/govulncheck)

## Consequences

Restrictive in code formatting. Time spent "fixing" things that might not need to be fixed at this time.

## Other Possible Options

If golangci-lint is breaks these are a good minimum and easy to turn on.
- [Staticcheck](https://staticcheck.dev/)
- [revive](https://revive.run/)

## Not an Option

- https://github.com/fe3dback/go-arch-lint
  - I'm not feeling this will be a big problem. 
