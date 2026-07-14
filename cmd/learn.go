/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

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
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("learn called")
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
