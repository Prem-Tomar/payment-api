# Original Issue #75

Historical snapshot before inventory consolidation. This file preserves the original request; it is not completion evidence.

Title: Day G62: Preserve amount width across the application boundary

## Goal

Preserve the accepted integer amount when mapping a request into the create use case.

## Context And System Design

The DTO uses int64 while the application currently uses int. The width of int depends on the target architecture; a narrowing conversion can change a value before the application sees it.

Earlier boundaries separated transport from application data. Preserving representation belongs at their mapping point; later domain calls need that same information without changing numeric meaning.

## Implementation Task

Use a consistent amount representation across request mapping and application input so accepted values retain their meaning. Keep real financial arithmetic outside this task.

## Completion Criteria

- An accepted request amount reaches the application unchanged.
- The input shape does not depend on 32-bit versus 64-bit int width.
- Validation and honest placeholder behavior remain separate responsibilities.

## Scope

Implementation only. Tests and documentation are reviewer-owned. No real processor, database, ledger behavior, financial calculations, production observability system, or general framework is required.

## Research Hints

Search `Go int int64 architecture conversion`, `integer narrowing Go`, and `DTO application input mapping`.

## Question of the Day - G62

Previously learned concepts: Integer request fields, DTO validation, and application input mapping.

Scenario: A positive int64 value of 2147483648 is converted to int on a 32-bit target after passing a greater-than-zero check.

1. Could the earlier validation still guarantee the converted value is positive? Trace where information can change, and explain why validating input differs from preserving its representation.
2. Explain the system-design responsibility involved and connect it to earlier implemented work. Future uses in the context above are previews, not assumed knowledge.

Answer every part in your own words directly in the relevant PR description or comment. An issue-only answer does not meet the placement rule. No extra implementation, tests, or documentation are required to answer. The reviewer preserves the answer, identifies correct and incomplete parts, explains corrections, and asks targeted follow-ups. The learning response is expected but does not block merging an otherwise correct implementation.


