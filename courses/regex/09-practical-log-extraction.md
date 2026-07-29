---
id: regex-09-practical-log-extraction
title: "Practical Log & Config Extraction"
index: 9
description: "Combine anchors, classes, quantifiers, and capturing groups for real-world text extraction."
difficulty: "Advanced"
exercises:
  - id: extract-ip-addresses
    prompt: "Extract IPv4 addresses (such as 192.168.1.50) from web server access log lines."
    answer: '[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}'
    test_input: "2026-07-28 22:00:01 192.168.1.50 GET /index.html 200\n2026-07-28 22:00:02 10.0.0.1 POST /login 401\nno IP on this log line"
  - id: extract-http-status
    prompt: "Capture the HTTP method (GET/POST/DELETE) and HTTP response status code (3 digits) from server log entries."
    answer: '\b(GET|POST|PUT|DELETE)\s+/[^\s]*\s+([0-9]{3})\b'
    captures: true
    test_input: "GET /api/v1/users 200\nPOST /login 401\nDELETE /item/5 500\ninvalid request line"
  - id: extract-env-vars
    prompt: "Capture the variable name and assigned value from exported environment lines in shell scripts."
    answer: '^export\s+([A-Za-z_][A-Za-z0-9_]*)=([^ ]+)$'
    captures: true
    test_input: "export PORT=8080\nexport DB_HOST=localhost\n# export IGNORED=1\nexport SECRET_KEY=xyz123"
---

# Practical Log & Config Extraction

## What You Will Learn

By the end of this lesson, you will be able to:

- combine anchors (`^`, `$`, `\b`), character classes (`[0-9]`, `[A-Za-z_]`), and quantifiers (`{1,3}`, `+`);
- extract IP addresses and status codes from production log files; and
- capture configuration keys and values from Unix shell scripts and configuration files.

## The Core Idea

In real-world text processing, you rarely use a single regex feature in isolation. Production logs, CSV datasets, and server configuration files require **combining multiple concepts**:

1. **Anchors** (`^`, `$`, `\b`) ensure matching starts and ends at exact boundaries.
2. **Character Classes** (`[0-9]`, `[A-Za-z_]`) specify allowed characters for each token.
3. **Quantifiers** (`+`, `{n,m}`) define repetition lengths.
4. **Capturing Groups** (`(...)`) extract specific fields for downstream processing.

## Real-World Examples

### 1. IPv4 Address Matching
To match an IP address like `192.168.1.1`:

```text
[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}
```

Each octet consists of 1 to 3 digits separated by literal dots (`\.`).

### 2. HTTP Server Log Extraction
To capture the method (`GET`/`POST`) and response code (`200`/`404`/`500`) from log lines:

```text
\b(GET|POST|PUT|DELETE)\s+/[^\s]*\s+([0-9]{3})\b
```

- `\b(GET|POST|PUT|DELETE)` captures Group 1 (the HTTP method).
- `\s+/[^\s]*\s+` matches spaces, the URL path, and trailing spaces.
- `([0-9]{3})\b` captures Group 2 (the 3-digit status code).

### 3. Shell Export Variables
To capture key-value pairs from lines like `export DB_HOST=localhost`:

```text
^export\s+([A-Za-z_][A-Za-z0-9_]*)=([^ ]+)$
```

- `^export\s+` ensures the line starts with `export` followed by whitespace.
- `([A-Za-z_][A-Za-z0-9_]*)` captures Group 1 (the environment variable name).
- `=` matches the literal assignment operator.
- `([^ ]+)$` captures Group 2 (the value, up to end of line, avoiding trailing spaces).

## Key Takeaways

- Complex regexes build on simple foundational blocks (anchors, classes, quantifiers, groups).
- Always test against negative inputs (comments, malformed lines, unexpected spaces) to prevent false matches.
- Use capturing groups when you need to extract specific fields out of structured log lines.
