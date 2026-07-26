---
id: regex-01-literals
title: "Regex Literals"
index: 1
description: "Learn how to match exact text characters."
difficulty: "Beginner"
exercises:
  - id: timeout
    prompt: "Find the word timeout in a configuration file."
    answer: "timeout"
  - id: status-ok
    prompt: "Find the exact text status=ok in a service log."
    answer: "status=ok"
  - id: error
    prompt: "Find ERROR without matching lowercase error."
    answer: "ERROR"
---

# Regex Literals

## What You Will Learn

By the end of this lesson, you will be able to:

- match an exact sequence of characters;
- recognize that literals are case-sensitive; and
- use a literal pattern to find text in a realistic configuration file.

## The Core Idea

A literal is the simplest regular expression: it matches the characters you
type, in the order you type them. The pattern `cat` matches `cat` in this text:

```text
The cat is sleeping.
```

It does not match `Cat`, `catalog`, or `catch` unless your search tool allows a
match to occur inside a larger word. Regex engines usually search for a match
anywhere in the input, so `cat` can match the first three characters of
`catalog`.

Literals are normally case-sensitive. The pattern `port` matches `port`, but
not `Port` or `PORT`.

## A Developer Scenario

Suppose `app.conf` contains:

```text
host=localhost
port=8080
protocol=https
```

To find the port setting with `grep`, use:

```bash
grep 'port' app.conf
```

The literal `port` matches the `port=8080` line. It also matches `protocol`
because that line begins with the same four characters. Later lessons will
show how anchors and other syntax can make a match more precise.

## Practice

For each task, write a literal regex pattern:

1. Find the word `timeout` in a configuration file.
2. Find the exact text `status=ok` in a service log.
3. Find `ERROR` without matching lowercase `error`.

Try each pattern against several input lines, including lines that differ only
in capitalization or contain the pattern as part of a longer word.

## Check Your Answers

The patterns are `timeout`, `status=ok`, and `ERROR`, respectively. Each answer
uses only the characters that must appear in the target text; no special regex
syntax is needed yet.

## Key Takeaways

- Plain text is a valid regex pattern.
- Character order matters.
- Literal matching is usually case-sensitive.
- A literal may match part of a larger word; use anchors when exact boundaries
  matter.
