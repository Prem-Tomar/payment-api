# SEQ-04 | G15/G39/G51/G52: Make application results and API output agree

## Goal

Make the public response describe what the application really did.

## Learn Before Building

### 1. Results and public representations

An application result describes what the use case actually did. A response DTO is the public JSON shape used to describe that result to a client. A mapper translates between these shapes. An API contract is the promise clients rely on, including status codes, field names, and meanings. A placeholder is intentionally unfinished learning behavior. An idempotency key labels the caller's logical operation for retry handling. Requiring its header at this boundary does not yet implement a replay engine.

```mermaid
flowchart LR
  A["Actual application outcome"] --> B["Application result"] --> C["Response mapper"] --> D["HTTP status and JSON"]
```

**Worked example:** A preview calculation can return a preview result without claiming that a permanent resource was stored. Conversely, a response saying a new resource was created must have an implementation behind that promise. HTTP 201 means creation succeeded; an attractive response body cannot make unsupported work real.

**Common mistake:** Choosing a success label in the handler independently of the application result can make the status, ID, and body disagree.

**Why and where:** Single Responsibility keeps application decisions separate from public formatting. This extends typed responses and prepares for consistent stored-record lookup.

### 2. State names and allowed transitions

A state is a named stage in a resource's life. A typed string gives that name a distinct Go type. A transition is a change from one state to another under allowed rules. A state machine defines those rules. The State behavioral pattern can place different behavior behind different states; merely adding a string type is not that pattern.

```mermaid
flowchart LR
  A["Result has a stage"] --> B["Agreed public vocabulary"] --> C["Mapper preserves meaning"]
  D["Future domain rules"] --> E["Allow or reject transitions"]
```

**Worked example:** A support ticket could use open and closed. A named string can still receive an explicitly converted unknown value unless code validates it. Spelling is part of a public contract: two similar words are not automatically interchangeable for a client. The payment vocabulary here includes requires_confirmation.

**Common mistake:** Adding an HTTP-side state machine would take ownership of financial rules that belong to the domain.

**Why and where:** Precise vocabulary helps clients and logs describe the same outcome. Future behavioral design discussions can deepen the idea without adding financial behavior now.

### 3. Database preview: a row and an outcome are different

A database row is a stored record. Persistence means accepted data can survive the API process ending. A store adapter connects a database tool to an application contract. An operational learning record is not proof that a real payment was charged.

```mermaid
flowchart LR
  A["Application outcome"] --> B["Optional stored learning metadata"] --> C["Later scoped lookup"]
```

**Worked example:** Storing a request timestamp proves that metadata was saved. It does not prove a processor accepted money movement. The response must describe the supported operation.

**Common mistake:** Using persistence language to imply a real payment lifecycle that has not been implemented.

**Why and where:** This connects truthful output to later SQL storage. No database implementation is added to this existing Go assignment.


## Prerequisites

- SEQ-02: [issue #59](https://github.com/Prem-Tomar/payment-api/issues/59).

## Picture First

```mermaid
flowchart LR
  N0["Use case result"]
  N1["Allowed state"]
  N2["Response mapper"]
  N3["Client sees outcome"]
  N0 --> N1
  N1 --> N2
  N2 --> N3
```

Read left to right: Use case result -> Allowed state -> Response mapper -> Client sees outcome.

## Combined Inventory

Complete the related parts below together. Source requests and saved originals are linked for reference.

- [Original #20](https://github.com/Prem-Tomar/payment-api/issues/20): Day G15: Add first payment intent route shell. [Saved original](../legacy/issue-20.md).
- [Original #46](https://github.com/Prem-Tomar/payment-api/issues/46): Day G39: Add payment intent lifecycle state model. [Saved original](../legacy/issue-46.md).
- [Original #60](https://github.com/Prem-Tomar/payment-api/issues/60): Day G51: Add application result status boundary. [Saved original](../legacy/issue-60.md).
- [Original #62](https://github.com/Prem-Tomar/payment-api/issues/62): Day G52: Add payment intent response mapper boundary. [Saved original](../legacy/issue-62.md).

## Implementation Parts

1. Keep placeholder creation explicit; do not claim a resource exists when none is created (#20).
2. Use the intended lifecycle vocabulary, including requires_confirmation (#46).
3. Let the application result own its status (#60).
4. Map that result through a small response mapper (#62).

## Completion Criteria

- Status code, body, ID, and placeholder meaning agree.
- The application result and HTTP DTO remain separate. The application owns its deliberate typed state and allowed vocabulary; the HTTP mapper stays small and close to the API layer.
- Preserve the versioned create route, request parsing, boundary validation, and required idempotency-key check before the honest placeholder response.
- No real processor, lifecycle engine, or payment behavior is added.
- Show the connected success path and at least one relevant failure path in the PR discussion. A short explanation or demonstration is enough; no learner-written test suite or document is required.

## Scope

This grouping keeps the original Go requirements. Database content in this assignment is coaching only; no SQL or database code is added to these legacy goals. Keep existing exclusions. Do not add real money movement, processor calls, or financial-domain logic.

Routine tests and documentation are reviewer-owned.

<details>
<summary>Learning discussion - not extra implementation requirements</summary>

## What, Why, And Where

- What: Make the public response describe what the application really did.
- Why: the components must keep the same promise when they are connected.
- Where: follow the picture above through the existing project files.
- Earlier learning: Builds on typed response values, fake services, the create port, and lookup.
- Later use: A later stored learning record can be read after creation. Database storage will not decide financial lifecycle rules.


## Software And Design Principle

Single Responsibility; State is a behavioral-pattern preview. A typed state alone is not a full state machine.

Use a pattern only when it solves a visible problem. A Go function or small interface may be enough. [Short principle guide](../../software-design-principles.md).

## Database Learning Connection

```mermaid
flowchart LR
  N0["Learning result"]
  N1["SQL record later"]
  N2["Lookup"]
  N3["Honest response"]
  N0 --> N1
  N1 --> N2
  N2 --> N3
```

Read left to right: Learning result -> SQL record later -> Lookup -> Honest response.

This picture is a future preview. Explain the connection if useful, but do not implement SQL as an extra requirement of this assignment.
[In-depth SQL guide](../../sql-learning-guide.md).

## Question of the Day

Previously learned concepts: Builds on typed response values, fake services, the create port, and lookup.

```mermaid
flowchart LR
  N0["Named string type"]
  N1["Assigned value"]
  N2["Public response"]
  N0 --> N1
  N1 --> N2
```

Read left to right: Named string type -> Assigned value -> Public response.

A named string type is used for a resource state. Does that type alone guarantee that every value is allowed? Explain using the earlier typed-status work.

Use your own words. Any programming language, pseudocode, or a sketch is welcome. If the question concerns Go or SQL semantics, explain that original behavior correctly.

## Interview Problem

```mermaid
flowchart LR
  N0["POST result"]
  N1["Client expectation"]
  N2["GET result"]
  N0 --> N1
  N1 --> N2
```

Read left to right: POST result -> Client expectation -> GET result.

POST returns created, but GET cannot find the ID. What promise did the client hear? Show a truthful placeholder outcome without adding storage.

## Research Hints

Search `HTTP 201 resource semantics`, `Go response mapper`, and `typed state versus state machine`.

Look up the concept, inspect the relevant project code, then make your own implementation decision. Hints are not a complete solution.

Reply in your own words in the PR discussion. Any programming language, pseudocode, or diagram is welcome. [Answer guidance](../../learning-and-interview-guide.md).

</details>

