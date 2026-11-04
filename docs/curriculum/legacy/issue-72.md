# Original Issue #72

Historical snapshot before inventory consolidation. This file preserves the original request; it is not completion evidence.

Title: Day G59: Scope idempotency records by merchant

## Goal

Make the idempotency record boundary explicitly merchant-scoped.

## Context

Idempotency associates a caller's retry key with one operation. Different merchants may choose the same key. Merchant scope identifies which caller's records a lookup belongs to; it is different from the merchant reference in a request body.

## Task

Extend the application idempotency record and its lookup boundary so merchant identity participates alongside the operation and retry key. Keep HTTP header parsing separate from this application shape.

## Completion Criteria

- A record can represent merchant identity, operation, and retry key.
- The lookup interface makes all three values explicit.
- The boundary can represent identical retry keys for different merchants without treating them as the same request.
- Current route behavior remains stable.

## Scope

Implementation only. Tests and documentation are reviewer-owned. No durable store, replay engine, real duplicate-payment prevention, processor integration, or authorization system.

## Research Hints

Search `multi tenant idempotency key scope`, `Go composite key struct`, and `application repository interface`. Compare a globally unique key with a key meaningful inside one caller's scope.

## Question of the Day - G59

Previously learned concepts: Idempotency-key parsing, merchant scope, and the application record shape.

Scenario: Two merchants use the same retry key for the same operation, but their request bodies differ.

1. Why can a globally keyed record confuse independent callers? Explain the role of merchant identity, operation, and request identity without treating the body merchant_ref as authorization.
2. Explain what this example teaches about responsibility or failure boundaries in the system, and relate it to one earlier implementation.

System-design connection: Idempotency means retries should not create accidental duplicate work. Earlier record boundaries represent the idea; later storage and replay behavior must preserve the caller's scope.

Answer every part in your own words directly in the relevant PR description or comment; an issue-only answer does not satisfy this placement rule. Explain your assumptions and reasoning. No extra implementation, test-writing, or documentation is required to answer. The reviewer preserves your original answer and marks it Correct, Partially correct, Needs correction, or Awaiting answer, with targeted corrections and follow-ups. This learning response is expected but does not block merging an otherwise correct implementation.


