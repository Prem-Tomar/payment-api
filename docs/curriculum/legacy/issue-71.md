# Original Issue #71

Historical snapshot before inventory consolidation. This file preserves the original request; it is not completion evidence.

Title: Day G58: Bound payment intent path identifier length

## Goal

Bound payment-intent path identifier length consistently for lookup and confirmation.

## Context

A path identifier is the resource name supplied in the URL. The API currently trims it but has no explicit length rule. Boundary validation should reject unusable input before application work.

## Task

Choose and name a small maximum length for the learning API's payment-intent path identifier. Apply it consistently to lookup and confirmation after trimming, preserving the current behavior for ordinary identifiers.

## Completion Criteria

- Both routes reject empty and overlong identifiers with a clear validation response.
- Acceptable identifiers reach the existing route behavior.
- The maximum and validation responsibility are easy to find.
- The code does not infer ownership or existence from the identifier's spelling.

## Scope

Implementation only. Tests and documentation are reviewer-owned. No real ID generator, database lookup redesign, or production identifier standard.

## Research Hints

Search `Gin path parameter validation`, `Go string length bytes runes`, and `API identifier length limit`. Decide whether the limit counts bytes or characters and keep the rule consistent.

## Question of the Day - G58

Previously learned concepts: Lookup and confirmation route shells, path-parameter parsing, and trimming.

Scenario: Lookup accepts an identifier that confirmation rejects even though both refer to the same resource name.

1. Why is that inconsistent for a caller? Where should a shared input rule be applied, and what must be specified if a length rule can count either bytes or characters?
2. Explain what this example teaches about responsibility or failure boundaries in the system, and relate it to one earlier implementation.

System-design connection: A path identifier is the resource name in the URL. Earlier parsing establishes its boundary; later handlers should share its contract without inferring ownership or existence from its shape.

Answer every part in your own words directly in the relevant PR description or comment; an issue-only answer does not satisfy this placement rule. Explain your assumptions and reasoning. No extra implementation, test-writing, or documentation is required to answer. The reviewer preserves your original answer and marks it Correct, Partially correct, Needs correction, or Awaiting answer, with targeted corrections and follow-ups. This learning response is expected but does not block merging an otherwise correct implementation.


