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

func TestPrintRegexFeedbackShowsSeparateResults(t *testing.T) {
	exercise := course.Exercise{
		Answer:    "[A-Fa-f0-9]",
		TestInput: "A f 7 G z 0 9",
	}
	var output bytes.Buffer

	printRegexFeedback(&output, "[A-Z]", exercise)
	feedback := output.String()
	for _, want := range []string{
		"YOUR RESULT",
		"Pattern: [A-Z]",
		"Matches found: 2 [\"A\" \"G\"]",
		"EXPECTED RESULT",
		"Pattern: [A-Fa-f0-9]",
		"Matches expected: 5 [\"A\" \"f\" \"7\" \"0\" \"9\"]",
		"Difference: your regex found 2 match(es); 5 expected.",
		redText + "A" + resetText,
		greenText + "f" + resetText,
	} {
		if !strings.Contains(feedback, want) {
			t.Errorf("feedback %q does not contain %q", feedback, want)
		}
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

	if !strings.Contains(output.String(), "Valid characters selected: [\"A\" \"f\" \"7\" \"0\" \"9\"]") {
		t.Fatalf("selection = %q, want selected-character output", output.String())
	}
	if !strings.Contains(output.String(), redText+"A"+resetText) {
		t.Fatalf("selection = %q, want red match", output.String())
	}
}

func TestRegexOutputTreatsAnchorsAsLineAnchors(t *testing.T) {
	got, err := regexOutput("^ERROR", "ERROR disk full\nINFO ERROR disk full")
	if err != nil {
		t.Fatalf("regexOutput() returned an error: %v", err)
	}

	if len(got) != 1 || got[0] != "ERROR" {
		t.Fatalf("matches = %q, want [\"ERROR\"]", got)
	}
}
