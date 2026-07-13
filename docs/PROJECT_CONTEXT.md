# TextForge Project Context

## What is TextForge?

TextForge is an open-source interactive CLI that teaches developers
regex and Unix text-processing through hands-on exercises.

It is inspired by Rustlings.

The project is NOT a regex engine.

The project is NOT another cheat sheet.

The project is a learning platform.

---

## Target Audience

- Developers
- Linux users
- CS students
- DevOps engineers

---

## Learning Philosophy

Teach practical skills through solving realistic terminal problems.

Never teach syntax in isolation.

Every lesson should answer

"When will I actually use this?"

---

## Core Principles

- Learn by doing
- Small lessons
- Immediate feedback
- Offline-first
- Fast CLI
- Community-driven
- Open Source

---

## MVP

Regex only.

Commands

textforge learn

textforge practice

textforge search

textforge play

textforge progress

---

## Future Courses

grep

sed

awk

jq

find

shell

---

## Things We Are NOT Building

Cloud accounts

Authentication

AI lesson generation

Telemetry

Heavy GUI

---

## Technology

Language

Go

CLI

Cobra

Storage

YAML + Markdown

Tests

Go testing

License

MIT

---

## Design Rule

Code should never contain lesson content.

Lessons live under

courses/

Recipes live under

recipes/

Go only implements the engine.

---

## North Star

Become the Rustlings for Unix text processing.
