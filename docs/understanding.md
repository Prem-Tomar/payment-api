# payment-api Understanding

## Goal

Define the Go service track clearly before writing service code.

The `payment-api` service will be the high-volume REST boundary for payment operations. It should
be fast, operationally simple, and easy to deploy, while relying on Rust crates for financial
correctness and heavy deterministic processing.

## Real Problem

Public payment APIs receive messy external input: JSON payloads, headers, idempotency keys,
authentication context, client mistakes, retries, timeouts, and duplicate requests. That boundary
must be robust, but it should not become the source of financial truth.

If Go duplicates money arithmetic, fee formulas, lifecycle rules, or ledger logic, the system can
drift into inconsistent behavior. The API should validate transport-level shape and then delegate
domain decisions to the Rust-owned contracts.

## Architecture Boundary

Go should handle:

- parsing HTTP requests.
- validating required request fields.
- rejecting malformed JSON and unsupported media types.
- normalizing simple transport data such as headers and request IDs.
- applying request-size limits.
- coordinating idempotency checks when persistence is introduced.
- mapping domain errors to stable HTTP responses.
- exposing health, readiness, metrics, and later tracing.

Rust should handle:

- exact currency and money representation.
- payment amount invariants.
- fee calculation and fee evidence.
- payment lifecycle transition rules.
- deterministic ledger behavior when that crate exists.
- memory-intensive or correctness-critical processing.

## Coaching Style For This Repo

Use the same learning style as the Rust track:

- Start with the goal and real systems problem.
- Explain the boundary before writing code.
- Keep each task small enough for daily progress.
- Review architecture/domain alignment separately from Go code quality.
- Teach the concept behind every review issue.
- Avoid expanding into new behavior without explicit scope.

## Question of the Day

Starting with G40, every new issue includes one focused question phrased as an interviewer asking
the developer to explain a previously covered topic. This is a recall and reasoning exercise,
not another implementation requirement. Existing issues are updated when requested.

### Choosing the Question

- Check earlier issues, implementations, and review discussions to confirm the topic was covered.
- Name the earlier topic and link its issue. Use a concrete scenario and ask the developer to
  explain why something works, a tradeoff, or a failure case in their own words.
- Rotate across all covered topics over successive issues: Go service boundaries and packages;
  HTTP routes and responses; middleware order, request IDs and logging; configuration, timeouts
  and shutdown; health and readiness; request parsing, validation and body limits; idempotency;
  API-key and merchant boundaries; rate limiting; application interfaces, stores and state models.
  Add later topics only after they have been covered. One issue does not need to cover everything.
- Inspect recent questions before choosing the next one. Revisit a weak answer later with a
  different scenario. Do not paste a model answer into the question or ask for out-of-scope code.
- Aim for staff/principal-level reasoning through depth, not obscure trivia or unlearned technology.
  Start with a commonly missed mechanism, then ask the developer to trace its consequences under
  failure, concurrency, cancellation, or retries as relevant to that topic. Follow through to
  ownership, tradeoffs, caller-visible behavior, observability, and a way to verify the design.
  Use two or three follow-up prompts within the same scenario rather than a disconnected quiz.
  Distinguish language guarantees, assumptions, and design choices; accept well-reasoned alternatives.
  Examples such as `defer` timing or panic recovery illustrate the expected depth, not a fixed
  syllabus. Apply this questioning standard across every previously covered topic.
- Use `.github/ISSUE_TEMPLATE/learning_task.md` when drafting issues. Include the same sections
  when using an API or other tool that bypasses GitHub templates.

### Answering and Reviewing

The developer must provide an answer in their own words in an issue comment or the associated
PR's **Question of the day** section. Include the question in the PR too; a link to the issue
answer is sufficient and avoids answering twice. Explain the reasoning and give a small example
where useful. This does not add a learner-written test or repository documentation deliverable.

During review, record the question/answer link and one learning result: **Correct**, **Partially
correct**, **Needs correction**, or **Awaiting answer**. Explain what is right, what needs changing,
and why. Where needed, add a corrected explanation under **Reviewer correction** beside the
original answer or in a linked reply. Preserve the developer's original wording and distinguish
reviewer feedback from it; never silently replace it or write an answer on the developer's behalf.
Assess both the mechanism and the engineering reasoning: assumptions, failure cases, tradeoffs,
observable consequences, and verification. Do not grade only terminology or predicted output.

Record the implementation verdict separately. A missing, incomplete, or incorrect answer, or a
pending answer review, must not block a PR whose implementation is correct and passes the normal
implementation checks. Leave learning follow-up in the issue or PR, including after merge when
needed. Do not request changes, fail CI, or withhold implementation approval solely for the answer.
If the discussion reveals an actual code defect, assess that defect under the normal code-review
rules and explain the concrete implementation impact separately.

### First Question: G40

Previously covered topic: Go `defer`, used for shutdown/context cleanup in
[G2](https://github.com/Prem-Tomar/payment-api/issues/4), and final-outcome logging in
[G5](https://github.com/Prem-Tomar/payment-api/issues/7).

**Interviewer asks:** A teammate says, "A deferred function runs after the function has returned,
so it cannot change the result received by the caller." Review that claim using this example:

```go
package main

import "fmt"

func buildResult() (result int) {
    result = 1
    defer fmt.Println("logged:", result)
    defer func() { result++ }()
    return 5
}

func main() {
    fmt.Println("returned:", buildResult())
}
```

1. Predict the exact output, in order. Walk through when deferred-call arguments are evaluated,
   when the return expression is evaluated and assigned, when each deferred call runs, and when
   control reaches the caller. Distinguish a named result from an ordinary local variable.
2. If this pattern appeared in an API use case, what could go wrong when a log captures one
   result but deferred code changes the result returned to the caller? Would changing the log
   to a closure always fix it, or does registration order still matter? Explain your review advice.
3. When, if ever, should deferred cleanup change a returned result or error? Explain how you
   would handle both an operation failure and a cleanup failure without hiding important evidence,
   and describe a small test that verifies the caller-visible outcome and the final-outcome log.

Ask this question in G40 and its PR. Explain the proposed checks; implementing them is not a
learner deliverable for this question. No additional payment or persistence implementation is required.

## Initial Non-Goals

Do not add these until explicitly scoped:

- database schema or migrations.
- authentication or authorization.
- payment processor clients.
- ledger posting.
- refunds, disputes, settlement, or reconciliation.
- async workers or message queues.
- generated OpenAPI contracts.
- container or Kubernetes deployment.
- Rust FFI, WASM, or sidecar integration.

## First Implementation Direction

The first Go implementation should establish a minimal service skeleton only after the boundary is
reviewed. The first endpoint should probably be a health/readiness endpoint, not payment creation,
because it teaches service layout and validation discipline without pretending the payment workflow
is ready for public traffic.
