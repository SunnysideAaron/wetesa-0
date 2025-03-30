# ADR-010 Basic Coding Layout of API

## Status

Accepted

## Context

## Decision

Go does not have a standard layout. We can do what ever we want.

For now wee will use:

### /cmd

### /internal

Internal is not importable by other packages. This is enforced by go. 
- "Learning Go (An Idiomatic Approach to Real-World Programming)" by Jon Bodner page 234
- [reddit: internal package](https://www.reddit.com/r/golang/comments/1ak3mmm/internal_package/)

### /internal/config

Standard practice to put loading of configuration in a separate config package.

### /internal/database

### /internal/server

## Why / Notes

## Consequences

## Other Options

### Possibilities:

- https://github.com/avelino/awesome-go?tab=readme-ov-file#project-layout
- [reddit: internal package](https://www.reddit.com/r/golang/comments/1ak3mmm/internal_package/)
  - lists several options in a comment.
- [project-layout](https://github.com/golang-standards/project-layout)
- [Organizing a Go module](https://go.dev/doc/modules/layout#server-project)

Example code:

- [exposure-notifications-server](https://github.com/google/exposure-notifications-server)
- [AutoVerse: A Modular Go Framework for RESTful APIs](https://github.com/Muga20/Go-Modular-Application)
  - didn't like this

### Not an option:

Vertical Slice Architecture?
- https://www.jimmybogard.com/vertical-slice-architecture/
- https://www.milanjovanovic.tech/blog/vertical-slice-architecture
- Decided not to use Vertical Slice Architecture because there are many packages
 that do code generation. When this example gets forked it will be easier to use
  those packages. In theory.