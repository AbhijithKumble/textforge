# TextForge — One-Month TODO

Goal: ship a reliable pre-alpha MVP that teaches a complete beginner regex
course from the terminal, using realistic developer problems and a
Rustlings-style practice loop.

## Current Baseline

- [x] Cobra CLI foundation and version command
- [x] Markdown/YAML lesson loading from `courses/`
- [x] Interactive practice with retry and local progress persistence
- [x] First lesson: regex literals
- [x] Basic loader and progress tests
- [ ] Working `search` and `play` commands
- [ ] Complete beginner regex course
- [ ] CI, release workflow, and stable documentation

## Week 1 — Stabilize the Foundation (Must Have)

- [ ] Write `ARCHITECTURE.md` describing commands, lesson metadata, exercise
  validation, and progress storage.
- [ ] Remove generated/dead Cobra boilerplate and make every exposed command
  return useful errors.
- [ ] Add `go vet ./...`, formatting, and test instructions to the contributor
  workflow.
- [ ] Add GitHub Actions CI for `gofmt` checking, `go vet ./...`, and
  `go test ./...`.
- [ ] Define and validate the lesson schema: required IDs, unique exercise IDs,
  non-empty prompts/answers, and deterministic lesson ordering.
- [ ] Add a safe progress reset command or documented reset workflow.

## Week 2 — Finish the Exercise Engine (Must Have)

- [ ] Separate lesson explanation, exercise prompt, test input, expected output,
  hint, and solution metadata; keep all lesson content outside Go code.
- [ ] Add hidden test cases so exercises validate behavior, not only one exact
  answer string.
- [ ] Support hints, clear failure messages, retry, quit, and resume behavior.
- [ ] Make `practice` handle multiple lessons cleanly and show accurate course
  completion statistics.
- [ ] Add command-level tests for `learn`, `practice`, and `progress` using
  temporary courses and progress files.

## Week 3 — Build the Regex Course (Must Have)

- [ ] Add lessons in this order: literals, character classes, quantifiers,
  anchors, groups, alternation, escaping, greedy/lazy matching, and practical
  log/config extraction.
- [ ] Give every lesson a real-world scenario, a short explanation, examples,
  at least three exercises, hints, and hidden cases.
- [ ] Add small fixtures under `examples/` for logs, CSV data, and config files.
- [ ] Review every lesson for beginner clarity, Unix relevance, and progressive
  difficulty; dogfood the full course with a clean progress file.

## Week 4 — Complete the MVP Surface (Should Have)

- [ ] Implement `textforge search <term>` over local recipes in `recipes/`, with
  regex, explanation, mistakes, and grep/sed/awk examples.
- [ ] Implement `textforge play` as an offline regex tester with input text,
  match results, and capture-group output.
- [ ] Improve `progress` with per-course totals, completion percentage, and the
  next recommended lesson.
- [ ] Update README, architecture docs, quick-start instructions, and version
  information to match the shipped behavior.
- [ ] Add a tagged pre-alpha release workflow and a short demo recording or
  terminal transcript.

## Definition of Done

From a clean checkout, a new user can run `go run . learn`, read the course,
run `go run . practice`, receive immediate feedback, quit safely, resume later,
and inspect progress. `go test ./...`, `go vet ./...`, and CI pass, and every
lesson teaches a practical text-processing problem rather than syntax alone.

## Explicit Non-Goals This Month

Do not build cloud accounts, authentication, telemetry, AI lesson generation,
a GUI, or the future grep/sed/awk/jq courses. Keep the month focused on the
offline regex learning platform.
