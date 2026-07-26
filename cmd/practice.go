/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bufio"
	"fmt"
	"io"
	"regexp"
	"strings"

	"github.com/AbhijithKumble/textforge/internal/course"
	"github.com/AbhijithKumble/textforge/internal/progress"
	"github.com/spf13/cobra"
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
						fmt.Fprintf(cmd.OutOrStdout(), "\nTest input:\n  %s\n", exercise.TestInput)
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
						printRegexSelection(cmd.OutOrStdout(), answer, exercise.TestInput)
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

func regexOutput(pattern, input string) ([]string, error) {
	re, err := regexp.Compile(pattern)
	if err != nil {
		return nil, err
	}
	return re.FindAllString(input, -1), nil
}

func printRegexFeedback(out io.Writer, answer string, exercise course.Exercise) {
	selected, err := regexOutput(answer, exercise.TestInput)
	if err != nil {
		fmt.Fprintf(out, "Your regex is invalid: %v\n", err)
		return
	}

	expected, err := regexOutput(exercise.Answer, exercise.TestInput)
	if err != nil {
		return
	}
	fmt.Fprintf(out, "Your regex selected: %q\n", selected)
	fmt.Fprintf(out, "Expected selection: %q\n", expected)
}

func printRegexSelection(out io.Writer, pattern, input string) {
	selected, err := regexOutput(pattern, input)
	if err != nil {
		return
	}
	fmt.Fprintf(out, "Valid characters selected: %q\n", selected)
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
