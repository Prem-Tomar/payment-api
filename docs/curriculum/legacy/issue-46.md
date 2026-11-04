# Original Issue #46

Historical snapshot before inventory consolidation. This file preserves the original request; it is not completion evidence.

Title: Day G39: Add payment intent lifecycle state model

## Goal

Add a small payment-intent lifecycle state model so future handlers can describe where a payment intent is in the flow without pretending money movement has happened.

## Context

A lifecycle state is a named step in a resource's progress, such as `requires_payment_method`, `requires_confirmation`, or `created`. It helps the API communicate what can happen next without adding processor calls or ledger behavior.

A model here means a small application-level type that represents allowed states. It should not be a database table or a full domain engine yet.

## Task

Create a narrow payment-intent status/state type and use it in the application result shape where payment-intent responses are prepared.

Keep this focused on representing allowed states. Do not add a database, payment processor, ledger entry, Rust integration, or real payment confirmation.

## Acceptance Criteria

- Payment-intent responses use a deliberate status/state value instead of an arbitrary string.
- The allowed states are easy to find and explain.
- The state type stays separate from HTTP-only DTO concerns.
- Existing request validation and API-key behavior remain unchanged.
- The response still does not claim real money movement.

## Hint

Search for `payment intent lifecycle states`, `Go string enum pattern`, and `API resource status field`. Compare how a simple typed constant can keep response values consistent without building a full state machine yet.

## Question of the Day - G39

Previously learned concepts: Typed response statuses, application results, and placeholder services.

Scenario: Two application paths use slightly different strings for the same intended payment-intent state.

1. What could happen when a caller compares the state exactly? Does giving a string a named Go type automatically restrict all its possible values? Explain where responsibility for allowed values belongs.
2. Explain what this example teaches about responsibility or failure boundaries in the system, and relate it to one earlier implementation.

System-design connection: A state is a named stage in a resource's progress. Earlier typed results provide the shared vocabulary; a later confirmation implementation could use it to decide which actions are allowed.

Answer every part in your own words directly in the relevant PR description or comment; an issue-only answer does not satisfy this placement rule. Explain your assumptions and reasoning. No extra implementation, test-writing, or documentation is required to answer. The reviewer preserves your original answer and marks it Correct, Partially correct, Needs correction, or Awaiting answer, with targeted corrections and follow-ups. This learning response is expected but does not block merging an otherwise correct implementation.


