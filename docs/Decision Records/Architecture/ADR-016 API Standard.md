# ADR-017 API Standard

## Status

Accepted

## Context

There is no standard. We do want to be consistent throughout the API.

## Decision

- [API Standards by Michael Bissell](https://www.michaelbissell.com/2d5a25c0-8d0c-11ed-b6fc-b5eee5a22130/API-Standards)
  - like this. Short and sweet. Use this ish.
- [Set the right tone (for error messages)](https://developers.google.com/tech-writing/error-messages/set-tone)

TODO subcollections, totalCount, created, and modified (date) in response. This 
is all waiting on other code I may or may not include in this example.

TODO field selection

TODO PATCH

TODO long running asynchronous methods, 202 accepted but not completed.

## Why / Notes

As the API grows we will probably need to tighten standards. Start simple till we
know what we need.

## Consequences

## Other Possible Options

- [Microsoft: RESTful web API design](https://learn.microsoft.com/en-us/azure/architecture/best-practices/api-design)
  - Implement asynchronous methods, 202 accepted
  - field selection
- [Zalando RESTful API and Event Guidelines](https://opensource.zalando.com/restful-api-guidelines/)
  - Good stuff but very opinionated.
- [Google API Improvement Proposals](https://google.aip.dev/)
- [Google Cloud: API design guide](https://cloud.google.com/apis/design)
  - Some good info but a bit much.
  - "the Google API Style Guide is (unofficially) deprecated" true? 
    [reddit](https://www.reddit.com/r/ExperiencedDevs/comments/vc8em5/do_you_have_an_api_design_guide/)

## Not an Option

