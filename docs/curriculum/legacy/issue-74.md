# Original Issue #74

Historical snapshot before inventory consolidation. This file preserves the original request; it is not completion evidence.

Title: Day G61: Identify the invalid field in API validation responses

## Goal

Let callers identify the rejected request field without parsing a human-readable message.

## Context And System Design

A field name identifies the input control to correct. Typed codes already distinguish broad failures; field metadata can make the response useful when multiple references share one code.

Earlier validation mapping established a stable error contract. Field metadata belongs at that HTTP boundary; later client forms can use it to guide correction.

## Implementation Task

Add a small optional public field identifier to validation-error responses and populate it through the existing mapper. Keep JSON field names separate from internal Go struct names.

## Completion Criteria

- Relevant validation responses identify the rejected public field.
- Malformed JSON and unexpected server errors do not invent a field.
- Codes, statuses, and existing success behavior remain stable.

## Scope

Implementation only. Tests and documentation are reviewer-owned. No real processor, database, ledger behavior, financial calculations, production observability system, or general framework is required.

## Research Hints

Search `Gin validator FieldError JSON field name`, `API validation error field metadata`, and `Go json omitempty`.

## Question of the Day - G61

Previously learned concepts: Binding validation, field-specific codes, and shared response helpers.

Scenario: Both merchant_ref and customer_ref can return the same INVALID_REFERENCE code.

1. How would a client know which input to highlight? Why should a response use a public JSON field name rather than require callers to understand the Go struct?
2. Explain the system-design responsibility involved and connect it to earlier implemented work. Future uses in the context above are previews, not assumed knowledge.

Answer every part in your own words directly in the relevant PR description or comment. An issue-only answer does not meet the placement rule. No extra implementation, tests, or documentation are required to answer. The reviewer preserves the answer, identifies correct and incomplete parts, explains corrections, and asks targeted follow-ups. The learning response is expected but does not block merging an otherwise correct implementation.


