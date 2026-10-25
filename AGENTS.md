# Payment API working rules

Read `README.md` and `docs/understanding.md` before creating learning issues or reviewing PRs.
Keep work within this repository and the task's stated learning scope.
Use a review branch and PR; do not push or merge directly to `main` without Prem's specific approval.

## Question of the day

- Every new issue must include one interviewer-style question about previous completed implementations
  or payment-system concepts already learned in this project,
  using the workflow in `docs/understanding.md`. This applies to issues created through tools
  as well as the GitHub issue template.
- Identify the earlier implemented topic, its issue or PR, and the relevant code path. Verify that
  the implementation exists; a roadmap item, open task, or discussion alone is not evidence of
  completion. For a payment concept, cite the earlier implementation, issue, or review where it
  was learned. Rotate across Go mechanics, service design, and payment-system reasoning. Do not
  ask about future topics or treat the newly assigned task as prior learning. Never invent an answer.
- Probe commonly missed mechanisms and staff/principal-level judgment: failure modes, ownership,
  tradeoffs, observability, and verification. Use a focused scenario with progressive follow-ups;
  avoid trivia-only questions and assess reasoning as well as factual accuracy.
- Starting with G40, repeat each issue's exact question in the associated PR. The developer must
  answer every question directly in the PR description or a PR comment, in their own words.
  An issue-only answer or a link to an issue answer does not satisfy this rule. For a PR covering
  multiple issues, include and answer each issue's question separately.
- Review the answer for correctness and explain any corrections alongside the original answer.
- Make feedback granular: explain each claim's mechanism, starting state, state changes, timing,
  and observable consequence. Trace relevant code statement by statement; explain why values
  arise (including zero initialization), rather than only naming the correct value. Distinguish
  evaluation, assignment, deferred invocation, and return to the caller where relevant. Apply
  the same causal detail to service and payment concepts, not just Go execution examples.
- A missing, incomplete, or incorrect answer, or a pending answer review, never blocks merging
  an otherwise correct implementation. Keep learning feedback separate from implementation
  findings and merge decisions. Do not add a required CI check for the answer.
