# ADR-002 Cryptography

## Status

Accepted / Pending

## Context
Often repeated advice is **not** to develop your own cryptography. It is an easy problem to get wrong and the consequences to getting it wrong can be catastrophic.

## Decision
We will not develop our own cryptography. Relying on well established 3rd party libraries and tools.

TODO Put which tools we decided on here.

TODO list what all needs cryptography

1. encrypt db?
2. SSN, DOB, Email, Address. (no one should ever be able to get these raw) Name?
3. User Log in

## Consequences
We will have to commit to keeping these tools up to date. Since they are well established vulnerabilities are quickly exploited.