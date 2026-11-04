# Original Issue #59

Historical snapshot before inventory consolidation. This file preserves the original request; it is not completion evidence.

Title: Day G50: Add create payment intent command boundary

Goal: introduce a small application command struct for create-payment-intent input so the HTTP handler stops passing only context into the use case. Learning focus: separating transport DTOs from application input. Implementation task: define a command type in the application layer with the fields the create use case needs, and update the interface shape to accept it without adding real payment behavior. Rules: no documentation task, no test requirement, no database, no processor client, no ledger behavior, no Rust integration. Completion criteria: the handler can map the Gin request DTO into an application command, the command is easy for a beginner to trace, and placeholder behavior remains explicit. Research hints: search for Go command object pattern, DTO to application mapping, and thin HTTP handler pattern; compare passing many parameters with passing one small struct.

## Question of the Day - G50

Previously learned concepts: The create use-case interface and separate HTTP/application data types.

Scenario: A handler binds a JSON DTO, then calls an application service through an interface.

1. Why might passing the DTO or Gin context directly make that service harder to reuse? Trace which information belongs in a small application input and which stays specific to HTTP.
2. Explain what this example teaches about responsibility or failure boundaries in the system, and relate it to one earlier implementation.

System-design connection: An application boundary separates one use case from its transport. Earlier interfaces make replacement possible; later callers can reuse the use case through another transport.

Answer every part in your own words directly in the relevant PR description or comment; an issue-only answer does not satisfy this placement rule. Explain your assumptions and reasoning. No extra implementation, test-writing, or documentation is required to answer. The reviewer preserves your original answer and marks it Correct, Partially correct, Needs correction, or Awaiting answer, with targeted corrections and follow-ups. This learning response is expected but does not block merging an otherwise correct implementation.


