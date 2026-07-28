---
id: regex-05-groups-alternation
title: "Regex Groups and Alternation"
index: 5
description: "Match one pattern from several alternatives."
difficulty: "Beginner"
exercises:
  - id: log-levels
    prompt: "Find every log entry with either ERROR or WARN."
    answer: "(ERROR|WARN)"
    test_input: "INFO started\nERROR disk full\nWARN high memory\nDEBUG retrying\nERROR timeout"
  - id: api-methods
    prompt: "Find GET or POST requests to /users, but not DELETE requests."
    answer: "(GET|POST) /users"
    test_input: "GET /users\nPOST /users\nDELETE /users\nGET /health"
  - id: deployment-environments
    prompt: "Find configuration entries for development or production."
    answer: "env=(development|production)"
    test_input: "env=development\nenv=test\nenv=production\nenv=staging"
---

# Regex Groups and Alternation

## What You Will Learn

By the end of this lesson, you will be able to:

- match one pattern from several alternatives with |;
- group alternatives with parentheses; and
- combine a group with surrounding literal text.

## The Core Idea

The alternation operator | means “or”. Parentheses group the alternatives
into one pattern:

| Pattern | Meaning |
| --- | --- |
| ERROR\|WARN | match ERROR or WARN |
| (ERROR\|WARN) | group the same alternatives |
| env=(development\|production) | match a complete configuration value |

Without grouping, the surrounding text may apply only to one side of the
alternation. For example, GET|POST /users can match GET anywhere, while
(GET|POST) /users requires both alternatives to be followed by /users.

## A Developer Scenario

Suppose an API log contains several request methods:

~~~text
GET /users
POST /users
DELETE /users
~~~

To find requests that read or create users, use:

~~~bash
grep -oE '(GET|POST) /users' app.log
~~~

This selects the GET and POST requests but leaves the DELETE request out.

## Practice

Each test input includes valid alternatives and near-matches. Compare the
number and content of selected matches when you try a broader or narrower
pattern.

## Key Takeaways

- | means “or”.
- Parentheses group alternatives.
- Grouping lets surrounding text apply to every alternative.
- Use grep -E for grouped and alternation patterns.
