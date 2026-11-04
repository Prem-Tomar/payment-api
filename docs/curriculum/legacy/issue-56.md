# Original Issue #56

Historical snapshot before inventory consolidation. This file preserves the original request; it is not completion evidence.

Title: Day G48: Add request metadata input to create use case

## Goal

Pass request metadata into the create-payment-intent application boundary instead of only building it in the HTTP handler.

## Context

Request metadata is information about the request, such as request ID and idempotency key. Body DTO data is the JSON payload. Keeping these separate helps the application boundary receive the context it needs without mixing unrelated data.

This builds on the metadata struct already introduced.

## Task

Update the create-payment-intent use case boundary so it can receive a small metadata input alongside request data or context.

Keep this focused on boundary shape. Do not add persistence, idempotency records, payment processor calls, ledger behavior, Rust integration, or tracing systems.

## Acceptance Criteria

- The handler passes request metadata into the application boundary.
- Metadata remains separate from the JSON body DTO.
- Missing required metadata still fails before application work continues.
- Placeholder create behavior remains honest.

## Hint

Search for `Go application service input struct`, `request metadata struct`, and `Gin context value`. Compare passing many function arguments with passing a small input struct.

## Question of the Day - G48

Previously learned concepts: Request IDs, idempotency keys, merchant scope, and request metadata.

Scenario: A request has a request ID, a retry key, a merchant header, and a merchant_ref field in its JSON body.

1. Explain the distinct purpose of each value. Why should a body reference not silently replace the merchant scope or the request ID used by logging?
2. Explain what this example teaches about responsibility or failure boundaries in the system, and relate it to one earlier implementation.

System-design connection: Metadata describes the request rather than its business payload. Earlier middleware extracts these values; later application coordination can receive them explicitly without depending on Gin.

Answer every part in your own words directly in the relevant PR description or comment; an issue-only answer does not satisfy this placement rule. Explain your assumptions and reasoning. No extra implementation, test-writing, or documentation is required to answer. The reviewer preserves your original answer and marks it Correct, Partially correct, Needs correction, or Awaiting answer, with targeted corrections and follow-ups. This learning response is expected but does not block merging an otherwise correct implementation.


