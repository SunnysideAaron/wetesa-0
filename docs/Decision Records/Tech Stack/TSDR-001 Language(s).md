# TSDR-001 Language(s) 

## Status

Accepted

## Context



## Decision

Go


## Why / Notes

  - Compiled, easy to read, easy to maintain, and fast.
  - Concurrency

## Consequences

Project is locked into Go.

## Other Options

Possibilities:
- Rust
  - Pro: Compiled. Go has better out of the box support for concurrency.
- Java
  - Oracle dependant.
- php
  - Script not compiled. Will be a bottle neck at scale.
- Python
  - Script not compiled. Will be a bottle neck at scale.
- Node.js
  - Javascript is a terrible language. It was a hack developed in a few days. As
    fast as it is, Node.js is at it's core just a way for javascript developers
    to not have to learn another language.

Not an option:
- c#
  - Microsoft dependant. Cost
- c++
  - Ecosystem is a mess. Not used for much web development today beyond highly
    optimized for speed code.
