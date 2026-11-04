# SEQ-02 | G44/G48/G49/G50: Carry input and request lifetime through the use case

## Goal

Make application calls explicit and make each request produce one final outcome.

## Prerequisites

The existing service/store foundations are enough. No unfinished assignment is required.

## Picture First

```mermaid
flowchart LR
  N0["Body and metadata"]
  N1["Application input"]
  N2["Context-aware use case"]
  N3["One response"]
  N0 --> N1
  N1 --> N2
  N2 --> N3
```

Read left to right: Body and metadata -> Application input -> Context-aware use case -> One response.

## Combined Inventory

Complete the related parts below together. Source requests and saved originals are linked for reference.

- [Original #55](https://github.com/Prem-Tomar/payment-api/issues/55): Day G44: Add request cancellation handling in create use case. [Saved original](../legacy/issue-55.md).
- [Original #56](https://github.com/Prem-Tomar/payment-api/issues/56): Day G48: Add request metadata input to create use case. [Saved original](../legacy/issue-56.md).
- [Original #58](https://github.com/Prem-Tomar/payment-api/issues/58): Day G49: Add timeout error mapping at application boundary. [Saved original](../legacy/issue-58.md).
- [Original #59](https://github.com/Prem-Tomar/payment-api/issues/59): Day G50: Add create payment intent command boundary. [Saved original](../legacy/issue-59.md).

## Implementation Parts

1. Observe cancellation before returning a create result (#55).
2. Pass request metadata separately from the JSON DTO (#56).
3. Map cancellation and deadline errors to distinct, single responses (#58).
4. Use one small application command/input shape (#59).

## Completion Criteria

- The use case receives the intended input, metadata, and request context.
- Direct and wrapped deadline errors produce one JSON error, with no trailing object.
- Canceled, invalid, not-implemented, and unexpected failures remain distinct.
- Show the connected success path and at least one relevant failure path in the PR discussion. A short explanation or demonstration is enough; no learner-written test suite or document is required.

## Scope

This grouping keeps the original Go requirements. Database content below is coaching only; no SQL or database code is added to these legacy goals. Keep existing exclusions. Do not add real money movement, processor calls, or financial-domain logic.

Routine tests and documentation are reviewer-owned.

<details>
<summary>Learning discussion - not extra implementation requirements</summary>

## What, Why, And Where

- What: Make application calls explicit and make each request produce one final outcome.
- Why: the components must keep the same promise when they are connected.
- Where: follow the picture above through the existing project files.
- Earlier learning: Builds on shutdown cleanup, request timeout middleware, metadata, and the application interface.
- Later use: Database calls can later use the same context. A pool wait or slow query must also obey the request lifetime.

Words used here: Context: Go's request-lifetime signal. Deadline: the time when work must stop. Command input: the data for one action; it is not automatically the full Command design pattern.

## Software And Design Principle

Dependency Inversion: application work depends on a small contract, not Gin. Keep ownership of the final response clear.

Use a pattern only when it solves a visible problem. A Go function or small interface may be enough. [Short principle guide](../../software-design-principles.md).

## Database Learning Connection

```mermaid
flowchart LR
  N0["Request context"]
  N1["SQL call later"]
  N2["Wait or query"]
  N3["Stop signal"]
  N0 --> N1
  N1 --> N2
  N2 --> N3
```

Read left to right: Request context -> SQL call later -> Wait or query -> Stop signal.

This picture is a future preview. Explain the connection if useful, but do not implement SQL as an extra requirement of this assignment.
[In-depth SQL guide](../../sql-learning-guide.md).

## Question of the Day

Previously learned concepts: Builds on shutdown cleanup, request timeout middleware, metadata, and the application interface.

```mermaid
flowchart LR
  N0["Deadline ends"]
  N1["Context signal"]
  N2["Function checks signal"]
  N0 --> N1
  N1 --> N2
```

Read left to right: Deadline ends -> Context signal -> Function checks signal.

A context reaches its deadline. Does passing it to a function automatically stop that function? Explain passing the signal versus observing it.

Use your own words. Any programming language, pseudocode, or a sketch is welcome. If the question concerns Go or SQL semantics, explain that original behavior correctly.

## Interview Problem

```mermaid
flowchart LR
  N0["Deadline error"]
  N1["First JSON"]
  N2["Second JSON?"]
  N0 --> N1
  N1 --> N2
```

Read left to right: Deadline error -> First JSON -> Second JSON?.

The handler writes a timeout error and then runs its generic error branch. Trace the two writes. Where must the response path end?

## Research Hints

Search `Go errors.Is context deadline`, `Gin response written return`, and `application service input struct`.

Look up the concept, inspect the relevant project code, then make your own implementation decision. Hints are not a complete solution.

For discussion, any programming language, pseudocode, or diagram is accepted. QOD revisits earlier work; the interview examples explore different scenarios. [How learning answers are reviewed](../../learning-and-interview-guide.md).

</details>

