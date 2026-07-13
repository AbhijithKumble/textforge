## Why TextForge?

There are already excellent regex resources like RegexOne, RegexLearn, regex101, and regular-expressions.info.

TextForge doesn't aim to replace them.

Instead, it focuses on something different:

- Learn by solving realistic developer tasks instead of isolated syntax exercises.
- Learn entirely from the terminal.
- Progress from regex to grep, sed, awk, jq, and other Unix text-processing tools.
- Practice with real-world datasets such as logs, CSV files, `/etc/passwd`, and configuration files.
- Search for solutions when you know the problem but not the regex.

TextForge is designed to teach **how developers actually use regex at work**, not just how regex syntax works.


# 🗺️ TextForge Roadmap

> **Master Unix text processing through interactive terminal exercises.**

This roadmap outlines the planned direction of **TextForge**. It is intended as a guide, not a strict schedule. Features and priorities may change based on community feedback and project evolution.

---

# Philosophy

TextForge is built around a few core principles:

* Learn by solving real problems.
* Learn inside the terminal.
* Practice over memorization.
* Small lessons with immediate feedback.
* Open source and community driven.

The goal is **not** to become another documentation website or cheat sheet.

The goal is to become the easiest and most enjoyable way to learn Unix text processing.

---

# Current Status

**Current Version**

```
Pre-Alpha
```

**Current Focus**

* Project foundation
* CLI architecture
* Lesson engine
* Exercise engine

---

# Version Roadmap

---

# v0.1 — Foundation

**Goal**

Build the project foundation.

### Features

* Project structure
* Cobra CLI
* CI/CD
* Release workflow
* Documentation
* Version command

### Commands

```bash
textforge
textforge version
textforge learn
textforge practice
textforge search
```

---

# v0.2 — Lesson Engine

**Goal**

Support loading lessons from external files.

### Features

* Course loader
* Markdown lessons
* YAML metadata
* Lesson navigation
* Course discovery

Directory structure

```
courses/
    regex/
        01-literals/
        02-character-classes/
```

---

# v0.3 — Exercise Engine

**Goal**

Interactive learning similar to Rustlings.

### Features

* Interactive exercises
* Hidden test cases
* Hint system
* Retry support
* Pass/Fail validation

Example

```
$ textforge practice

Question:
Extract all numbers.

Regex:
```

---

# v0.4 — Regex Course

**Goal**

Complete beginner-friendly regex course.

Topics include:

* Literals
* Character Classes
* Quantifiers
* Groups
* Alternation
* Anchors
* Escaping
* Greedy vs Lazy
* Practical Regex

Lessons will focus on solving real developer tasks instead of memorizing syntax.

---

# v0.5 — Regex Recipe Search

**Goal**

Search practical regex recipes.

Examples

```
textforge search email

textforge search ipv4

textforge search uuid

textforge search whitespace
```

Every recipe includes

* Regex
* Explanation
* Examples
* Common mistakes
* grep example
* sed example
* awk example

---

# v0.6 — Playground

**Goal**

Offline regex playground.

Features

* Live matching
* Capture groups
* Match highlighting
* Test input

Example

```
textforge play
```

---

# v0.7 — Progress Tracking

Track learning progress.

Features

* Completed lessons
* Resume learning
* Statistics
* Course completion
* Local progress storage

---

# v1.0 — Regex Academy

The first stable release.

Includes

* Complete Regex Course
* Search
* Playground
* Progress Tracking
* Documentation
* Tests
* Examples

---

# Future Courses

After Regex, TextForge expands into the Unix text-processing ecosystem.

## grep

Topics

* Basic Search
* Recursive Search
* Context
* Counting
* Regular Expressions
* Pipelines

---

## sed

Topics

* Substitution
* Addressing
* Groups
* Multiple Commands
* Files
* Practical Editing

---

## awk

Topics

* Fields
* Variables
* Conditions
* Loops
* Functions
* CSV Processing

---

## find

Topics

* Name Matching
* File Types
* Permissions
* Time Filters
* exec
* xargs Integration

---

## jq

Topics

* JSON Queries
* Arrays
* Objects
* Filters
* Pipelines

---

## Shell Pipelines

Topics

* Pipes
* stdin
* stdout
* stderr
* Command Chaining
* Process Substitution

---

# Long-Term Vision

TextForge aims to become the interactive learning platform for Unix text processing.

Future possibilities include

* Daily Challenges
* Community Courses
* Plugin System
* Interactive Terminal UI
* Multi-language Support
* Community Lesson Packs
* VS Code Extension
* Web Companion
* Achievement System

---

# What We Are NOT Building

To keep the project focused, these are intentionally out of scope for the foreseeable future.

* User accounts
* Cloud sync
* AI-generated lessons
* Online authentication
* Complex analytics
* Enterprise features

TextForge should remain lightweight, fast, offline-first, and developer-friendly.

---

# Contributing

Contributions are welcome.

You can contribute by

* Writing lessons
* Adding recipes
* Improving documentation
* Reporting bugs
* Improving the CLI
* Creating new courses

See **CONTRIBUTING.md** for details.

---

# Guiding Principle

Every feature should help answer one question:

> **Will this help someone become better at solving real-world text-processing problems?**

If the answer is **yes**, it belongs in TextForge.

If not, it should probably wait.

