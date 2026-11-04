# Original Issue #76

Historical snapshot before inventory consolidation. This file preserves the original request; it is not completion evidence.

Title: Day G63: Correlate API error responses with request IDs

## Goal

Connect client-visible API errors to the request ID already used in server logs.

## Context And System Design

Correlation means linking records about the same request. A request ID is an identifier for that request, not a credential or merchant identity.

Earlier middleware and logging already created the correlation path. The response boundary can expose its identifier; later observability tools can use the same association without requiring a tracing system now.

## Implementation Task

Include the established request ID in errors emitted by the shared response helper when it is available. Keep response creation from generating a different ID or exposing request secrets.

## Completion Criteria

- Errors reuse the request ID established by middleware.
- Missing request-ID context has a small deliberate fallback.
- Success response behavior and merchant scope stay unchanged.

## Scope

Implementation only. Tests and documentation are reviewer-owned. No real processor, database, ledger behavior, financial calculations, production observability system, or general framework is required.

## Research Hints

Search `request correlation ID error response`, `Gin context middleware value`, and `structured logging request id`.

## Question of the Day - G63

Previously learned concepts: Request-ID middleware, structured access logs, and shared error helpers.

Scenario: A caller reports a generic error message, and several log entries share that same message.

1. Which earlier request metadata can connect that report to one request? Why would generating a new ID while writing the error break that connection, and why is a request ID not proof of identity?
2. Explain the system-design responsibility involved and connect it to earlier implemented work. Future uses in the context above are previews, not assumed knowledge.

Answer every part in your own words directly in the relevant PR description or comment. An issue-only answer does not meet the placement rule. No extra implementation, tests, or documentation are required to answer. The reviewer preserves the answer, identifies correct and incomplete parts, explains corrections, and asks targeted follow-ups. The learning response is expected but does not block merging an otherwise correct implementation.


