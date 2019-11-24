package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var AddName string
var AddFeedURL string

const numOfAddArgsAllowed = 2

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add [name] [feedURL]",
	Short: "Add podcast",
	Long: `Add a new podcast to the database

You need to give a name and the url in that order. The name should be one word
and consist of characters, that you can use as file name on your file system.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == numOfAddArgsAllowed {
			AddName = args[0]
			AddFeedURL = args[1]
		} else {
			fmt.Println("invalid number of arguments, got", len(args), "but expected", numOfAddArgsAllowed)
			cmd.Help()
		}
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
