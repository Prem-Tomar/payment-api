# Original Issue #77

Historical snapshot before inventory consolidation. This file preserves the original request; it is not completion evidence.

Title: Day G64: Pass request context into readiness checks

## Goal

Make the injectable readiness boundary receive the HTTP request's lifetime signal.

## Context And System Design

Readiness describes whether the service can accept work. A context carries cancellation and deadlines; dependency checks may eventually need to stop when that request ends.

Earlier injected checks separated readiness from HTTP and earlier contexts carried request lifetime. Their connection belongs at the checker interface; later real dependency checks can cooperate with cancellation.

## Implementation Task

Extend the existing readiness-check interface to receive a context and pass the request context from the readiness handler. Keep the default check and operational route behavior simple.

## Completion Criteria

- The handler passes its request context to the checker.
- The default checker remains cheap and explicit.
- Success and unavailable responses remain predictable without adding real dependencies.

## Scope

Implementation only. Tests and documentation are reviewer-owned. No real processor, database, ledger behavior, financial calculations, production observability system, or general framework is required.

## Research Hints

Search `Go context injectable health check`, `readiness check request context`, and `context cancellation propagation`.

## Question of the Day - G64

Previously learned concepts: Injectable readiness checks, per-request context, and server timeouts.

Scenario: A readiness dependency check ignores its caller's cancellation signal and keeps waiting after the caller has gone away.

1. Does merely creating a request deadline stop that check? Explain what passing and observing context each contribute, and how this differs from a server's socket read timeout.
2. Explain the system-design responsibility involved and connect it to earlier implemented work. Future uses in the context above are previews, not assumed knowledge.

Answer every part in your own words directly in the relevant PR description or comment. An issue-only answer does not meet the placement rule. No extra implementation, tests, or documentation are required to answer. The reviewer preserves the answer, identifies correct and incomplete parts, explains corrections, and asks targeted follow-ups. The learning response is expected but does not block merging an otherwise correct implementation.


