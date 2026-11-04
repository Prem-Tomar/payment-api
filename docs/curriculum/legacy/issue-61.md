# Original Issue #61

Historical snapshot before inventory consolidation. This file preserves the original request; it is not completion evidence.

Title: Day G53: Add fake payment intent ID sequence boundary

Goal: replace the single hard-coded fake payment intent ID with a tiny fake ID generator boundary that still has no real storage or payment behavior. Learning focus: isolating generated values behind small interfaces. Implementation task: introduce a minimal ID generator dependency for the fake implementation and keep the generated ID visibly fake. Rules: no documentation task, no test requirement, no database, no production store, no processor client, no ledger behavior, no Rust integration. Completion criteria: the fake implementation no longer owns the literal ID inline, generated IDs are clearly placeholder values, and replacing the generator later would be straightforward. Research hints: search for Go interface for ID generator, deterministic fake IDs, and dependency injection without framework; compare a simple function dependency with a struct interface.

## Question of the Day - G53

Previously learned concepts: Fake application implementations, in-memory records, and lookup behavior.

Scenario: A fake create action returns the same ID for every request. Later learning code starts using that ID as a map key.

1. When is the fixed value useful, and when can it hide collisions or overwrite records? How should a caller recognize placeholder behavior before assuming a real resource exists?
2. Explain what this example teaches about responsibility or failure boundaries in the system, and relate it to one earlier implementation.

System-design connection: A fake is a controlled substitute used during learning. Earlier boundaries isolate it; later ID generation can vary identifiers without teaching that a fake performs real payment work.

Answer every part in your own words directly in the relevant PR description or comment; an issue-only answer does not satisfy this placement rule. Explain your assumptions and reasoning. No extra implementation, test-writing, or documentation is required to answer. The reviewer preserves your original answer and marks it Correct, Partially correct, Needs correction, or Awaiting answer, with targeted corrections and follow-ups. This learning response is expected but does not block merging an otherwise correct implementation.


