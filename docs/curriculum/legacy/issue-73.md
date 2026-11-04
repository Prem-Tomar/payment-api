# Original Issue #73

Historical snapshot before inventory consolidation. This file preserves the original request; it is not completion evidence.

Title: Day G60: Validate merchant scope in the lookup use case

## Goal

Reject missing merchant scope at the application lookup boundary.

## Context

HTTP middleware supplies merchant identity today, but another application caller could invoke the lookup use case directly. A boundary should make required input meaningful even outside one transport path.

## Task

Add a small application input check for absent or whitespace-only merchant identity before accessing the payment-intent store. Return a recognizable application error and map that error to a clear client response in the HTTP layer.

## Completion Criteria

- Missing merchant identity stops lookup before the store is called.
- The application error can be recognized without comparing message text.
- Valid merchant-scoped lookup and not-found behavior remain distinct.
- Operational endpoints remain independent of merchant scope.

## Scope

Implementation only. Tests and documentation are reviewer-owned. No merchant registry, credential validation, authorization policy, database, or general validation framework.

## Research Hints

Search `Go sentinel error errors.Is`, `application service input validation`, and `tenant scoped repository`. A sentinel error is a named error value callers can recognize consistently.

## Question of the Day - G60

Previously learned concepts: Merchant-scope middleware, the lookup use case, and the scoped store interface.

Scenario: A background caller invokes the lookup use case directly with empty merchant identity, bypassing HTTP middleware.

1. Which earlier HTTP guarantee no longer applies? Why should invalid input be distinguished from a valid scoped lookup that simply finds no record?
2. Explain what this example teaches about responsibility or failure boundaries in the system, and relate it to one earlier implementation.

System-design connection: Application boundaries can have callers other than HTTP. Earlier merchant-scoped reads establish the requirement; later reuse should preserve it without assuming every caller passed through Gin.

Answer every part in your own words directly in the relevant PR description or comment; an issue-only answer does not satisfy this placement rule. Explain your assumptions and reasoning. No extra implementation, test-writing, or documentation is required to answer. The reviewer preserves your original answer and marks it Correct, Partially correct, Needs correction, or Awaiting answer, with targeted corrections and follow-ups. This learning response is expected but does not block merging an otherwise correct implementation.


