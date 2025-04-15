# ADR-009 Automated Testing

## Status

Accepted

## Context

Automated testing gives developers confidence that code changes didn't break the function.

## Decision

Each API endpoint / method tested for happy path. This will verify a lot of the code.

Automated testing 100% coverage? No. But do test complex business logic. This
example doesn't have any. Would add some unit tests if it made sense.

## Why

Automated testing has burned me hard on multiple occasions. Need further research.

What is an appropriate level of testing for a solo dev?

## Notes

[Test Driven Development Wars: Detroit vs London, Classicist vs Mockist](https://medium.com/@adrianbooth/test-driven-development-wars-detroit-vs-london-classicist-vs-mockist-9956c78ae95f)

## Consequences

Anything not tested is at risk of breaking with changes. Testing too many things
costs time. Have to balance.

## Other Possible Options



## Not an Option

- [Learn Go with tests](https://quii.gitbook.io/learn-go-with-tests)
  - TDD Examples. Personally I believe TDD is only appropriate when the spec is
    well defined. Not a good fit for greenfield and a lot of unknowns.
- https://jestjs.io/docs/snapshot-testing
  - No UI for this example.
- [Awesome Go: Testing](https://github.com/avelino/awesome-go?tab=readme-ov-file#testing)
- [Awesome Go: Mock](https://github.com/avelino/awesome-go?tab=readme-ov-file#mock)  