# Original Issue #66

Historical snapshot before inventory consolidation. This file preserves the original request; it is not completion evidence.

Title: Day G56: Preserve currency and payment method in application input

## Goal

Carry currency and payment method from the create request into the application input.

## Context

The HTTP request DTO contains the fields sent by the caller. The application input is the small value passed from the handler into the use case. Mapping is the step where the handler copies accepted request fields into that application input.

## Implementation Task

Extend the create-payment-intent application input with the request's currency and payment method, and map both fields in the Gin handler. Keep transport DTO types separate from application types; choose simple application field types that preserve the accepted values.

## Completion Criteria

- The application input preserves amount, currency, payment method, merchant reference, and customer reference from the accepted request.
- The handler remains the clear mapping point between JSON DTO and application input.
- No payment behavior is added; the placeholder response stays honest.

## Scope

Implementation only. Tests and documentation are handled by the reviewer, not required from the learner. No database, processor client, ledger behavior, Rust integration, or broad mapping framework.

## Research Hints

Search for `Go request DTO application command mapping`, `payment amount minor units int64`, and `Go named string type conversion`. Compare which values belong to the transport layer and which should cross into a use case.

## Question of the Day - G56

Previously learned concepts: The request DTO, payment-method field, and application input mapping.

Scenario: A valid request supplies amount, currency, and payment method, but the application input preserves only the amount.

1. What information has been lost at the boundary, even if the placeholder response looks successful? Why should preserving accepted values be considered separately from implementing real payment behavior?
2. Explain what this example teaches about responsibility or failure boundaries in the system, and relate it to one earlier implementation.

System-design connection: A boundary should carry the information its use case needs. Earlier DTOs describe the transport; later application behavior relies on explicit inputs rather than reaching back into the HTTP request.

Answer every part in your own words directly in the relevant PR description or comment; an issue-only answer does not satisfy this placement rule. Explain your assumptions and reasoning. No extra implementation, test-writing, or documentation is required to answer. The reviewer preserves your original answer and marks it Correct, Partially correct, Needs correction, or Awaiting answer, with targeted corrections and follow-ups. This learning response is expected but does not block merging an otherwise correct implementation.


