package cmd

import (
	"github.com/spf13/cobra"
)

// AddFilterCondition is the condition for the new filter
var AddFilterCondition string

// AddFilterField is the field of the feed, which the new filter shall be applied on
var AddFilterField string

// AddFilterKeyword is the search keyword for the new filter
var AddFilterKeyword string

// AddFilterPodcastName is the name of the podcast for which the new filter is
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
	addFilterCmd.Flags().StringVarP(&AddFilterKeyword, "keyword", "", "", "keyword to search for in the field (not case-sensitive")
	addFilterCmd.Flags().StringVarP(&AddFilterPodcastName, "podcast", "", "", "name of the podcast the filter shall be added to")
	addFilterCmd.MarkFlagRequired("keyword")
	addFilterCmd.MarkFlagRequired("podcast")
	rootCmd.AddCommand(addFilterCmd)
}
