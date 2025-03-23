# ADR-005 Error Handling

## Status

Accepted, Proposed, Deprecated or Superseded (list DR)

## Context

How exactly to handle errors. Wrapping, how to structure messages. stack trace?

https://blog.lobocv.com/posts/richer_golang_errors/

https://www.youtube.com/watch?v=CxcxRgwWtAk

https://www.reddit.com/r/golang/comments/1iwmeaw/in_larger_programs_how_do_you_handle_errors_so/

https://www.reddit.com/r/golang/comments/1in0tiw/simple_strategy_to_understand_error_handling_in_go/

pkg.errors = wrapping errors
	look into from podcast
	wrapping errror msg is about the thing you just called not the what is doing the calling.

waterfall library for msging?

https://pkg.go.dev/net/http#pkg-constants
  - lists constants for http status codes


## Decision



## Why / Notes



## Consequences



## Other Options

Possibilities:
https://github.com/avelino/awesome-go?tab=readme-ov-file#error-handling

Not an option:

