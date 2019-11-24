package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var AddFilterCondition string
var AddFilterKeyword string
var AddFilterField string

// addFilterCmd represents the addFilter command
var addFilterCmd = &cobra.Command{
	Use:   "addFilter",
	Short: "adds a filter to selected feeds",
	Long:  `adds a filter to selected feeds`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("addFilter called")
	},
}

func init() {
	addCmd.Flags().StringVarP(&AddFilterCondition, "condition", "", "", "condition, currently only IN or NOT")
	addCmd.Flags().StringVarP(&AddFilterKeyword, "keyword", "", "", "keyword to search for in the field")
	addCmd.Flags().StringVarP(&AddFilterField, "field", "", "", "which field of the feed item, currently only title")
	rootCmd.AddCommand(addFilterCmd)
}
