# SEQ-07 | G54/G56/G62: Complete the request-input path

## Goal

Make accepted HTTP input reach the application without lost fields or changed numbers.

## Learn Before Building

### 1. Body size and JSON shape are separate checks

A request body is the bytes sent after HTTP headers. JSON is a format for representing structured data. A decoder reads those bytes into Go values. Parsing checks whether the input follows the format. A byte limit bounds how much input may be consumed. EOF means the reader reached the end of the input.

```mermaid
flowchart LR
  B["Incoming body bytes"] --> L["Bound byte count"] --> D["Decode the intended value"] --> E["Check whether input is finished"]
```

**Worked example:** A catalogue endpoint expects one object. A body containing that object followed by spaces still represents one value. A decoder can finish one value before it reaches the end of the reader. Knowing that the first value decoded does not describe the remaining bytes.

**Common mistake:** Assuming Content-Length, the declared body length header, always exists or is trustworthy. A reader must enforce its own consumption limit. Streaming means reading incrementally rather than requiring the whole body in memory.

**Why and where:** This extends the existing body-size guard and JSON parsing. The size guard controls resource use; parsing controls accepted meaning. Keep their error outcomes distinct.

### 2. Mapping must preserve accepted information

A DTO is the transport data shape. Application input is the shape passed to the use case. Mapping copies accepted values between those boundaries. Integer width is how many bits a type uses; its representable range determines which amounts fit. int64 is a signed 64-bit integer. Go int can be 32 or 64 bits depending on the target architecture, so it is not an architecture-independent substitute. A narrower type can represent fewer integers.

```mermaid
flowchart LR
  A["Accepted DTO fields"] --> B["Deliberate mapping"] --> C["Application input"] --> D["Use case receives same meaning"]
```

**Worked trace:** A toy three-digit storage field cannot represent a four-digit accepted value. The equivalent software risk appears when a wide accepted number is converted into a smaller integer type. Currency and payment method can also be lost if a mapper copies only amount.

**Common mistake:** Assuming compilation proves that a numeric conversion preserves every value, or normalizing a field in the DTO but forwarding an older copy.

**Why and where:** Single Responsibility means parsing, intentional cleanup, and application work have different reasons to change. Preserve accepted information through the complete path rather than adding policy to a mapper.

### 3. Database preview: typed values stay data

A SQL parameter is a value passed separately from a database command. A column type bounds what a database field can hold. A store adapter translates the application contract into storage operations.

```mermaid
flowchart LR
  A["Accepted value"] --> B["Application input"] --> C["Future SQL parameter"] --> D["Compatible column"]
```

**Worked example:** A future inventory count accepted as a wide integer still needs a database column that can represent it. A safe parameter protects the SQL command's structure, but does not repair an earlier narrowing conversion.

**Common mistake:** Using database validation as a reason to silently change accepted values before they reach storage.

**Why and where:** This extends today's end-to-end field mapping. Database code remains outside this legacy Go assignment.


## Prerequisites

- SEQ-02: [issue #59](https://github.com/Prem-Tomar/payment-api/issues/59).

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

Complete the related parts below together. Source requests and saved originals are linked for reference.

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
- Existing body-size and error-code rules still apply: oversized input and field validation keep their existing categories.
- The HTTP DTO and application input remain separate; the Gin handler is the explicit mapping point.
- Amount representation is independent of 32-bit versus 64-bit architecture.
- Show the connected success path and at least one relevant failure path in the PR discussion. A short explanation or demonstration is enough; no learner-written test suite or document is required.

## Scope

This grouping keeps the original Go requirements. Database content in this assignment is coaching only; no SQL or database code is added to these legacy goals. Keep existing exclusions. Do not add real money movement, processor calls, or financial-domain logic.

Routine tests and documentation are reviewer-owned.

<details>
<summary>Learning discussion - not extra implementation requirements</summary>

## What, Why, And Where

- What: Make accepted HTTP input reach the application without lost fields or changed numbers.
- Why: the components must keep the same promise when they are connected.
- Where: follow the picture above through the existing project files.
- Earlier learning: Builds on JSON binding, body limits, reference cleanup, and the application interface.
- Later use: Later SQL columns must receive the same accepted values. SQL parameters keep values separate from query text.


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

Use your own words. Any programming language, pseudocode, or a sketch is welcome. If the question concerns Go or SQL semantics, explain that original behavior correctly.

## Interview Problem

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

## Research Hints

Search `Go json Decoder trailing data`, `Go int int64 conversion`, and `DTO application input mapping`.

Look up the concept, inspect the relevant project code, then make your own implementation decision. Hints are not a complete solution.

Reply in your own words in the PR discussion. Any programming language, pseudocode, or diagram is welcome. [Answer guidance](../../learning-and-interview-guide.md).

</details>

