# Original Issue #70

Historical snapshot before inventory consolidation. This file preserves the original request; it is not completion evidence.

Title: Day G57: Add consistent middleware error codes

## Goal

Return consistent machine-readable error codes from API-key and merchant-scope middleware.

## Context

Middleware checks a request before its handler runs. A machine-readable code is a stable string a client can inspect without parsing a human message. These middleware errors currently have a different response shape from handler errors.

## Task

Make missing API-key, missing merchant-scope, and malformed merchant-scope responses use the shared error-code shape while preserving their existing HTTP statuses. Keep shared response code usable across the current packages.

## Completion Criteria

- Each listed failure includes an intentional code and readable message.
- Handler and middleware errors use a consistent JSON shape.
- Valid requests and operational routes keep their behavior.
- Package dependencies remain easy to trace.

## Scope

Implementation only. Tests and documentation are reviewer-owned. No credential database, authorization engine, real merchant registry, or error framework.

## Research Hints

Search `Go shared HTTP error response`, `Gin middleware error response`, and `Go import cycle`. An import cycle means packages depend back on each other; consider where a small shared response responsibility belongs.

## Question of the Day - G57

Previously learned concepts: Shared JSON helpers, API-key middleware, merchant scope, and typed error codes.

Scenario: A handler error has a code and message, but a middleware error contains only a message.

1. What burden does this put on a client handling both outcomes? Where can shared response behavior live without making packages depend back on each other?
2. Explain what this example teaches about responsibility or failure boundaries in the system, and relate it to one earlier implementation.

System-design connection: An import cycle means packages depend back on each other. Earlier package separation and response helpers should support a consistent API contract; later clients can reuse one error-handling path.

Answer every part in your own words directly in the relevant PR description or comment; an issue-only answer does not satisfy this placement rule. Explain your assumptions and reasoning. No extra implementation, test-writing, or documentation is required to answer. The reviewer preserves your original answer and marks it Correct, Partially correct, Needs correction, or Awaiting answer, with targeted corrections and follow-ups. This learning response is expected but does not block merging an otherwise correct implementation.


