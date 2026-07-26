# Repository Guidelines

## Project Structure & Module Organization

TextForge is a Go CLI and terminal-based course platform. The executable entry point is `main.go`; Cobra command definitions live in `cmd/`, and reusable application logic belongs in `internal/` (currently the lesson loader is in `internal/course/`). Learning material is stored as Markdown under `courses/`, with the current regex lesson at `courses/regex/01-literals.md`. Project documentation is in `README.md`, `ARCHITECTURE.md`, `ROADMAP.md`, `VISION.md`, and `docs/`. Keep future exercises, recipes, examples, and static assets in their corresponding top-level directories.

## Build, Test, and Development Commands

Run these from the repository root:

```bash
go run .                 # Show the TextForge CLI help
go run . learn           # List available lessons
go run . practice        # Resume the interactive exercise session
go run . progress        # Show saved exercise progress
go build ./...           # Compile all packages
go test ./...            # Run all package tests
gofmt -w main.go cmd internal  # Format changed Go code
```

Progress is saved to `.textforge/progress.json`; use `--progress` to select another file during development.

## Coding Style & Naming Conventions

Follow standard Go style and let `gofmt` determine formatting (tabs, conventional import grouping, and trailing newline). Use mixed-case exported names with Go doc comments, lower-case package names, and descriptive lower-case file names. Cobra command variables should follow the existing `*Cmd` pattern (for example, `learnCmd`). Keep course metadata and content compatible with the loader’s YAML front matter format.

## Testing Guidelines

Use Go’s built-in `testing` package and place tests beside the implementation in files named `*_test.go` (for example, `internal/course/course_test.go`). Cover successful lesson loading, malformed front matter, missing files, and progress persistence. Run `go test ./...` before submitting changes; add focused tests for new parsing or command behavior.

## Commit & Pull Request Guidelines

The existing history is small and uses brief, direct summaries. Write concise imperative subjects such as `Add lesson metadata validation`; keep unrelated changes in separate commits. Pull requests should explain the user-visible behavior, identify the main files changed, include test commands and results, and link a related issue when one exists. Include terminal output or screenshots when changing interactive CLI behavior.
