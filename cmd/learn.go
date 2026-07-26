/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

// learnCmd represents the learn command
var learnCmd = &cobra.Command{
	Use:   "learn",
	Short: "Start or resume interactive learning lessons",
	Long: `Start or resume interactive text-processing lessons.

    This command loads structured lessons from the courses directory, explains
    concepts in detail, and lets you navigate step-by-step through topics like
    regex literals, character classes, and shell pipelines.`,
	Args: cobra.MaximumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		lessons, err := loadLessons()
		if err != nil {
			return err
		}
		if len(args) == 0 {
			out := cmd.OutOrStdout()
			fmt.Fprintln(out, "TextForge Learning")
			fmt.Fprintln(out, "=================")
			fmt.Fprintf(out, "\n%d lessons available\n\n", len(lessons))
			for index, lesson := range lessons {
				fmt.Fprintf(out, "[%d] %s\n", index+1, lesson.Title)
				fmt.Fprintf(out, "    %s · %d exercises\n", lesson.Difficulty, len(lesson.Exercises))
				fmt.Fprintf(out, "    %s\n", lesson.Description)
				fmt.Fprintf(out, "    ID: %s\n\n", lesson.ID)
			}
			fmt.Fprintln(out, "Start learning:")
			fmt.Fprintln(out, "  textforge learn <lesson-id>")
			fmt.Fprintln(out, "  textforge practice <lesson-id>")
			return nil
		}

		for _, lesson := range lessons {
			if lesson.ID == args[0] {
				out := cmd.OutOrStdout()
				fmt.Fprintf(out, "Lesson %d · %s\n", lesson.Index, lesson.Title)
				fmt.Fprintf(out, "%s · %d exercises\n", lesson.Difficulty, len(lesson.Exercises))
				fmt.Fprintln(out, "────────────────────────────────────────")
				fmt.Fprintf(out, "\n%s\n\n", lesson.Description)
				fmt.Fprintln(out, strings.TrimSpace(lesson.Content))
				fmt.Fprintln(out, "\n────────────────────────────────────────")
				fmt.Fprintf(out, "Ready to practice?  textforge practice %s\n", lesson.ID)
				return nil
			}
		}
		return fmt.Errorf("lesson %q not found", args[0])
	},
}

func init() {
	rootCmd.AddCommand(learnCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// learnCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// learnCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
