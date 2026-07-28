---
id: regex-07-capturing-groups
title: "Regex Capturing Groups"
index: 7
description: "Extract useful parts from a larger match."
difficulty: "Intermediate"
exercises:
  - id: capture-usernames
    prompt: "Match user entries and capture the username after user=."
    answer: 'user=([A-Za-z]+)'
    captures: true
    test_input: "user=alice\nuser=bob\nadmin=root\nuser=carol"
  - id: capture-config
    prompt: "Capture the key and value from each configuration entry."
    answer: '^([A-Za-z_]+)=([^ ]+)$'
    captures: true
    test_input: "host=api.example.com\nport=8080\ninvalid key=val\nbad=foo bar\nenv=production"
  - id: capture-time
    prompt: "Capture the hour, minute, and second from each timestamp."
    answer: '([0-9]{2}):([0-9]{2}):([0-9]{2})'
    captures: true
    test_input: "started at 09:45:12\nfinished at 17:03:08\ninvalid 9:5:2"
---

# Regex Capturing Groups

## What You Will Learn

By the end of this lesson, you will be able to:

- group part of a pattern with parentheses;
- capture useful text from a larger match; and
- distinguish a complete match from its captured groups.

## The Core Idea

Parentheses do two jobs in regex:

- they group parts of a pattern; and
- they capture the text matched inside the parentheses.

For example, user=([A-Za-z]+) matches the complete text user=alice, while
Group 1 contains only alice.

The complete match is always shown first. Captured groups are numbered from
left to right:

~~~text
Pattern: user=([A-Za-z]+)
Full match: user=alice
Group 1: alice
~~~

## A Developer Scenario

Suppose a configuration file contains key-value pairs:

~~~text
host=api.example.com
port=8080
~~~

To capture the key and value separately, use:

~~~bash
grep -oE '^([A-Za-z_]+)=([^ ]+)$' config.txt
~~~

grep prints the complete matching line. TextForge additionally shows the
captured groups so you can see the extracted key and value.

## Practice

Look at both the full match and each captured group. A pattern can match the
right text but still be wrong if its groups capture the wrong parts.

## Key Takeaways

- Parentheses create capturing groups.
- The full match includes everything matched by the complete pattern.
- Groups contain selected parts of that full match.
- Groups are numbered from left to right.
- Capturing groups are useful for extracting fields from logs and config files.
