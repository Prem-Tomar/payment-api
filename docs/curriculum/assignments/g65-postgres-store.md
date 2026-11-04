# G65: Build a working PostgreSQL store

## Goal

Deliver one usable SQL-backed store behind the existing application interface.

## Picture First

```mermaid
flowchart LR
  N0["Go configuration"]
  N1["Managed SQL pool"]
  N2["PostgreSQL table"]
  N3["Scoped Save and Get"]
  N0 --> N1
  N1 --> N2
  N2 --> N3
```

Read left to right: Go configuration -> Managed SQL pool -> PostgreSQL table -> Scoped Save and Get.

## Start And Dependencies

Ready to start now. Reuse the completed store interface and merchant-scoped lookup. You do not need to finish every queued HTTP task first.

## Implementation Parts

1. Create a configurable database/sql handle using the pgx standard-library driver. Bound the connection check with a context. Keep credentials out of logs.
2. Add a versioned SQL migration for learning record metadata. Store merchant identity, record ID, and creation time. Use deliberate types, required values, and scoped identity.
3. Implement Save with bound SQL parameters and a clear duplicate outcome. Do not build query text from caller values.
4. Implement merchant-scoped Get with a context, explicit columns, and correct no-row handling.

## Completion Criteria

- The same application store contract can use the SQL adapter.
- A saved row can be read by its merchant and ID after reconnecting.
- Another merchant cannot read that row; duplicate identity does not silently overwrite it.
- The database handle has a clear close owner. No HTTP route is switched to SQL in this task.
- Show the connected success path and at least one relevant failure path in the PR discussion. A short explanation or demonstration is enough; no learner-written test suite or document is required.

## What, Why, And Where

- What: Deliver one usable SQL-backed store behind the existing application interface.
- Why: the components must keep the same promise when they are connected.
- Where: follow the picture above through the existing project files.
- Earlier learning: Builds on the store interface, scoped map keys, config loading, and context propagation.
- Later use: G66 wires the store into runtime behavior. Later tasks add transactions and query analysis.

Words used here: SQL: a language for reading and changing tables. Migration: a saved schema change. Pool: reusable database connections. Parameter: a value passed separately from SQL text.

## Software And Design Principle

Dependency Inversion and Liskov Substitution: the SQL store must keep the existing store contract. GRASP Creator asks which component has the information to construct and own the store.

Use a pattern only when it solves a visible problem. A Go function or small interface may be enough. [Short principle guide](../../software-design-principles.md).

## Database Learning Connection

```mermaid
flowchart LR
  N0["Save or Get"]
  N1["SQL parameter values"]
  N2["Merchant-scoped row"]
  N3["Read after restart"]
  N0 --> N1
  N1 --> N2
  N2 --> N3
```

Read left to right: Save or Get -> SQL parameter values -> Merchant-scoped row -> Read after restart.

This assignment includes database implementation. See the SQL practice guide for clause order, row prediction, failure behavior, and query-plan coaching.
[In-depth SQL guide](../../sql-learning-guide.md).

## Question of the Day

Previously learned concepts: Builds on the store interface, scoped map keys, config loading, and context propagation.

```mermaid
flowchart LR
  N0["Same interface"]
  N1["Memory behavior"]
  N2["SQL behavior"]
  N3["Same promise?"]
  N0 --> N1
  N1 --> N2
  N2 --> N3
```

Read left to right: Same interface -> Memory behavior -> SQL behavior -> Same promise?.

The memory store and SQL store both implement one Go interface. Is successful compilation enough to prove they are interchangeable? Use missing records and duplicate records as examples.

Explain in your own words. Use short sentences, bullets, or a small sketch. We assess the reasoning, not English grammar.

## Interview Practice

**Active level: Small.** Start here. The mentor activates a harder level only after the learner can explain the simple case and an edge case, use a familiar example without step-by-step help, and says they are comfortable. Task number and time spent do not automatically raise the interview level.

### Small - active

```mermaid
flowchart LR
  N0["Save row"]
  N1["Restart API"]
  N2["Read record"]
  N0 --> N1
  N1 --> N2
```

Read left to right: Save row -> Restart API -> Read record.

Your API process restarts. Compare a map record and a SQL row. Which behavior must a client be told to expect?

<details>
<summary>Moderate - later practice, not active yet</summary>

```mermaid
flowchart LR
  N0["Merchant A pi-1"]
  N1["Stored identity"]
  N2["Merchant B pi-1"]
  N0 --> N1
  N1 --> N2
```

Read left to right: Merchant A pi-1 -> Stored identity -> Merchant B pi-1.

Two merchants both use record ID pi-1. Choose the identity rule and explain how a lookup stays inside one merchant's scope.

New concepts in this preview must be explained before this level is assigned.

</details>

<details>
<summary>Strong - future practice, not active yet</summary>

```mermaid
flowchart LR
  N0["Many requests"]
  N1["Shared handles"]
  N2["Connection budget"]
  N3["Database"]
  N0 --> N1
  N1 --> N2
  N2 --> N3
```

Read left to right: Many requests -> Shared handles -> Connection budget -> Database.

Later, many API instances open database connections at once. Explain why one sql.DB per request would be a poor ownership model and what capacity you would measure first.

The mentor introduces any new concepts first. This is a preview, not a requirement to implement another system now.

</details>

## Research Hints

Search `database/sql managing connections`, `pgx stdlib`, `PostgreSQL composite primary key`, and `QueryRowContext sql.ErrNoRows`.

Look up the concept, inspect the relevant project code, then make your own implementation decision. Hints are not a complete solution.

## Scope And Coaching Agreement

Use local learning data only. PostgreSQL stores operational/request metadata. Keep financial calculations, ledger rules, processor calls, and real payment lifecycle decisions outside this assignment.

Reviewers own routine tests and documentation. Implementation review checks the listed behavior. Questions and interview answers belong in the related PR description or comment, in the learner's own words. Preserve original answers and add focused corrections. Missing, partial, or incorrect learning answers are non-blocking for an otherwise correct implementation. No extra code, tests, or documentation are needed to answer.

