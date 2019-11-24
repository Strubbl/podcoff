package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var AddName string
var AddFeedURL string

const numOfArgs = 2

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add [name] [feedURL]",
	Short: "Add a podcast",
	Long: `Add a new podcast to the database

You need to give a name and the url in that order.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("add called with the following args:", args)
		fmt.Println("num of args:", len(args))
		if len(args) == numOfArgs {
			AddName = args[0]
			AddFeedURL = args[1]
		} else {
			fmt.Println("invalid number of arguments, got", len(args), "but expected", numOfArgs)
			cmd.Help()
		}
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
