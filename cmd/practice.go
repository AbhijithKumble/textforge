/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bufio"
	"fmt"
	"io"
	"reflect"
	"regexp"
	"strings"

	"github.com/AbhijithKumble/textforge/internal/course"
	"github.com/AbhijithKumble/textforge/internal/progress"
	"github.com/spf13/cobra"
)

const (
	redText   = "\033[31m"
	greenText = "\033[32m"
	resetText = "\033[0m"
)

// practiceCmd represents the practice command
var practiceCmd = &cobra.Command{
	Use:   "practice",
	Short: "Run interactive exercises to test your skills",
	Long: `Launch the interactive exercise runner to test your skills (similar to Rustlings).

    You will be presented with real-world problems and data. Work on the exercises
    locally, and TextForge will automatically validate your solutions against
    hidden test cases, providing immediate feedback and hints when you're stuck.`,
	Args: cobra.MaximumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		lessons, err := loadLessons()
		if err != nil {
			return err
		}
		store, err := progress.Load(progressFile)
		if err != nil {
			return err
		}

		reader := bufio.NewReader(cmd.InOrStdin())
		lessonCount := 0
		for lessonIndex, lesson := range lessons {
			if len(args) == 1 && lesson.ID != args[0] {
				continue
			}

			unfinished := false
			for _, exercise := range lesson.Exercises {
				progressID := lesson.ID + "/" + exercise.ID
				if store.IsComplete(progressID) {
					continue
				}
				unfinished = true
				break
			}
			if !unfinished {
				continue
			}

			lessonCount++
			fmt.Fprintf(cmd.OutOrStdout(), "\nLesson %d/%d: %s\n", lessonIndex+1, len(lessons), lesson.Title)
			fmt.Fprintln(cmd.OutOrStdout(), "────────────────────────────────────────")

			for exerciseIndex, exercise := range lesson.Exercises {
				progressID := lesson.ID + "/" + exercise.ID
				if store.IsComplete(progressID) {
					continue
				}

				for {
					fmt.Fprintf(cmd.OutOrStdout(), "\nExercise %d/%d\n%s\n", exerciseIndex+1, len(lesson.Exercises), exercise.Prompt)
					if exercise.TestInput != "" {
						fmt.Fprintln(cmd.OutOrStdout(), "\nTest input:")
						for _, line := range strings.Split(exercise.TestInput, "\n") {
							fmt.Fprintf(cmd.OutOrStdout(), "  %s\n", line)
						}
					}
					fmt.Fprint(cmd.OutOrStdout(), "\nEnter a regex (or q to quit): ")
					answer, err := readAnswer(reader)
					if err != nil {
						if err == io.EOF {
							return nil
						}
						return fmt.Errorf("read answer: %w", err)
					}
					answer = strings.TrimSpace(answer)
					if strings.EqualFold(answer, "q") || strings.EqualFold(answer, "quit") {
						fmt.Fprintln(cmd.OutOrStdout(), "Leaving practice. Your saved progress is kept.")
						return nil
					}
					if !validAnswer(answer, exercise) {
						fmt.Fprintln(cmd.OutOrStdout(), "Not quite. Try again.")
						if exercise.TestInput != "" {
							printRegexFeedback(cmd.OutOrStdout(), answer, exercise)
						}
						continue
					}

					if exercise.TestInput != "" {
						if exercise.Captures {
							printCaptureSelection(cmd.OutOrStdout(), answer, exercise.TestInput)
						} else {
							printRegexSelection(cmd.OutOrStdout(), answer, exercise.TestInput)
						}
					}
					store.Complete(progressID)
					if err := store.Save(progressFile); err != nil {
						return err
					}
					fmt.Fprintln(cmd.OutOrStdout(), "Correct! Progress saved.")
					fmt.Fprintln(cmd.OutOrStdout())
					break
				}
			}

			fmt.Fprintf(cmd.OutOrStdout(), "Lesson complete: %s!\n", lesson.Title)
			if nextLessonExists(lessons, lessonIndex, args, store) {
				fmt.Fprint(cmd.OutOrStdout(), "Press Enter for the next lesson, or q to quit: ")
				choice, err := readAnswer(reader)
				if err != nil || strings.EqualFold(strings.TrimSpace(choice), "q") || strings.EqualFold(strings.TrimSpace(choice), "quit") {
					fmt.Fprintln(cmd.OutOrStdout(), "Great work! See you next time.")
					return nil
				}
			}
		}

		if len(args) == 1 {
			if lessonCount == 0 {
				return fmt.Errorf("no unfinished exercises found for lesson %q", args[0])
			}
			return nil
		}
		if lessonCount == 0 {
			fmt.Fprintln(cmd.OutOrStdout(), "All lessons are complete. Great work!")
		}
		return nil
	},
}

