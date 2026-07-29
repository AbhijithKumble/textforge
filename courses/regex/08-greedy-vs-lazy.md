---
id: regex-08-greedy-vs-lazy
title: "Greedy vs. Lazy Matching"
index: 8
description: "Understand the difference between greedy and lazy quantifiers."
difficulty: "Intermediate"
exercises:
  - id: extract-html-tags
    prompt: "Match individual HTML tags (like <b> or </b>) without spanning across multiple tags on the same line."
    answer: '<.*?>'
    test_input: "<b>Hello</b> <i>World</i>\n<div>Content</div>\n<p>First</p> <p>Second</p>"
  - id: extract-quoted-strings
    prompt: "Match text enclosed in double quotes without spanning across multiple quoted values on the same line."
    answer: '".*?"'
    test_input: 'name="alice" role="admin"\nkey="value1" extra="value2"\nsingle="val"'
  - id: extract-bracketed-tags
    prompt: "Match each bracketed tag like [INFO] or [AUTH] individually without combining adjacent brackets."
    answer: '\[.*?\]'
    test_input: "[INFO] [AUTH] User login\n[WARN] [DB] Connection pool low\n[ERROR] Server crash"
---

# Greedy vs. Lazy Matching

## What You Will Learn

By the end of this lesson, you will be able to:

- understand the default **greedy** behavior of quantifiers (`*`, `+`, `?`, `{n,m}`);
- use **lazy** (non-greedy) quantifiers by appending a `?` (`*?`, `+?`, `??`, `{n,m}?`); and
- prevent quantifiers from accidentally matching across multiple delimiters like quotes, brackets, or HTML tags.

## The Core Idea

By default, regular expression quantifiers are **greedy**. A greedy quantifier matches as much text as possible while still allowing the rest of the pattern to match.

For example, given this HTML string:

```text
<b>Hello</b> <i>World</i>
```

The pattern `<.*>` matches from the first `<` all the way to the **last** `>` on the line:

```text
<b>Hello</b> <i>World</i>  <-- Full line matched!
```

Because `.*` is greedy, it consumes `b>Hello</b> <i>World` as part of a single match.

To stop at the **first** matching closing bracket, append a `?` to make the quantifier **lazy** (or non-greedy):

```text
<.*?>
```

The lazy quantifier `*?` matches as **few** characters as possible:

```text
Match 1: <b>
Match 2: </b>
Match 3: <i>
Match 4: </i>
```

## Greedy vs. Lazy Comparison

| Quantifier | Type | Behavior | Example on `"a" "b"` | Match Result |
| :--- | :--- | :--- | :--- | :--- |
| `.*` | Greedy | Matches maximum characters | `".*"` | `"a" "b"` |
| `.*?` | Lazy | Matches minimum characters | `".*?"` | `"a"` and `"b"` |
| `.+` | Greedy | Matches 1 or more (max) | `".+"` | `"a" "b"` |
| `.+?` | Lazy | Matches 1 or more (min) | `".+?"` | `"a"` and `"b"` |

## A Developer Scenario

Suppose you are parsing a log file containing bracketed metadata:

```text
[2026-07-28 20:45:00] [INFO] [AUTH] User logged in
```

Using `\[.*\]` will match `[2026-07-28 20:45:00] [INFO] [AUTH]`.

To capture each bracketed tag independently for further processing with `grep -oE`, use a lazy quantifier:

```bash
grep -oE '\[.*?\]' server.log
```

This outputs:

```text
[2026-07-28 20:45:00]
[INFO]
[AUTH]
```

## Key Takeaways

- Standard quantifiers (`*`, `+`, `?`) are **greedy** and consume as much text as possible.
- Appending `?` (`*?`, `+?`, `??`) turns them into **lazy** quantifiers that match as little text as possible.
- Use lazy quantifiers when matching text between paired delimiters such as quotes (`".*?"`), brackets (`\[.*?\]`), or HTML/XML tags (`<.*?>`).
