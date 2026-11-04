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

## Validation Rule

Always validate changes before pushing a PR update.

Run tests:

```sh
go test ./...
```

Run Go's built-in lint-style static checks:

```sh
go vet ./...
```

Format changed Go files before review. In PowerShell:

```powershell
gofmt -w (git ls-files '*.go')
```

In bash or zsh:

```sh
gofmt -w $(git ls-files '*.go')
```

## Learning And Coaching

Every new learning issue and successor issue includes a concrete Question of the Day
using the [learning task template](.github/ISSUE_TEMPLATE/learning-task.md).
Questions revisit earlier implementation work and connect it to system-design purpose,
component placement, and future applications. Answers belong in the relevant PR discussion.
Learning follow-ups are non-blocking; implementation defects remain review findings.
See [AGENTS.md](AGENTS.md) for the full coaching and review rules.

The [learning inventory](docs/curriculum/index.md) groups related goals into larger assignments.
Technical coaching is visual-first and uses plain English. Small interview exercises are active;
Moderate and Strong exercises are previews until the learner is comfortable. Read the
[coaching guide](docs/learning-and-interview-guide.md),
[software-design principle guide](docs/software-design-principles.md), and
[SQL learning guide](docs/sql-learning-guide.md).

## Current Status

The Gin implementation through G47 is under review in PR #63. Current learning work is grouped
in the inventory above. G65 starts a PostgreSQL adapter using the existing store boundary;
later database tasks depend on that working adapter. The curriculum update itself does not
implement learner-owned SQL or application behavior.

The first task is documented in [docs/tasks/day-g1-api-boundary.md](docs/tasks/day-g1-api-boundary.md).
The current Gin service foundation is documented in
[docs/tasks/day-g2-to-g9-gin-service-foundation.md](docs/tasks/day-g2-to-g9-gin-service-foundation.md).
