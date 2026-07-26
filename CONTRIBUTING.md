# Contributing to TextForge

Thanks for helping improve TextForge. Contributions can include Go code,
lessons, documentation, tests, and ideas for improving the learning
experience.

## Getting started

Clone the repository and run the tests:

~~~bash
git clone https://github.com/AbhijithKumble/textforge.git
cd textforge
go mod download
go test ./...
~~~

Run the CLI locally with:

~~~bash
go run . learn
go run . practice
~~~

## Project structure

- cmd/ — Cobra commands and terminal UI.
- internal/course/ — lesson discovery and Markdown front matter.
- internal/progress/ — saved exercise progress.
- courses/ — Markdown lessons and exercise metadata.
- docs/ — supporting project documentation.

## Adding a lesson

Add a Markdown file under the appropriate course directory. Include YAML front
matter with an ID, title, index, description, difficulty, and exercises:

~~~yaml
---
id: regex-04-anchors
title: "Regex Anchors"
index: 4
description: "Match text at the beginning or end of a line."
difficulty: "Beginner"
exercises:
  - id: line-start
    prompt: "Find lines that begin with ERROR."
    answer: "^ERROR"
    test_input: "ERROR disk full\nINFO started"
---
~~~

Each exercise should:

- have a clear, task-oriented prompt;
- include test_input when behavior-based validation is appropriate;
- use realistic examples;
- distinguish matching and non-matching cases;
- introduce one main concept at a time.

## Go changes

Use standard Go style and run gofmt on changed files:

~~~bash
gofmt -w main.go cmd internal
~~~

Add tests beside the implementation in files named *_test.go. Cover
successful behavior, invalid input, and persistence changes where relevant.

Before submitting changes, run:

~~~bash
GOCACHE=/tmp/textforge-go-cache CGO_ENABLED=0 go test ./...
go build ./...
~~~

## Commit and pull requests

Use short, imperative commit subjects:

~~~text
Add regex anchor lesson
Improve practice feedback
Document local installation
~~~

Keep unrelated changes separate. Pull requests should describe:

- what changed for users;
- the main files changed;
- how the change was tested;
- screenshots or terminal output for interactive UI changes.

Please update documentation when changing commands, lesson formats, or
user-visible behavior.
