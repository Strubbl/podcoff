package cmd

import (
	"github.com/spf13/cobra"
)

var Check bool

// checkCmd represents the check command
var checkCmd = &cobra.Command{
	Use:   "check",
	Short: "check all podcast feeds",
	Long:  `fetch feeds and check for new podcasts, but do not download them`,
	Run: func(cmd *cobra.Command, args []string) {
		Check = true
	},
}

func init() {
	rootCmd.AddCommand(checkCmd)
}
