package course

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestLoadLesson(t *testing.T) {
	path := writeLesson(t, `---
id: regex-01-literals
title: Regex Literals
index: 1
description: Match exact text.
difficulty: Beginner
---

# Regex Literals

Match plain text exactly.
`)

	lesson, err := LoadLesson(path)
	if err != nil {
		t.Fatalf("LoadLesson() returned an unexpected error: %v", err)
	}

	if lesson.ID != "regex-01-literals" {
		t.Errorf("ID = %q, want %q", lesson.ID, "regex-01-literals")
	}
	if lesson.Title != "Regex Literals" {
		t.Errorf("Title = %q, want %q", lesson.Title, "Regex Literals")
	}
	if lesson.Index != 1 {
		t.Errorf("Index = %d, want 1", lesson.Index)
	}
	if !strings.Contains(lesson.Content, "# Regex Literals") {
		t.Errorf("Content does not contain the lesson heading: %q", lesson.Content)
	}
}

func TestLoadLessonMissingFile(t *testing.T) {
	_, err := LoadLesson(filepath.Join(t.TempDir(), "missing.md"))
	if err == nil {
		t.Fatal("LoadLesson() returned nil error for a missing file")
	}
}

func TestLoadLessonMalformedFrontMatter(t *testing.T) {
	path := writeLesson(t, `---
title: [invalid
---

Lesson content.
`)

	_, err := LoadLesson(path)
	if err == nil {
		t.Fatal("LoadLesson() returned nil error for malformed front matter")
	}
}

func writeLesson(t *testing.T, content string) string {
	t.Helper()

	path := filepath.Join(t.TempDir(), "lesson.md")
	if err := os.WriteFile(path, []byte(content), 0o600); err != nil {
		t.Fatalf("write test lesson: %v", err)
	}

	return path
}
