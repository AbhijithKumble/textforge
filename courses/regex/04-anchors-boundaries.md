---
id: regex-04-anchors-boundaries
title: "Regex Anchors and Boundaries"
index: 4
description: "Match text at specific positions."
difficulty: "Beginner"
exercises:
  - id: line-start
    prompt: "Find ERROR only when it appears at the start of a line."
    answer: "^ERROR"
    test_input: "ERROR disk full\nERROR timeout\nINFO ERROR disk full\nWARN started"
  - id: line-end
    prompt: "Find filenames that end with .log."
    answer: "[.]log$"
    test_input: "app.log\nserver.log\nnotes.txt\nbackup.log.tmp\napp.log.bak"
  - id: whole-word
    prompt: "Match ERROR as a complete word, but not ERROR_CODE."
    answer: '\bERROR\b'
    test_input: "ERROR ERROR_CODE CODE_ERROR INFO ERROR"
---

# Regex Anchors and Boundaries

## What You Will Learn

By the end of this lesson, you will be able to:

- match text at the beginning of a line with ^;
- match text at the end of a line with $; and
- match a complete word with \b.

## The Core Idea

Some regex symbols describe a position instead of matching a character:

| Symbol | Meaning | Example |
| --- | --- | --- |
| ^ | beginning of a line | ^ERROR |
| $ | end of a line | [.]log$ |
| \b | word boundary | \bERROR\b |

The pattern ERROR matches those letters anywhere in a line. The pattern
^ERROR matches them only when they start the line. Similarly, ERROR$ would
require the text to end with ERROR.

## A Developer Scenario

Suppose a log contains messages and filenames:

~~~text
ERROR disk full
INFO ERROR disk full
app.log
notes.txt
~~~

To find only log lines that begin with ERROR, use:

~~~bash
grep -E '^ERROR' app.log
~~~

To find filenames ending in .log, use:

~~~bash
grep -E '[.]log$' files.txt
~~~

## Practice

Compare the selected output carefully. The test input contains near-matches
that fail because the text starts or ends in the wrong place, or because the
word is part of a larger identifier.

## Key Takeaways

- Anchors match positions, not visible characters.
- ^ restricts a match to the beginning of a line.
- $ restricts a match to the end of a line.
- \b prevents a word from matching inside a larger word.
