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
  - **PENDING**
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
- ADR-007 Authentication
  - **PENDING**
- ADR-008 Audit Tables
  - **PENDING**
- ADR-009 Automated Testing
  - **PENDING**
- ADR-010 Basic Code Layout of API
  - **PENDING**
- ADR-011 Validation
  - **Pending**
- ADR-012 Linters
  - **Pending**
- ADR-013 Coding Standards
  - **Pending**
- ADR-014 Routing
  - Standard library for routing. Adaptor pattern for middleware.
- ADR-015 Middleware
  - Installed middleware:
    - AllowQuerySemicolons
    - CORS
    - Logging
    - MaxBytesHandler 
    - Recover
    - RequestID
    - TimeoutHandler
- ADR-016 gRPC
  - Not using gRPC.


## TODO 

**PENDING**

DB default encodings
utf8 (other encodings?)

UTC for server and db

ACID Compliant

DB knows user logged in. Not just a general web log in.

how to handle api breaking changes.