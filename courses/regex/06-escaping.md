---
id: regex-06-escaping
title: "Regex Escaping"
index: 6
description: "Match regex metacharacters as literal text."
difficulty: "Beginner"
exercises:
  - id: literal-dot
    prompt: "Find filenames that end with the literal .log suffix."
    answer: '\.log$'
    test_input: "app.log\nserver.log\nappXlog\nserver.log.bak"
  - id: literal-plus-question
    prompt: "Find literal plus and question-mark characters."
    answer: '(\+|\?)'
    test_input: "C++ uses ?\nC# uses #\nC+ uses +\nunknown"
  - id: literal-delimiters
    prompt: "Find [ERROR], (WARN), and INFO|DEBUG as literal text."
    answer: '(\[ERROR\]|\(WARN\)|INFO\|DEBUG)'
    test_input: "[ERROR] (WARN) INFO|DEBUG\nERROR WARN INFO DEBUG\n[INFO] (DEBUG)"
  - id: literal-backslash
    prompt: "Find every literal backslash in the paths and escape sequence."
    answer: '\\'
    test_input: |-
      path=C:\tmp
      path=/var/log
      escape=\n
---

# Regex Escaping

## What You Will Learn

By the end of this lesson, you will be able to:

- match punctuation that normally has a regex meaning;
- escape operators such as plus and question mark;
- escape brackets, parentheses, and pipes; and
- match a literal backslash.

## The Core Idea

Some characters have special meanings in regex:

| Character | Normal meaning | Escaped form |
| --- | --- | --- |
| . | any character | \. |
| + | one or more | \+ |
| ? | optional | \? |
| [ ] | character class | \[ \] |
| ( ) | group | \( \) |
| \| | alternation | \| |

Put a backslash before a metacharacter when you want to match that character
literally. For example, a dot matches any character, but \. matches only a
literal dot.

To match one literal backslash, the regex pattern is two backslashes: \\.

## A Developer Scenario

Suppose a directory listing contains log files:

~~~text
app.log
server.log
server.log.bak
~~~

To find names ending in the literal .log suffix, use:

~~~bash
grep -oE '\.log$' files.txt
~~~

Use single quotes around shell regexes so the shell preserves the backslashes
before grep receives the pattern.

## Practice

Each test input includes valid matches and near-matches. Compare the selected
output when you leave a metacharacter unescaped.

## Key Takeaways

- A backslash changes a metacharacter into a literal character.
- A dot is not the same as a literal dot.
- Operators such as +, ?, and | need escaping when matched as text.
- Brackets, parentheses, and backslashes also need careful escaping.
- Single quotes help preserve regex backslashes in shell commands.
