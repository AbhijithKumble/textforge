/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/AbhijithKumble/textforge/internal/progress"
	"github.com/spf13/cobra"
)

// progressCmd represents the progress command
var progressCmd = &cobra.Command{
	Use:   "progress",
	Short: "Check your learning progress",
	Long: `Display your current learning statistics and completed exercises.

    Shows a breakdown of your progress across different courses (Regex, grep, sed, etc.)
    and helps you resume exactly where you left off.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		lessons, err := loadLessons()
		if err != nil {
			return err
		}
		store, err := progress.Load(progressFile)
		if err != nil {
			return err
		}

		completed, total := 0, 0
		for _, lesson := range lessons {
			for _, exercise := range lesson.Exercises {
				total++
				if store.IsComplete(lesson.ID + "/" + exercise.ID) {
					completed++
				}
			}
		}

		fmt.Fprintf(cmd.OutOrStdout(), "Progress: %d/%d exercises complete\n", completed, total)
		for _, lesson := range lessons {
			lessonCompleted := 0
			for _, exercise := range lesson.Exercises {
				if store.IsComplete(lesson.ID + "/" + exercise.ID) {
					lessonCompleted++
				}
			}
			fmt.Fprintf(cmd.OutOrStdout(), "  %-24s %d/%d\n", lesson.ID, lessonCompleted, len(lesson.Exercises))
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(progressCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// progressCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// progressCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
