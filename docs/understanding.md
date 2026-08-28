# payment-api Understanding

## Goal

Define the Go service track clearly before writing service code.

The `payment-api` service will be the high-volume REST boundary for payment operations. It should
be fast, operationally simple, and easy to deploy, while relying on Rust crates for financial
correctness and heavy deterministic processing.

## Real Problem

Public payment APIs receive messy external input: JSON payloads, headers, idempotency keys,
authentication context, client mistakes, retries, timeouts, and duplicate requests. That boundary
must be robust, but it should not become the source of financial truth.

If Go duplicates money arithmetic, fee formulas, lifecycle rules, or ledger logic, the system can
drift into inconsistent behavior. The API should validate transport-level shape and then delegate
domain decisions to the Rust-owned contracts.

## Architecture Boundary

Go should handle:

- parsing HTTP requests.
- validating required request fields.
- rejecting malformed JSON and unsupported media types.
- normalizing simple transport data such as headers and request IDs.
- applying request-size limits.
- coordinating idempotency checks when persistence is introduced.
- mapping domain errors to stable HTTP responses.
- exposing health, readiness, metrics, and later tracing.

Rust should handle:

- exact currency and money representation.
- payment amount invariants.
- fee calculation and fee evidence.
- payment lifecycle transition rules.
- deterministic ledger behavior when that crate exists.
- memory-intensive or correctness-critical processing.

## Coaching Style For This Repo

Use the same learning style as the Rust track:

- Start with the goal and real systems problem.
- Explain the boundary before writing code.
- Keep each task small enough for daily progress.
- Review architecture/domain alignment separately from Go code quality.
- Teach the concept behind every review issue.
- Avoid expanding into new behavior without explicit scope.

## Initial Non-Goals

Do not add these until explicitly scoped:

- database schema or migrations.
- authentication or authorization.
- payment processor clients.
- ledger posting.
- refunds, disputes, settlement, or reconciliation.
- async workers or message queues.
- generated OpenAPI contracts.
- container or Kubernetes deployment.
- Rust FFI, WASM, or sidecar integration.

## First Implementation Direction

The first Go implementation should establish a minimal service skeleton only after the boundary is
reviewed. The first endpoint should probably be a health/readiness endpoint, not payment creation,
because it teaches service layout and validation discipline without pretending the payment workflow
is ready for public traffic.
