# Learning And Review Rules

- Learner assignments deliver a connected result in the real code path. At this stage,
  combine 2-4 related goals rather than assigning one helper or type at a time.
  Reviewers own routine tests and documentation; do not add these as learner requirements.
- Read `docs/learning-and-interview-guide.md` and `docs/curriculum/index.md` when planning
  assignments or coaching. Preserve source requirements when consolidating inventory.
  Close merged source issues as superseded/not planned, never as evidence of achievement.
- Use visual-first technical coaching in issues, PR summaries, inline findings, QOD reviews,
  and follow-up explanations: a small labelled diagram or numbered execution trace, then
  a short explanation and one focused question. Use Mermaid on GitHub and a text fallback.
  Add a relevant visual for every problem statement and coaching section. Status-only
  updates do not need a decorative diagram. Explain what the visual shows in plain words.
- Use short sentences and common English. Define technical terms beside their first use.
  Keep exact code names visible. Accept short answers, bullets, and sketches; assess
  understanding rather than English fluency. Do not reduce technical depth for language.
- Every assignment includes Small, Moderate, and Strong interview problem statements,
  each with a visual. Start with Small active. Moderate/Strong are labelled previews and
  stay inactive until the learner explains the simpler case and an edge case with little
  help and is comfortable. Do not advance based only on time, task number, or code volume.
- Connect implementation to software design principles using concrete responsibility
  decisions. Cover SOLID, GRASP Creator, creational patterns, and behavioral patterns
  gradually. Define each name and distinguish Creator responsibility from a factory
  mechanism. Prefer idiomatic Go functions/interfaces; never force a named pattern.
- Database implementation may start now at G65 using the existing store boundary.
  Earlier open Go goals keep their scope: database content there is a coaching preview.
  Follow the PostgreSQL/SQL progression in `docs/sql-learning-guide.md`. Database setup,
  schema, queries, transactions, and query analysis are implementation learning, not
  documentation substitutes. Keep financial arithmetic and lifecycle invariants Rust-owned.
- Every new or successor learning issue MUST include a concrete Question of the Day.
  Use `.github/ISSUE_TEMPLATE/learning-task.md`. Replace all placeholders before publishing.
  After creating or editing an issue, read its saved body back and verify the question,
  three interview levels, active-level label, visuals, and scope are present.
- A Question of the Day revisits previously implemented concepts. Ask for reasoning about
  a short scenario, behavior, failure, or tradeoff; do not include the answer or a full solution.
- Explain what the concept is, why it matters to system design, where it fits, how it builds
  on earlier learning, and where later work may use it. Define unfamiliar terms briefly.
  Future connections are previews, not assumed knowledge or added implementation scope.
- Learners answer every part in their own words directly in the relevant PR discussion.
  Preserve original answers; review them as Correct, Partially correct, Needs correction,
  or Awaiting answer, with targeted corrections and follow-ups.
- Missing or incomplete learning answers are non-blocking for an otherwise correct PR.
  Keep learning assessment separate from implementation defects and validation results.
- Issues include a goal, context, bounded task, completion criteria, exclusions, and research
  hints. Hints should help learners search and reason without dictating the complete solution.
- PR feedback includes a master summary and tight inline comments explaining the ask,
  involved components, consequence, design connection, and a useful research hint.
  Resolve achieved or superseded feedback while preserving remaining actions clearly.
- Close an achieved implementation issue and create one substantial successor, assigned to
  `vishaltomar004`, with its own Question of the Day. Do not expand review scope to future work.
  Consolidation is not completion and does not trigger one new issue per archived source.
- Do not implement learner-owned application logic during a review unless explicitly asked.
  Maintain review-owned tests when the intended contract evolves; retain failing regressions
  for actual gaps and report their meaning honestly. Follow README validation commands.
- Merge only with explicit user authorization. Work through PRs rather than pushing to main.
