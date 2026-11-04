# G54/G56/G62: Complete the request-input path

## Goal

Make accepted HTTP input reach the application without lost fields or changed numbers.

## Picture First

```mermaid
flowchart LR
  N0["JSON body"]
  N1["Parse one object"]
  N2["Clean input"]
  N3["Application input"]
  N0 --> N1
  N1 --> N2
  N2 --> N3
```

Read left to right: JSON body -> Parse one object -> Clean input -> Application input.

## Combined Inventory

This is one delivery assignment, not several separate PR tasks. Original goals are grouped below. A closed source issue means superseded, not achieved. Complete and review each listed goal before claiming it is done.

- [Original #64](https://github.com/Prem-Tomar/payment-api/issues/64): Day G54: Reject trailing JSON in create payment intent requests. [Saved original](../legacy/issue-64.md).
- [Original #66](https://github.com/Prem-Tomar/payment-api/issues/66): Day G56: Preserve currency and payment method in application input. [Saved original](../legacy/issue-66.md).
- [Original #75](https://github.com/Prem-Tomar/payment-api/issues/75): Day G62: Preserve amount width across the application boundary. [Saved original](../legacy/issue-75.md).

## Implementation Parts

1. Reject a second JSON value or trailing junk; allow trailing spaces (#64).
2. Pass currency and payment method with the accepted request fields (#66).
3. Keep the amount width unchanged between the DTO and application input (#75).

## Completion Criteria

- Bad input stops before the use case.
- Accepted amount, currency, method, and references reach the use case unchanged except intentional cleanup.
- Existing body-size and error-code rules still apply.
- Show the connected success path and at least one relevant failure path in the PR discussion. A short explanation or demonstration is enough; no learner-written test suite or document is required.

## What, Why, And Where

- What: Make accepted HTTP input reach the application without lost fields or changed numbers.
- Why: the components must keep the same promise when they are connected.
- Where: follow the picture above through the existing project files.
- Earlier learning: Builds on JSON binding, body limits, reference cleanup, and the application interface.
- Later use: Later SQL columns must receive the same accepted values. SQL parameters keep values separate from query text.

Words used here: DTO: the struct used to read or write HTTP data. Normalization: changing accepted input to one agreed form.

## Software And Design Principle

Single Responsibility Principle: parsing, cleanup, and use-case work have different reasons to change.

Use a pattern only when it solves a visible problem. A Go function or small interface may be enough. [Short principle guide](../../software-design-principles.md).

## Database Learning Connection

```mermaid
flowchart LR
  N0["Accepted input"]
  N1["SQL parameters later"]
  N2["Typed columns"]
  N3["Stored row"]
  N0 --> N1
  N1 --> N2
  N2 --> N3
```

Read left to right: Accepted input -> SQL parameters later -> Typed columns -> Stored row.

This picture is a future preview. Explain the connection if useful, but do not implement SQL as an extra requirement of this assignment.
[In-depth SQL guide](../../sql-learning-guide.md).

## Question of the Day

Previously learned concepts: Builds on JSON binding, body limits, reference cleanup, and the application interface.

```mermaid
flowchart LR
  N0["Small body"]
  N1["First object valid"]
  N2["Second object"]
  N3["Accept or reject?"]
  N0 --> N1
  N1 --> N2
  N2 --> N3
```

Read left to right: Small body -> First object valid -> Second object -> Accept or reject?.

A body is under the size limit. It has one valid object and then a second object. Does the size check prove the whole body is valid? Explain the separate jobs of size checking and parsing.

Explain in your own words. Use short sentences, bullets, or a small sketch. We assess the reasoning, not English grammar.

## Interview Practice

**Active level: Small.** Start here. The mentor activates a harder level only after the learner can explain the simple case and an edge case, use a familiar example without step-by-step help, and says they are comfortable. Task number and time spent do not automatically raise the interview level.

### Small - active

```mermaid
flowchart LR
  N0["Positive input"]
  N1["Type conversion"]
  N2["Use-case input"]
  N0 --> N1
  N1 --> N2
```

Read left to right: Positive input -> Type conversion -> Use-case input.

A handler accepts a positive int64 amount, then converts it to a smaller integer type. Point to the place where information can change. What must the application input promise?

<details>
<summary>Moderate - later practice, not active yet</summary>

```mermaid
flowchart LR
  N0["Route A"]
  N1["Cleanup rule"]
  N2["Use case"]
  N3["Route B"]
  N0 --> N1
  N1 --> N2
  N2 --> N3
```

Read left to right: Route A -> Cleanup rule -> Use case -> Route B.

Two routes clean the same input in different ways. Design one clear owner for cleanup without making the application depend on Gin.

New concepts in this preview must be explained before this level is assigned.

</details>

<details>
<summary>Strong - future practice, not active yet</summary>

```mermaid
flowchart LR
  N0["Old client"]
  N1["Versioned API"]
  N2["Input mapping"]
  N3["New client"]
  N0 --> N1
  N1 --> N2
  N2 --> N3
```

Read left to right: Old client -> Versioned API -> Input mapping -> New client.

Later, an old client and a new client send different field sets. Plan a compatible input contract. Explain what you would reject and what you would preserve.

The mentor introduces any new concepts first. This is a preview, not a requirement to implement another system now.

</details>

## Research Hints

Search `Go json Decoder trailing data`, `Go int int64 conversion`, and `DTO application input mapping`.

Look up the concept, inspect the relevant project code, then make your own implementation decision. Hints are not a complete solution.

## Scope And Coaching Agreement

This grouping keeps the original Go requirements. Database content below is coaching only; no SQL or database code is added to these legacy goals. Keep existing exclusions. Do not add real money movement, processor calls, or financial-domain logic.

Reviewers own routine tests and documentation. Implementation review checks the listed behavior. Questions and interview answers belong in the related PR description or comment, in the learner's own words. Preserve original answers and add focused corrections. Missing, partial, or incorrect learning answers are non-blocking for an otherwise correct implementation. No extra code, tests, or documentation are needed to answer.

