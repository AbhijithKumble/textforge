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
			fmt.Println("Lessons:")
			for _, lesson := range lessons {
				fmt.Printf("  %-24s %s (%s)\n", lesson.ID, lesson.Title, lesson.Difficulty)
			}
			fmt.Println("\nRun `textforge learn <lesson-id>` to read a lesson.")
			return nil
		}

		for _, lesson := range lessons {
			if lesson.ID == args[0] {
				fmt.Printf("# %s\n\n%s\n", lesson.Title, strings.TrimSpace(lesson.Content))
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
