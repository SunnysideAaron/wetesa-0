# TSDR-000 What and TLDR

## What we are building.

An example CRUD api. With a minimum of dependencies.

## TL;DR of decisions

- TSDR-001 Language(s)
  - Go
- TSDR-002 Data storage
  - PostgreSQL
- TSDR-003 Docker
  - Using Bitnami PostgreSQL Image
- TSDR-004 API framework
  - Use the standard library.
- TSDR-005 SQL Driver
  - pgx
- TSDR-006 DB Initial Data Load
  - Postgress image docker initdb folder 
- TSDR-008 Live Reload of Code
  - [air](https://github.com/air-verse/air)
- TSDR-009 Non-dependencies
  - Do not attempt Authentication, Authorization, or Cryptography.
  - Don't use any of the database packages that are not in the standard library.
