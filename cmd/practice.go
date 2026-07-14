/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

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
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("practice called")
	},
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
