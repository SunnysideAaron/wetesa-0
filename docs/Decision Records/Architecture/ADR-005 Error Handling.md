# ADR-005 Error Handling

## Status

Accepted

## Context

How exactly to approach errors.

## Decision

- wrap errors
- stack trace will be handled by slog
- msg will be context only. Do not include calling or called function names. **Pending** reevaluate later.
- Don't use words like "error", "failed", "went wrong" "error occurred", "problem found", "failed to ..." in error messages. **TODO** Linter
- don't use the ":" character anywhere else except the end of the message. **TODO** Linter

For now this should be enough to get going. I'll reevaluate after using.

## Why / Notes

- [Standard Library Errors](https://pkg.go.dev/errors)
- [Standard Libary: http constants](https://pkg.go.dev/net/http#pkg-constants)
- [A concise guide to error handling in Go](https://medium.com/@andreiboar/a-concise-guide-to-error-handling-in-go-611a42e589ad)
- [Error Wrapping in Go: A Guide to Enhance Debugging](https://erik.cat/blog/error-wrapping-go/)
- [Reddit: In larger programs, how do you handle errors?](https://www.reddit.com/r/golang/comments/1iwmeaw/in_larger_programs_how_do_you_handle_errors_so/)
  - Lots of good discussion here. Summary:
  - wrap messages
  - Include function being called, function doing the calling, or context only. Pick one. Don't mix and match.
    - seems many devs prefer "function being called"
	- standard library uses the caller not what is being called. Is this just for legacy reasons?
  - don't prefix function names. 
    - fmt.Errorf("SomeFunc: run other func: %v", err)
    - Note that some packages in the standard library do this. like bufio. Effective Go as well? Is this just for legacy reasons?
	- this means you aren't adding context. refactor
  - Don't need a stack trace
    - there are solutions to add stack traces to errors. Since we are using slog and wrapping errors this wont be needed. With proper wrapping, errors aren't usually that deep in Go?
  - Don't use words like "error", "failed", "went wrong" "error occurred", "problem found", "failed to ..." in error messages.
  - don't use the ":" character anywhere else except the end of the message.
  - In context describe not what went wrong, but what were you doing. Couple of words is frequently enough. (do i agree?)
  - don't add function arguments to error messages. Callers will add if it needs. Instead add what the caller doesn't know. 
  - If a function returns more than one error always add context. If a function returns just 1 err somewhere you can consider omitting the error wrapping. (do i agree? seems safer to always add context especially with refactoring. function has 1 error today. tomorrow? tomorrows problem, refactor then?) 
  - many different solutions.
    - add all the context
	- make a ton of custom error types
	- make a bespoke error solution.
    - choosing to use panic and recovery as a project rule
	- choosing to make a custom error library that always adds context in a consistent way
	  - (requiring everyone to if err: customstuff.wrap(err))
	- use logging.
  - don't need 3rd party packages. Most were created before Context wrapping, multi errors, structured logging were added to standard library.
  - AI generated context will often break use ":" or "error", "failed" etc. don't trust.
- Consider using a UserError(). Errors for users.
- How does standard library handle errors?
  - bufio includes package in message. Does not seem like all packages do this. Pre wrapping?
  - ErrInvalidUnreadByte = errors.New("bufio: invalid use of UnreadByte")


## Additional Notes

- [Stop Fighting Go Errors: A Modern Approach to Handling Them](https://dev.to/zakariachahboun/mastering-error-handling-in-go-a-pragmatic-approach-leg)
  - [Simple strategy to understand error handling in Go ](https://www.reddit.com/r/golang/comments/1in0tiw/simple_strategy_to_understand_error_handling_in_go/)
- [Don’t just check errors, handle them gracefully](https://dave.cheney.net/2016/04/27/dont-just-check-errors-handle-them-gracefully) 2016-APR
  - Dated?
- [Sentinel errors and errors.Is() slow your code down by 500%](https://www.dolthub.com/blog/2024-05-31-benchmarking-go-error-handling/)
  - This may or may not be actually practical. Take with grain of salt and do further research.Yes any value conversion takes
   measurable time. In very specific cases this might speed some slow code. Seems like premature optimization. Some ideas
    OK like how to structure testing for errors for speed.

## Consequences



## Other Possible Options



## Not an Option

- [awesome-go: Error Handling](https://github.com/avelino/awesome-go?tab=readme-ov-file#error-handling)
- [simplerr](https://github.com/lobocv/simplerr)
  - [Advanced Go Error Handling Made Simple](https://blog.lobocv.com/posts/richer_golang_errors/) (2022-MAR)
- [errtrace](https://github.com/bracesdev/errtrace)
- [Handling errors LIKE a 10x ENGINEER in Golang - Golang Service Pattern](https://www.youtube.com/watch?v=CxcxRgwWtAk)
  - just not feeling it.