---
id: regex-02-character-classes
title: "Regex Character Classes"
index: 2
description: "Match one character from a defined set or range."
difficulty: "Beginner"
exercises:
  - id: digit
    prompt: "Find any single digit in a user ID such as user42."
    answer: "[0-9]"
    test_input: "user42 userA"
  - id: hex-character
    prompt: "Given the request ID request_id=F02A7C, find one hexadecimal character."
    answer: "[A-Fa-f0-9]"
    test_input: "A f 7 G z 0 9"
  - id: uppercase-level
    prompt: "Find every uppercase letter in log levels such as INFO, WARN, or ERROR."
    answer: "[A-Z]"
    test_input: "INFO WARN ERROR ok"
---

# Regex Character Classes

## What You Will Learn

By the end of this lesson, you will be able to:

- match one character from a list of allowed characters;
- use ranges such as `[0-9]` and `[A-Z]`; and
- exclude characters with a negated class such as `[^\"]`.

## The Core Idea

A character class matches exactly one character from the characters inside
the square brackets. The pattern `[aeiou]` matches one vowel:

```text
The quick brown fox.
```

It can match each vowel individually, but it does not match `qu` as a pair or
the complete word `quick`.

You can write ranges with a hyphen. `[0-9]` matches one digit, and `[A-Fa-f]`
matches one letter from `A` through `F`, in either case. Character classes
are useful when the exact character is unknown but its allowed set is known.

To exclude characters, put `^` immediately after the opening bracket. The
pattern `[^\"]` matches one character that is not a double quote. The `^` has
this special meaning only at the beginning of a character class; elsewhere it
is treated differently.

## A Developer Scenario

Suppose a service log contains request IDs and status information:

```text
request_id=ab39f1 status=200
request_id=F02A7C status=500
```

To find a single hexadecimal character, use:

```bash
grep -o '[A-Fa-f0-9]' app.log
```

The class `[A-Fa-f0-9]` accepts the characters used in hexadecimal values:
digits `0` through `9` and letters `A` through `F`, in either case. It does
not match `G` or `z`.

Character classes can be combined with literals. For example, `status=[0-9]`
matches `status=` followed by one digit. Later lessons will show how
quantifiers match a repeated number of characters.

## Practice

For each task, write a regex character class:

1. Find any single digit in a user ID such as `user42`.
2. Given the request ID `request_id=F02A7C`, find one hexadecimal character.
3. Find the first character of an uppercase log level such as `INFO`, `WARN`,
   or `ERROR`.

Try each pattern against matching and non-matching characters. For the
hexadecimal exercise, include `A`, `f`, `7`, and `G` in your test data.

## Check Your Answers

The patterns are `[0-9]`, `[A-Fa-f0-9]`, and `[A-Z]`, respectively. Each class
matches one character, not an entire word or number.

## Key Takeaways

- Square brackets define a character class.
- A character class matches one character from its set.
- Hyphens define ranges such as `0-9` and `A-Z`.
- A caret immediately after `[` negates the class.
- Character classes become more powerful when combined with literals and
  quantifiers.
