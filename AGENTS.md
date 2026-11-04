# Learning And Review Rules

- Learner issues teach through a small implementation in the real code path.
  Reviewers own routine tests and documentation; do not add these as learner requirements.
- Every new or successor learning issue MUST include a concrete Question of the Day.
  Use `.github/ISSUE_TEMPLATE/learning-task.md`. Replace all placeholders before publishing.
  After creating or editing an issue, read its saved body back and verify the question exists.
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
- Close an achieved implementation issue and create one focused successor, assigned to
  `vishaltomar004`, with its own Question of the Day. Do not expand review scope to future work.
- Do not implement learner-owned application logic during a review unless explicitly asked.
  Maintain review-owned tests when the intended contract evolves; retain failing regressions
  for actual gaps and report their meaning honestly. Follow README validation commands.
- Merge only with explicit user authorization. Work through PRs rather than pushing to main.
