# ADR-002 API Versioning

## Status

Accepted

## Context

API Versioning can be tricky. Research to attempt something from 0.1
    
## Decision

URL versioning. Beginning with v0.1

## Why

This is just a basic example. It is important to have at least basic versioning
from day one. Eventually how to make breaking changes will have to be formalized.
A "contract" between developers and consumers of the API. That could be a long
way off. Basic URL versioning will allow for changes in that in between time of 
getting something out the door for people to test and committing to a contract.

Move to v1.0 only when we are ready to commit to a contract to our users.

## Notes

Part of versioning of API should state how long we will support that version.
Users will not move off unless forced to do so.

API version token? What is that versus url?

Min version of API in App

content negotiation
	api version, way to change. content of http. what is being sent over wire. how do changes happen

https://medium.com/better-programming/breaking-changes-in-apis-bf45ddfedba0

## Consequences

Not having a formal contract for breaking changes may lead to consumer frustration
down the road.

## Other Possible Options

- [Microsoft: RESTful web API design](https://learn.microsoft.com/en-us/azure/architecture/best-practices/api-design)
  - lists some different ways to implement versioning.
- [Googles API Versioning](https://google.aip.dev/185)

## Not an Option

    No versioning.
