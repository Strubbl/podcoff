package cmd

import (
	"github.com/spf13/cobra"
)

var AddFilterCondition string
var AddFilterField string
var AddFilterKeyword string
var AddFilterPodcastName string

// addFilterCmd represents the addFilter command
var addFilterCmd = &cobra.Command{
	Use:   "addFilter",
	Short: "add filter to podcast",
	Long:  `adds a filter to a selected podcast`,
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {
	addFilterCmd.Flags().StringVarP(&AddFilterCondition, "condition", "", "NOT", "condition, currently only IN or NOT")
	addFilterCmd.Flags().StringVarP(&AddFilterField, "field", "", "title", "which field of the feed item, currently only title")
	addFilterCmd.Flags().StringVarP(&AddFilterKeyword, "keyword", "", "", "keyword to search for in the field")
	addFilterCmd.Flags().StringVarP(&AddFilterPodcastName, "podcast", "", "", "name of the podcast the filter shall be added to")
	addFilterCmd.MarkFlagRequired("keyword")
	addFilterCmd.MarkFlagRequired("podcast")
	rootCmd.AddCommand(addFilterCmd)
}
