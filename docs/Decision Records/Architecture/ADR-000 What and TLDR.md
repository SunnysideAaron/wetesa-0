# ADR-000 What and TLDR

## What we are building.

See "../Tech Stack/TSDR-000 What and TLDR.md"

## TL;DR of decisions

- ADR-001 Using Architecture Decision Records
  - Yes. For Architecture and Technology Stack.
- ADR-002 API Versioning
  - **PENDING**
- ADR-003 Semantic Versioning
  - Yes
- ADR-004 Database Columns
  - table_id not id
  - Use UUIDs for primary keys
  - PENDING Date columns
  - All times in UTC
- ADR-005 Error Handling
  - wrap errors
  - stack trace will be handled by slog
  - msg will be context only. Do not include calling or called function names. 
  - Don't use words like "error", "failed", "went wrong" "error occurred", "problem found", "failed to ..." in error messages.
  - don't use the ":" character anywhere else except the end of the message. 
- ADR-006 Logging Levels
  - Use default slog levels, in the following manor:
    - DEBUG (-4) Only turn on for in-depth troubleshooting.
    - INFO (0) default level in production. Enough information to troubleshoot basic problems.
    - WARN (4) Create a ticket. Something is wrong and needs fixing. Properly handled errors are info not warn.
    - ERROR (8) Call someone NOW! Something is wrong and needs immediate fixing.
  - Allow changing log level at runtime.
  - Allow different log levels for different parts of the code.
- ADR-006 Logging Output
  - log to STDOUT
- ADR-006 Logging Package
  - use slog
  - use LogAttrs()
- ADR-006 What to Log
  - Log request, response, and error stack trace
- ADR-007 Sensitive Information
  - Do not log sensitive information
- ADR-008 Audit Tables
  - **PENDING**
- ADR-009 Automated Testing
  - **PENDING**
- ADR-010 Package Layout
  - We can do what ever we want. See ADR-010 Package Layout for more notes.
- ADR-011 Validation
  - **Pending**
- ADR-012 Linters
  - [Staticcheck](https://staticcheck.dev/)
  - [revive](https://revive.run/)
  - [govulncheck](https://pkg.go.dev/golang.org/x/vuln/cmd/govulncheck)
- ADR-013 Go Coding Standards
  - [Uber](https://github.com/uber-go/guide/blob/master/style.md)
- ADR-014 Routing
  - Standard library for routing. Adaptor pattern for middleware.
- ADR-015 Middleware
  - Installed middleware:
    - AllowQuerySemicolons
    - CORS
    - IP
    - Logging
    - MaxBytesHandler 
    - Recover
    - RequestID
    - TimeoutHandler
- ADR-016 gRPC
  - Not using gRPC.
- ADR-017 API Base Design
  - **Pending**
- ADR-018 SQL Coding Standards
  - **Pending**
- ADR-019 CI/CD
  - Not using CI/CD. Yet.


## TODO 

**PENDING**

DB default encodings
utf8 (other encodings?)

UTC for server and db

ACID Compliant

DB knows user logged in. Not just a general web log in.

how to handle api breaking changes.