// validAnswer compares regex behavior when an exercise provides test input.
// Exercises without test input retain exact-answer validation for non-regex
// lessons and answers that are not intended to be interpreted as patterns.
func validAnswer(answer string, exercise course.Exercise) bool {
	if exercise.TestInput == "" {
		return answer == exercise.Answer
	}

	if exercise.Captures {
		expected, err := regexSubmatches(exercise.Answer, exercise.TestInput)
		if err != nil {
			return false
		}
		candidate, err := regexSubmatches(answer, exercise.TestInput)
		if err != nil {
			return false
		}
		return reflect.DeepEqual(expected, candidate)
	}

	expected, err := regexOutput(exercise.Answer, exercise.TestInput)
	if err != nil {
		return false
	}
	candidate, err := regexOutput(answer, exercise.TestInput)
	if err != nil {
		return false
	}

	return strings.Join(expected, "\x00") == strings.Join(candidate, "\x00")
}

func regexSubmatches(pattern, input string) ([][]string, error) {
	re, err := regexp.Compile(pattern)
	if err != nil {
		return nil, err
	}

	var matches [][]string
	for _, line := range strings.Split(input, "\n") {
		matches = append(matches, re.FindAllStringSubmatch(line, -1)...)
	}
	return matches, nil
}

func regexOutput(pattern, input string) ([]string, error) {
	re, err := regexp.Compile(pattern)
	if err != nil {
		return nil, err
	}

	// grep evaluates patterns one line at a time, so process multiline test
	// input line-by-line to give ^ and $ the expected behavior.
	var matches []string
	for _, line := range strings.Split(input, "\n") {
		matches = append(matches, re.FindAllString(line, -1)...)
	}
	return matches, nil
}

func printRegexFeedback(out io.Writer, answer string, exercise course.Exercise) {
	if exercise.Captures {
		printCaptureFeedback(out, answer, exercise)
		return
	}

	selected, err := regexOutput(answer, exercise.TestInput)
	if err != nil {
		fmt.Fprintf(out, "Your regex is invalid: %v\n", err)
		return
	}

	expected, err := regexOutput(exercise.Answer, exercise.TestInput)
	if err != nil {
		return
	}
	selectedInput, err := highlightRegex(answer, exercise.TestInput)
	if err != nil {
		return
	}
	expectedInput, err := highlightRegexColor(exercise.Answer, exercise.TestInput, greenText)
	if err != nil {
		return
	}
	fmt.Fprintln(out, "\nYOUR RESULT")
	fmt.Fprintf(out, "  Pattern: %s\n", answer)
	fmt.Fprintf(out, "  Matches found: %d %s\n", len(selected), formatMatches(selected))
	if len(selected) == 0 {
		fmt.Fprintln(out, "  Matched input: (none)")
	} else {
		fmt.Fprintln(out, "  Matched input (red):")
		printIndented(out, selectedInput)
	}
	fmt.Fprintln(out, "\nEXPECTED RESULT")
	fmt.Fprintf(out, "  Matches expected: %d %s\n", len(expected), formatMatches(expected))
	fmt.Fprintln(out, "  Expected input (green):")
	printIndented(out, expectedInput)
	fmt.Fprintf(out, "\n  Difference: your regex found %d match(es); %d expected.\n", len(selected), len(expected))
}

