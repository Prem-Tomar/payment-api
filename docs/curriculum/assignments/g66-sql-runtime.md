# SEQ-06 | G66: Integrate the SQL store with service lifecycle and errors

## Goal

Make SQL an explicit runtime option without mixing SQL details into handlers.

## Picture First

```mermaid
flowchart LR
  N0["Runtime config"]
  N1["Store constructor"]
  N2["Application use case"]
  N3["HTTP response"]
  N0 --> N1
  N1 --> N2
  N2 --> N3
```

Read left to right: Runtime config -> Store constructor -> Application use case -> HTTP response.

## Prerequisites

- SEQ-01: [issue #78](https://github.com/Prem-Tomar/payment-api/issues/78).
- SEQ-02: [issue #59](https://github.com/Prem-Tomar/payment-api/issues/59).
- SEQ-03: [issue #65](https://github.com/Prem-Tomar/payment-api/issues/65).
- SEQ-04: [issue #60](https://github.com/Prem-Tomar/payment-api/issues/60).
- SEQ-05: [issue #51](https://github.com/Prem-Tomar/payment-api/issues/51).

Start after G65. Fix the request-lifetime assignment's malformed deadline response before enabling real query calls in HTTP.

## Implementation Parts

1. Select memory or SQL storage at a clear construction point. Keep handlers dependent on application contracts.
2. Wire scoped lookup to the selected store. Keep placeholder creation honest until it has a supported write path.
3. Use a short context-aware database check for readiness in SQL mode. Keep the cheap health endpoint independent.
4. Map no-row, duplicate, cancellation, and unavailable-database outcomes without exposing SQL text or credentials. Close the shared handle during shutdown.

## Completion Criteria

- Changing storage mode does not require handler SQL.
- Operational routes and request cancellation remain predictable.
- Each failure produces one coherent HTTP response.
- SQL mode stops accepting work cleanly during shutdown.
- Show the connected success path and at least one relevant failure path in the PR discussion. A short explanation or demonstration is enough; no learner-written test suite or document is required.

## Scope

Use local learning data only. PostgreSQL stores operational/request metadata. Keep financial calculations, ledger rules, processor calls, and real payment lifecycle decisions outside this assignment.

Routine tests and documentation are reviewer-owned.

<details>
<summary>Learning discussion - not extra implementation requirements</summary>

## What, Why, And Where

- What: Make SQL an explicit runtime option without mixing SQL details into handlers.
- Why: the components must keep the same promise when they are connected.
- Where: follow the picture above through the existing project files.
- Earlier learning: Builds on G65, config, readiness injection, typed errors, and graceful shutdown.
- Later use: The same construction point can later add measured pool settings without changing route code.

Words used here: Composition root: the place that creates and connects dependencies. Adapter: code that fits one storage tool to an application interface.

## Software And Design Principle

Creator responsibility chooses who builds dependencies; a creational factory is only one possible mechanism. Single Responsibility keeps construction out of handlers.

Use a pattern only when it solves a visible problem. A Go function or small interface may be enough. [Short principle guide](../../software-design-principles.md).

## Database Learning Connection

```mermaid
flowchart LR
  N0["HTTP context"]
  N1["Application contract"]
  N2["Selected SQL adapter"]
  N3["PostgreSQL"]
  N0 --> N1
  N1 --> N2
  N2 --> N3
```

Read left to right: HTTP context -> Application contract -> Selected SQL adapter -> PostgreSQL.

This assignment includes database implementation. See the SQL practice guide for clause order, row prediction, failure behavior, and query-plan coaching.
[In-depth SQL guide](../../sql-learning-guide.md).

## Question of the Day

Previously learned concepts: Builds on G65, config, readiness injection, typed errors, and graceful shutdown.

```mermaid
flowchart LR
  N0["Existing interface"]
  N1["New adapter"]
  N2["Unchanged caller"]
  N0 --> N1
  N1 --> N2
```

Read left to right: Existing interface -> New adapter -> Unchanged caller.

The current readiness checker is injectable and the query service already accepts an interface. How do these earlier choices help replace a map without changing the HTTP contract?

Use your own words. Any programming language, pseudocode, or a sketch is welcome. If the question concerns Go or SQL semantics, explain that original behavior correctly.

## Interview Problem

```mermaid
flowchart LR
  N0["Database down"]
  N1["Health question"]
  N2["Readiness question"]
  N0 --> N1
  N1 --> N2
```

Read left to right: Database down -> Health question -> Readiness question.

The database is down. Decide what health and readiness should each report, and why they answer different questions.

## Research Hints

Search `Go sql.DB lifecycle`, `PingContext readiness`, `dependency injection composition root`, and `errors.Is database error mapping`.

Look up the concept, inspect the relevant project code, then make your own implementation decision. Hints are not a complete solution.

For discussion, any programming language, pseudocode, or diagram is accepted. QOD revisits earlier work; the interview examples explore different scenarios. [How learning answers are reviewed](../../learning-and-interview-guide.md).

</details>

