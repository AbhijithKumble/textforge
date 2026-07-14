/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// progressCmd represents the progress command
var progressCmd = &cobra.Command{
	Use:   "progress",
	Short: "Check your learning progress",
	Long: `Display your current learning statistics and completed exercises.

    Shows a breakdown of your progress across different courses (Regex, grep, sed, etc.)
    and helps you resume exactly where you left off.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("progress called")
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
