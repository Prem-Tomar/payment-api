# Original Issue #51

Historical snapshot before inventory consolidation. This file preserves the original request; it is not completion evidence.

Title: Day G43: Add rate limit response header boundary

## Goal

Add basic rate-limit response headers so clients can understand the learning-grade request limit without reading server code.

## Context

A response header is metadata sent with the HTTP response. Rate-limit headers can tell a client the limit, remaining requests, or when to retry. This issue is about making the boundary clearer, not building a production quota system.

A quota is a policy for how many requests are allowed. This track is still using an in-memory learning implementation, so keep the headers simple and honest.

## Task

Extend the current rate-limit middleware to return small, consistent headers for allowed and rejected requests.

Keep this focused on response metadata. Do not add Redis, a database, paid plans, distributed rate limiting, processor clients, or Rust integration.

## Acceptance Criteria

- Rate-limited responses include a clear retry hint.
- Allowed responses expose simple limit metadata when practical.
- Header names are consistent and easy to explain.
- The implementation stays tied to the current in-memory learner limiter.
- Operational endpoints remain unaffected.

## Hint

Search for `HTTP rate limit headers`, `Retry-After header`, and `429 Too Many Requests`. Compare common headers, then choose a minimal set that matches this learning-grade limiter.

## Question of the Day - G43

Previously learned concepts: Merchant extraction, the token-bucket limiter, and shared-map locking.

Scenario: Two merchants each begin with five available request tokens. One merchant immediately sends three requests while the other has sent none.

1. What should each merchant's remaining allowance mean? Why must checking and consuming a token be protected together when requests run concurrently?
2. Explain what this example teaches about responsibility or failure boundaries in the system, and relate it to one earlier implementation.

System-design connection: A token permits one request and time refills the allowance. Earlier merchant scope separates callers; later client pacing can use response headers to respect that same policy.

Answer every part in your own words directly in the relevant PR description or comment; an issue-only answer does not satisfy this placement rule. Explain your assumptions and reasoning. No extra implementation, test-writing, or documentation is required to answer. The reviewer preserves your original answer and marks it Correct, Partially correct, Needs correction, or Awaiting answer, with targeted corrections and follow-ups. This learning response is expected but does not block merging an otherwise correct implementation.


