# Original Issue #65

Historical snapshot before inventory consolidation. This file preserves the original request; it is not completion evidence.

Title: Day G55: Reject duplicate payment intent IDs in the memory store

## Goal

Prevent the learning in-memory store from silently replacing an existing payment-intent record during creation.

## Context

A duplicate ID means a record already uses the identifier being saved. A conflict error tells the application that creation cannot use that ID. A map assignment currently replaces the old entry without reporting it.

## Implementation Task

Give the store's creation/save operation a clear duplicate-ID rule and return a small application-level conflict error when the ID already exists. Keep the existence check and insert protected together so concurrent requests cannot both create the same ID.

## Completion Criteria

- Saving a new ID still succeeds.
- Saving the same ID again reports a recognizable application error and preserves the original record.
- The duplicate check and insert cannot race when requests run concurrently.
- The application depends on the store interface; existing HTTP routes remain stable until a later wiring task.

## Scope

Implementation only. Tests and documentation are handled by the reviewer, not required from the learner. No database, automatic retries, ID-generator redesign, payment processor, ledger behavior, or Rust integration.

## Research Hints

Search for `Go sentinel error errors.Is`, `Go map check insert mutex`, and `repository create conflict`. A sentinel error is a named error value callers can recognize. Compare create-only behavior with an operation that updates an existing record.

## Question of the Day - G55

Previously learned concepts: Shared maps, mutex protection, and the in-memory store interface.

Scenario: Two concurrent callers check whether the same record ID exists before either takes the lock used for insertion.

1. Can both conclude the ID is available even if each individual map write is locked? Explain the difference between avoiding a data race and enforcing one create decision.
2. Explain what this example teaches about responsibility or failure boundaries in the system, and relate it to one earlier implementation.

System-design connection: A mutex allows one protected operation at a time. Earlier locking protects shared memory; later create-only storage rules need the whole check-and-insert decision to be coordinated.

Answer every part in your own words directly in the relevant PR description or comment; an issue-only answer does not satisfy this placement rule. Explain your assumptions and reasoning. No extra implementation, test-writing, or documentation is required to answer. The reviewer preserves your original answer and marks it Correct, Partially correct, Needs correction, or Awaiting answer, with targeted corrections and follow-ups. This learning response is expected but does not block merging an otherwise correct implementation.


