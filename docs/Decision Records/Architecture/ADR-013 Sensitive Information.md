# ADR-007 Sensitive Information

## Status

Accepted

## Context

What to do with sensitive information.

## Decision

Do not log sensitive information, including:
- Passwords
- email addresses
- Phone numbers
- DOB
- Age
- Addresses
- Bank Accounts
- Credit card numbers
- Social security numbers
- Personal identification numbers (PINs)
- Health information
- Financial information
- Any other sensitive data. (add to this list)

**Pending** What to do if user accidentally puts password in user name field?

**Pending** Rough example on handleGetClient(). Get help on better implementing
the LogValuer interface. There seems to be something about LogValue not quite working as my understanding expects.
- or should db structs just have have a stringer method?
- Try looking into [ReplaceAttr](https://pkg.go.dev/log/slog#HandlerOptions.ReplaceAttr)

## Why

Why not log sensitive information? Because we should assume they will fall into
the wrong hands. Going for the logs is basic steps for any hacker. At the very
least we may be sending logs to a third party where they
are stored. Even if temporarily. With right to be deleted laws any personal info
logged also has to be able to be deleted. Easier to just not log it.

## Notes

## Consequences

Being sued for lots of money when someone's sensitive information is leaked.

