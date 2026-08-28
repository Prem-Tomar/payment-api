# payment-api

High-volume Go REST service boundary for the Paykit payment portfolio.

This repository is the Go-service track. It will eventually expose payment APIs, validate and
sanitize request data, coordinate service concerns, and call the Rust domain libraries where
financial truth must stay deterministic.

## Ownership Boundary

Go owns:

- REST request and response boundaries.
- HTTP routing and middleware.
- request validation and sanitization.
- idempotency entry points.
- service orchestration.
- operational concerns such as health checks, metrics, timeouts, and deployment shape.

Go does not own:

- exact money arithmetic.
- fee calculation truth.
- payment lifecycle invariants.
- ledger correctness.
- HFT or memory-intensive processing.

Those remain Rust-owned unless the architecture is explicitly changed.

## Contribution Rule

No direct merges or pushes to `main` are allowed unless Prem explicitly approves that specific
case. Normal work must happen through a review branch and GitHub pull request.

## Current Status

This repo starts with pure Go implementation tasks. Architecture notes are maintained separately
by Prem and the coaching process; GitHub issues should ask for implementation work, not contract
or documentation deliverables.

For the initial Go days, do not add OpenAPI, generated contracts, processor integration, database
storage, Rust integration, Docker, or deployment scaffolding unless Prem explicitly scopes that
work.

The first implementation direction is captured in
[docs/tasks/day-g1-api-boundary.md](docs/tasks/day-g1-api-boundary.md).
