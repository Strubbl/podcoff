package cmd

import (
	"github.com/spf13/cobra"
)

// MarkSkippedPodcastName represents the name of the podcasts which shall be marked as skipped
var MarkSkippedPodcastName string

// markSkippedCmd represents the markSkipped command
var markSkippedCmd = &cobra.Command{
	Use:   "markSkipped",
	Short: "Mark podcasts items as skipped for download",
	Long: `Mark all podcasts items, which are fresh, as 
	skipped for download.`,
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {
	markSkippedCmd.Flags().StringVarP(&MarkSkippedPodcastName, "podcast", "", "", "name of the podcast you want to mark as skipped")
	markSkippedCmd.MarkFlagRequired("podcast")
	rootCmd.AddCommand(markSkippedCmd)
}
