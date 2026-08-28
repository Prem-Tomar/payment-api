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

This repo is intentionally documentation-first. No Go module, executable service, endpoints,
database, processor adapter, or deployment scaffold has been added yet.

The first task is documented in [docs/tasks/day-g1-api-boundary.md](docs/tasks/day-g1-api-boundary.md).
