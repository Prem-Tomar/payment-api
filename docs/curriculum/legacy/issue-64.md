# Original Issue #64

Historical snapshot before inventory consolidation. This file preserves the original request; it is not completion evidence.

Title: Day G54: Reject trailing JSON in create payment intent requests

## Goal

Reject a second JSON value or other trailing data in a create-payment-intent request body.

## Context

A JSON value is one complete document, such as one object. Some decoders stop after that first value even when more bytes follow. The API should accept exactly one request object so the caller's full input is handled consistently.

## Implementation Task

Refine the Gin create-payment-intent parsing path to accept one valid object followed by optional whitespace, and reject additional JSON values or trailing junk with the existing invalid-request error response. Preserve the body-size guard and field validation.

## Completion Criteria

- One valid request object is accepted with or without trailing whitespace.
- A second JSON value and non-whitespace trailing junk are rejected before the application use case runs.
- Oversized requests and field validation retain their existing error categories.
- The change stays small and easy to trace.

## Scope

Implementation only. Tests and documentation are handled by the reviewer, not required from the learner. No payment processing, persistence, ledger behavior, Rust integration, or parsing framework.

## Research Hints

Search for `Go json Decoder trailing data`, `json Decoder EOF second decode`, and `Gin ShouldBindJSON decoder`. EOF means end of input. Compare accepting the first document with validating the entire body.

## Question of the Day - G54

Previously learned concepts: The body-size guard, JSON binding, and input validation.

Scenario: A request is below the body-size limit and contains one valid JSON object followed by another object.

1. What does the size guard prove, and what does it not prove about the complete document? Why might successfully decoding one value still be insufficient to accept the entire request body?
2. Explain what this example teaches about responsibility or failure boundaries in the system, and relate it to one earlier implementation.

System-design connection: A byte limit controls input volume; parsing controls structure. These earlier boundaries solve different problems and later need to work together for a clear request contract.

Answer every part in your own words directly in the relevant PR description or comment; an issue-only answer does not satisfy this placement rule. Explain your assumptions and reasoning. No extra implementation, test-writing, or documentation is required to answer. The reviewer preserves your original answer and marks it Correct, Partially correct, Needs correction, or Awaiting answer, with targeted corrections and follow-ups. This learning response is expected but does not block merging an otherwise correct implementation.


