# Original Issue #20

Historical snapshot before inventory consolidation. This file preserves the original request; it is not completion evidence.

Title: Day G15: Add first payment intent route shell

Goal: connect the Gin routing, DTO, validation, and idempotency-key boundary into a first payment-intent route shell without creating payments.

Learning focus:
- Handler composition at the edge of a service.
- Distinguishing API shape from domain execution.

Implementation task:
- Add a versioned route shell for creating a payment intent.
- Bind the request DTO and apply boundary validation.
- Require the idempotency key boundary for this mutating route.
- Return an intentional placeholder response that makes it clear real creation is not implemented yet.

Rules:
- No actual payment creation.
- No ledger behavior.
- No database writes.
- No processor clients.
- No Rust integration.
- No OpenAPI contract yet.

Completion criteria:
- The route shape is visible and versioned.
- Request parsing and idempotency-header validation happen before the placeholder response.
- The response does not pretend that a payment was created.
- The next task can replace the placeholder with a real application boundary.

Research hints:
- Search for: `Gin POST route`, `HTTP 501 not implemented`, and `REST create resource status codes`.
- Think carefully about whether a placeholder route should return 501, 202, or another status.
- Keep the route honest: no fake successful payment state.

## Question of the Day - G15

Previously learned concepts: Request DTOs, validation, versioned routes, and idempotency-header parsing.

Scenario: A caller sends a valid request body but omits its retry key. Another caller supplies a retry key but malformed JSON.

1. Which boundary should stop each request, and why should neither reach application work? Does accepting a retry-key header alone prove that duplicate work cannot occur?
2. Explain what this example teaches about responsibility or failure boundaries in the system, and relate it to one earlier implementation.

System-design connection: Earlier parsing and middleware responsibilities keep the API predictable. Later application coordination can consume validated input; that future use is a preview.

Answer every part in your own words directly in the relevant PR description or comment; an issue-only answer does not satisfy this placement rule. Explain your assumptions and reasoning. No extra implementation, test-writing, or documentation is required to answer. The reviewer preserves your original answer and marks it Correct, Partially correct, Needs correction, or Awaiting answer, with targeted corrections and follow-ups. This learning response is expected but does not block merging an otherwise correct implementation.


