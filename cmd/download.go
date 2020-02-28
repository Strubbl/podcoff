package cmd

import (
	"github.com/spf13/cobra"
)

// Download is used as a flag to make podcoff download all new podcasts
var Download bool

// downloadCmd represents the download command
var downloadCmd = &cobra.Command{
	Use:   "download",
	Short: "download new podcasts",
	Long:  `download all new podcasts`,
	Run: func(cmd *cobra.Command, args []string) {
		Download = true
	},
}

func init() {
	rootCmd.AddCommand(downloadCmd)
}
