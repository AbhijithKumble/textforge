# TextForge LLM Handoff

This document preserves the current project context for another LLM or
contributor continuing work on TextForge.

## Project goal

TextForge is a Go/Cobra terminal course platform intended to become the
Rustlings of Unix text processing. The current course teaches regex through
short lessons and interactive practice exercises.

The product priorities are:

- practical, command-focused learning;
- small lessons with immediate feedback;
- behavior-based regex validation instead of literal answer comparison;
- saved progress after every correct answer; and
- readable terminal-first learning and practice interfaces.

## Important repository rules

Read AGENTS.md before changing the repository. It already contains the
repository contribution rules and must be preserved.

Do not reset, discard, or overwrite unrelated working-tree changes. The
repository may contain lesson files that are intentionally untracked while
they are being developed.

Use apply_patch for source and documentation edits. Format Go changes with
gofmt.

## Current course sequence

The regex course currently contains seven lessons:

1. regex-01-literals — exact text matching.
2. regex-02-character-classes — sets and ranges such as [A-Z].
3. regex-03-quantifiers — repetition with +, ?, and {n}.
4. regex-04-anchors-boundaries — ^, $, and word boundaries.
5. regex-05-groups-alternation — parentheses and |.
6. regex-06-escaping — matching metacharacters literally.
7. regex-07-capturing-groups — extracting submatches.

Lessons are Markdown files with YAML front matter under courses/regex/.
Every behavior-tested exercise should include test_input. Capturing-group
exercises additionally set captures: true.

## Practice validation behavior

The practice command is implemented in cmd/practice.go.

For ordinary exercises with test_input:

- the expected and submitted patterns are compiled with Go regexp;
- test input is processed line-by-line to match grep behavior for ^ and $;
- all matches are compared, so equivalent patterns can pass;
- invalid patterns fail with a readable error; and
- feedback shows the submitted result, expected result, match counts, and
  highlighted input.

Terminal feedback uses ANSI colors:

- submitted matches are red;
- expected matches are green.

For exercises with captures: true:

- FindAllStringSubmatch is used;
- full matches and capture groups are both compared;
- feedback displays each full match and Group 1, Group 2, and so on.

Exercises without test_input retain exact string comparison. This preserves
support for non-regex or intentionally literal answers.

## Current user interface

learn with no lesson ID displays a course overview containing lesson number,
title, difficulty, description, exercise count, and commands for reading or
practicing.

learn with a lesson ID displays lesson metadata, the lesson Markdown content,
and a practice command.

practice displays:

- lesson and exercise progress;
- the task prompt;
- multiline test input with indentation;
- a clear regex entry prompt;
- colored match feedback;
- capture groups where enabled; and
- saved-progress messages.

Users can enter q or quit to leave practice. Correct answers are saved after
each exercise.

## Documentation

- README.md — project overview and roadmap.
- INSTALL.md — clone, build, run, test, and reset-progress instructions.
- CONTRIBUTING.md — contributor workflow and lesson format.
- AGENTS.md — repository-specific agent and contribution instructions.
- LLM_HANDOFF.md — this context-restoration document.

## Verification command

Run this from the repository root:

~~~bash
GOCACHE=/tmp/textforge-go-cache CGO_ENABLED=0 go test ./...
~~~

Also useful:

~~~bash
go run . learn
go run . learn regex-07-capturing-groups
go run . practice regex-07-capturing-groups
go run . progress
go build ./...
~~~

Use a temporary progress file when manually testing:

~~~bash
go run . practice --progress /tmp/textforge-progress.json
~~~

## Known considerations for future work

- Keep lesson exercises behavior-focused and include multiple positive and
  negative cases so overly broad patterns visibly fail.
- Preserve line-by-line matching unless the course deliberately introduces
  multiline regex behavior.
- Keep expected and submitted feedback visually distinct; red and green
  highlighting alone is not sufficient without clear labels and counts.
- When adding capture lessons, update both the Exercise metadata and the
  capture-aware validation tests.
- Search and play remain separate command areas and should not be assumed to
  be complete unless their implementations are inspected.
- Do not claim a commit or release unless the user explicitly requests it and
  the repository state confirms it.
