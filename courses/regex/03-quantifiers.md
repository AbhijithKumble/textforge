---
id: regex-03-quantifiers
title: "Regex Quantifiers"
index: 3
description: "Match repeated characters with quantifiers."
difficulty: "Beginner"
exercises:
  - id: three-digit-code
    prompt: "Find every sequence containing exactly three consecutive digits."
    answer: "[0-9]{3}"
    test_input: "id=7 retry=42 code=200 port=808"
  - id: one-or-more-letters
    prompt: "Find each sequence containing one or more letters."
    answer: "[A-Za-z]+"
    test_input: "id=7 status=ok"
  - id: optional-letter
    prompt: "Match both color and colour, but not collar."
    answer: "colou?r"
    test_input: "color colour collar"
---

# Regex Quantifiers

## What You Will Learn

By the end of this lesson, you will be able to:

- match an exact number of repeated characters;
- match one or more characters with `+`; and
- make a character optional with `?`.

## The Core Idea

Quantifiers describe how many times the preceding pattern may repeat:

| Quantifier | Meaning | Example |
| --- | --- | --- |
| `+` | one or more | `[A-Za-z]+` |
| `?` | zero or one | `colou?r` |
| `{3}` | exactly three | `[0-9]{3}` |
| `{2,4}` | between two and four | `[0-9]{2,4}` |

The quantifier applies to the character or group immediately before it. In
`[0-9]{3}`, the character class must match three times in a row.

## A Developer Scenario

Suppose a service log contains identifiers and status values:

```text
id=7 retry=42 code=200 port=808
```

To find the three-digit values with `grep`, use extended regular expressions:

```bash
grep -oE '[0-9]{3}' app.log
```

This selects `200` and `808`, but not the one-digit `7` or the two-digit `42`.

## Practice

Pay attention to how many characters each pattern selects. The test input is
shown with each exercise so you can compare short and repeated matches.

## Key Takeaways

- Quantifiers control repetition.
- `[0-9]{3}` matches exactly three consecutive digits.
- `+` requires at least one occurrence.
- `?` makes the preceding character optional.
- Use `grep -E` when writing `+`, `?`, or `{n}` in shell examples.