func printCaptureFeedback(out io.Writer, answer string, exercise course.Exercise) {
	actual, err := regexSubmatches(answer, exercise.TestInput)
	if err != nil {
		fmt.Fprintf(out, "Your regex is invalid: %v\n", err)
		return
	}
	expected, err := regexSubmatches(exercise.Answer, exercise.TestInput)
	if err != nil {
		return
	}

	fmt.Fprintln(out, "\nYOUR RESULT")
	fmt.Fprintf(out, "  Pattern: %s\n", answer)
	printCaptureRows(out, "  ", actual)
	fmt.Fprintln(out, "\nEXPECTED RESULT")
	printCaptureRows(out, "  ", expected)
	fmt.Fprintf(out, "\n  Difference: your regex found %d match(es); %d expected.\n", len(actual), len(expected))
}

func printCaptureRows(out io.Writer, prefix string, matches [][]string) {
	if len(matches) == 0 {
		fmt.Fprintf(out, "%sMatches: (none)\n", prefix)
		return
	}
	fmt.Fprintf(out, "%sMatches: %d\n", prefix, len(matches))
	for index, match := range matches {
		fmt.Fprintf(out, "%s  Match %d: %s\n", prefix, index+1, match[0])
		if len(match) == 1 {
			fmt.Fprintf(out, "%s    Groups: (none)\n", prefix)
			continue
		}
		for groupIndex, group := range match[1:] {
			fmt.Fprintf(out, "%s    Group %d: %s\n", prefix, groupIndex+1, group)
		}
	}
}

func printCaptureSelection(out io.Writer, pattern, input string) {
	matches, err := regexSubmatches(pattern, input)
	if err != nil {
		return
	}
	fmt.Fprintln(out, "Matches and captures:")
	printCaptureRows(out, "  ", matches)
}

func printRegexSelection(out io.Writer, pattern, input string) {
	selected, err := regexOutput(pattern, input)
	if err != nil {
		return
	}
	highlighted, err := highlightRegex(pattern, input)
	if err != nil {
		return
	}
	fmt.Fprintln(out, "Matches (red):")
	printIndented(out, highlighted)
	fmt.Fprintf(out, "Valid characters selected: %s\n", formatMatches(selected))
}

func formatMatches(matches []string) string {
	if len(matches) == 0 {
		return "(none)"
	}
	var builder strings.Builder
	builder.WriteString("[")
	for i, m := range matches {
		if i > 0 {
			builder.WriteString(" ")
		}
		if strings.HasPrefix(m, "\"") && strings.HasSuffix(m, "\"") && len(m) >= 2 {
			builder.WriteString(m)
		} else {
			builder.WriteString(fmt.Sprintf("%q", m))
		}
	}
	builder.WriteString("]")
	return builder.String()
}

func printIndented(out io.Writer, text string) {
	for _, line := range strings.Split(text, "\n") {
		fmt.Fprintf(out, "  %s\n", line)
	}
}

func highlightRegex(pattern, input string) (string, error) {
	return highlightRegexColor(pattern, input, redText)
}

func highlightRegexColor(pattern, input, color string) (string, error) {
	re, err := regexp.Compile(pattern)
	if err != nil {
		return "", err
	}

	lines := strings.Split(input, "\n")
	for lineIndex, line := range lines {
		indices := re.FindAllStringIndex(line, -1)
		if len(indices) == 0 {
			continue
		}

		var highlighted strings.Builder
		cursor := 0
		for _, match := range indices {
			highlighted.WriteString(line[cursor:match[0]])
			highlighted.WriteString(color)
			highlighted.WriteString(line[match[0]:match[1]])
			highlighted.WriteString(resetText)
			cursor = match[1]
		}
		highlighted.WriteString(line[cursor:])
		lines[lineIndex] = highlighted.String()
	}
	return strings.Join(lines, "\n"), nil
}

func readAnswer(reader *bufio.Reader) (string, error) {
	answer, err := reader.ReadString('\n')
	if err != nil && len(answer) == 0 {
		return "", err
	}
	return answer, nil
}

func nextLessonExists(lessons []*course.Lesson, current int, args []string, store *progress.Store) bool {
	for i := current + 1; i < len(lessons); i++ {
		if len(args) == 1 && lessons[i].ID != args[0] {
			continue
		}
		for _, exercise := range lessons[i].Exercises {
			if !store.IsComplete(lessons[i].ID + "/" + exercise.ID) {
				return true
			}
		}
	}
	return false
}

func init() {
	rootCmd.AddCommand(practiceCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// practiceCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// practiceCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
