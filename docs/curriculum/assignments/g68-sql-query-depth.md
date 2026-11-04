# G68: Build and explain useful SQL queries

## Goal

Deliver a small read boundary that explains stored learning operations accurately and efficiently.

## Picture First

```mermaid
flowchart LR
  N0["Scoped read request"]
  N1["SQL rows and joins"]
  N2["Bounded result"]
  N3["Measured query plan"]
  N0 --> N1
  N1 --> N2
  N2 --> N3
```

Read left to right: Scoped read request -> SQL rows and joins -> Bounded result -> Measured query plan.

## Start And Dependencies

Start after G67. Use only synthetic/local learning records. A query plan is database work to inspect, not a number to guess.

## Implementation Parts

1. Implement one scoped read query with explicit columns, filters, a stable order, and a bounded page size.
2. Add a read result that joins learning records with idempotency records. Handle missing matches and NULL values deliberately.
3. Add a small grouped summary of learning operations. Explain the difference between filtering rows and filtering groups.
4. Inspect the read query with EXPLAIN on representative local data. Add an appropriate index only after comparing the plan and write cost.

## Completion Criteria

- Merchant scope applies to every query and join.
- Paging has a stable order; NULL and missing matches do not silently remove needed rows.
- Summary counts match the raw learning rows.
- The chosen index is justified by actual plans. A sequential scan on a small table is not automatically a defect.
- Show the connected success path and at least one relevant failure path in the PR discussion. A short explanation or demonstration is enough; no learner-written test suite or document is required.

## What, Why, And Where

- What: Deliver a small read boundary that explains stored learning operations accurately and efficiently.
- Why: the components must keep the same promise when they are connected.
- Where: follow the picture above through the existing project files.
- Earlier learning: Builds on SQL insert/select, scoped records, linked retry rows, and application read boundaries.
- Later use: G69-G74 continue with updates, paging under concurrent writes, isolation, locks, schema changes, and recovery.

Words used here: Join: combine related rows. NULL: an absent or unknown value. Aggregate: combine several rows into a summary. Index: a structure that can help find rows. Query plan: the work the database chooses to run a query.

## Software And Design Principle

Interface Segregation and Single Responsibility: give read callers a small query contract rather than a large storage API. An index is a storage choice, not a new domain rule.

Use a pattern only when it solves a visible problem. A Go function or small interface may be enough. [Short principle guide](../../software-design-principles.md).

## Database Learning Connection

```mermaid
flowchart LR
  N0["Base rows"]
  N1["Join and group"]
  N2["Filter and order"]
  N3["Explain the plan"]
  N0 --> N1
  N1 --> N2
  N2 --> N3
```

Read left to right: Base rows -> Join and group -> Filter and order -> Explain the plan.

This assignment includes database implementation. See the SQL practice guide for clause order, row prediction, failure behavior, and query-plan coaching.
[In-depth SQL guide](../../sql-learning-guide.md).

## Question of the Day

Previously learned concepts: Builds on SQL insert/select, scoped records, linked retry rows, and application read boundaries.

```mermaid
flowchart LR
  N0["No base row"]
  N1["Not found"]
  N2["Matched row with NULL"]
  N0 --> N1
  N1 --> N2
```

Read left to right: No base row -> Not found -> Matched row with NULL.

A prior query finds no row and returns a not-found result. A LEFT JOIN can instead produce a row with missing columns. Why are these different cases for a Go result mapper?

Explain in your own words. Use short sentences, bullets, or a small sketch. We assess the reasoning, not English grammar.

## Interview Practice

**Active level: Small.** Start here. The mentor activates a harder level only after the learner can explain the simple case and an edge case, use a familiar example without step-by-step help, and says they are comfortable. Task number and time spent do not automatically raise the interview level.

### Small - active

```mermaid
flowchart LR
  N0["Stored rows"]
  N1["LIMIT ten"]
  N2["Stable page?"]
  N0 --> N1
  N1 --> N2
```

Read left to right: Stored rows -> LIMIT ten -> Stable page?.

A query returns ten rows without ORDER BY. Can a client rely on the same ten next time? State the promise you would need before calling it a page.

<details>
<summary>Moderate - later practice, not active yet</summary>

```mermaid
flowchart LR
  N0["Learning rows"]
  N1["Two join choices"]
  N2["Rows and counts"]
  N0 --> N1
  N1 --> N2
```

Read left to right: Learning rows -> Two join choices -> Rows and counts.

One learning record has two retry rows and another has none. Predict how INNER JOIN and LEFT JOIN affect the rows and counts.

New concepts in this preview must be explained before this level is assigned.

</details>

<details>
<summary>Strong - future practice, not active yet</summary>

```mermaid
flowchart LR
  N0["Concurrent inserts"]
  N1["Page boundary"]
  N2["Index choice"]
  N3["Measured plan"]
  N0 --> N1
  N1 --> N2
  N2 --> N3
```

Read left to right: Concurrent inserts -> Page boundary -> Index choice -> Measured plan.

Later, data grows and inserts happen between page requests. Compare offset paging and a stable cursor, then describe how EXPLAIN evidence would guide an index choice.

The mentor introduces any new concepts first. This is a preview, not a requirement to implement another system now.

</details>

## Research Hints

Search `PostgreSQL joins NULL count`, `WHERE GROUP BY HAVING`, `stable keyset pagination`, and `PostgreSQL EXPLAIN ANALYZE`.

Look up the concept, inspect the relevant project code, then make your own implementation decision. Hints are not a complete solution.

## Scope And Coaching Agreement

Use local learning data only. PostgreSQL stores operational/request metadata. Keep financial calculations, ledger rules, processor calls, and real payment lifecycle decisions outside this assignment.

Reviewers own routine tests and documentation. Implementation review checks the listed behavior. Questions and interview answers belong in the related PR description or comment, in the learner's own words. Preserve original answers and add focused corrections. Missing, partial, or incorrect learning answers are non-blocking for an otherwise correct implementation. No extra code, tests, or documentation are needed to answer.

