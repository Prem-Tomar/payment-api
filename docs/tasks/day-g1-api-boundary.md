# Day G1: Minimal Go Service Health Boundary

## Goal

Create the first minimal Go service boundary without implementing payment business logic.

## Real Problem

Payment APIs sit at the edge of the system. They receive external traffic, enforce transport
rules, protect the service from malformed input, and translate domain outcomes into HTTP responses.
They must be high-volume and operationally reliable, but they must not silently reimplement the
financial truth already owned by Rust libraries.

## Learning Target

- Go module layout.
- HTTP handler structure.
- Health and readiness endpoints.
- Table-driven handler tests.
- Keeping a service task small enough for review.

## Task

Implement a minimal Go HTTP service.

The implementation should include:

- a Go module.
- a small service entry point.
- `GET /healthz` for liveness.
- `GET /readyz` for readiness.
- simple JSON responses.
- table-driven tests for both endpoints.

## Expected Outcome

By the end of this task, another developer should have a tiny running service that proves:

- the repository compiles as a Go project.
- the service can be run locally.
- health and readiness behavior is tested.
- payment behavior is still intentionally absent.

## Out Of Scope

Do not implement:

- database access.
- authentication.
- idempotency storage.
- payment processor clients.
- ledger behavior.
- Rust integration.
- OpenAPI generation.
- Docker or deployment files.
- payment creation, capture, cancellation, refund, dispute, settlement, or reconciliation.

## Validation

- `go test ./...` passes.
- Health and readiness endpoints return expected status codes and JSON.
- The pull request explains implementation choices briefly in the PR description.
- No payment business endpoint is added.

## Suggested Day G2

Add graceful shutdown, request timeouts, and a small configuration shape without introducing
payment operations.
