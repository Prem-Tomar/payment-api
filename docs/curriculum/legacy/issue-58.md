# Original Issue #58

Historical snapshot before inventory consolidation. This file preserves the original request; it is not completion evidence.

Title: Day G49: Add timeout error mapping at application boundary

## Goal

Make timeout and cancellation errors map cleanly between the application boundary and HTTP response layer.

## Context

A timeout means the allowed time budget expired. Cancellation means work was stopped because the request context ended. Mapping means converting an internal error into the right HTTP response without leaking implementation details.

This builds on request timeout middleware and context-aware use cases.

## Task

Add a small error mapping path for context deadline or cancellation coming from the create-payment-intent use case.

Keep this focused on error mapping. Do not add background jobs, queues, processor clients, databases, Rust integration, or fake long sleeps in mergeable handlers.

## Acceptance Criteria

- Context deadline errors can become a clear timeout HTTP response.
- Canceled requests do not look like ordinary validation failures.
- Internal error text is not leaked to clients.
- Existing validation and not-implemented behavior remains distinct.

## Hint

Search for `context deadline exceeded errors.Is`, `HTTP 504 timeout`, and `Go error mapping HTTP handler`. Pay attention to the difference between a timeout, a canceled request, and a server panic.

## Question of the Day - G49

Previously learned concepts: Request deadlines, context propagation, recovery middleware, and response helpers.

Scenario: An application call returns a deadline error. The handler writes a timeout body and then continues into its ordinary error branch.

1. What can happen to the status and body after a response has already been written? Why is a deadline error different from a panic, and which component should own the final response?
2. Explain what this example teaches about responsibility or failure boundaries in the system, and relate it to one earlier implementation.

System-design connection: One request needs one coherent outcome. Earlier middleware and helpers establish response ownership; later error mapping should preserve it across application calls.

Answer every part in your own words directly in the relevant PR description or comment; an issue-only answer does not satisfy this placement rule. Explain your assumptions and reasoning. No extra implementation, test-writing, or documentation is required to answer. The reviewer preserves your original answer and marks it Correct, Partially correct, Needs correction, or Awaiting answer, with targeted corrections and follow-ups. This learning response is expected but does not block merging an otherwise correct implementation.


