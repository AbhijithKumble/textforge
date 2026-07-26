package cmd

import (
	"bytes"
	"strings"
	"testing"

	"github.com/AbhijithKumble/textforge/internal/course"
)

func TestValidAnswerAcceptsEquivalentCharacterClass(t *testing.T) {
	exercise := course.Exercise{
		Answer:    "[A-Fa-f0-9]",
		TestInput: "A f 7 G z 0 9",
	}

	if !validAnswer("[A-F0-9a-f]", exercise) {
		t.Fatal("equivalent character class should be accepted")
	}
}

func TestValidAnswerRejectsDifferentMatches(t *testing.T) {
	exercise := course.Exercise{
		Answer:    "[A-Fa-f0-9]",
		TestInput: "A f 7 G z 0 9",
	}

	if validAnswer("[A-Z]", exercise) {
		t.Fatal("different character class should be rejected")
	}
}

func TestValidAnswerRejectsInvalidRegex(t *testing.T) {
	exercise := course.Exercise{
		Answer:    "[A-Fa-f0-9]",
		TestInput: "A f 7 G z 0 9",
	}

	if validAnswer("[", exercise) {
		t.Fatal("invalid regex should be rejected")
	}
}

func TestPrintRegexFeedbackShowsSelections(t *testing.T) {
	exercise := course.Exercise{
		Answer:    "[A-Fa-f0-9]",
		TestInput: "A f 7 G z 0 9",
	}
	var output bytes.Buffer

	printRegexFeedback(&output, "[A-Z]", exercise)

	want := "Your regex selected: [\"A\" \"G\"]\nExpected selection: [\"A\" \"f\" \"7\" \"0\" \"9\"]\n"
	if output.String() != want {
		t.Fatalf("feedback = %q, want %q", output.String(), want)
	}
}

func TestPrintRegexFeedbackShowsInvalidRegex(t *testing.T) {
	exercise := course.Exercise{Answer: "[A-Z]", TestInput: "INFO"}
	var output bytes.Buffer

	printRegexFeedback(&output, "[", exercise)

	if !strings.HasPrefix(output.String(), "Your regex is invalid:") {
		t.Fatalf("feedback = %q, want invalid-regex message", output.String())
	}
}

func TestPrintRegexSelectionShowsValidCharacters(t *testing.T) {
	var output bytes.Buffer

	printRegexSelection(&output, "[A-Fa-f0-9]", "A f 7 G z 0 9")

	want := "Valid characters selected: [\"A\" \"f\" \"7\" \"0\" \"9\"]\n"
	if output.String() != want {
		t.Fatalf("selection = %q, want %q", output.String(), want)
	}
}
