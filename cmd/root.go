package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var ConfigJSON string
var Debug bool
var Verbose bool
var Version bool

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "podcoff",
	Short: "podcoff is a cli application to download podcasts.",
	Long: `podcoff is a cli application to download podcasts.

It is inspired by the famous greg application, which is not actively maintained anymore.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//	Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	rootCmd.PersistentFlags().StringVar(&ConfigJSON, "config", "", "config file (default is defaultConfigJSON)")
	rootCmd.PersistentFlags().BoolVarP(&Debug, "debug", "", false, "debug output")
	rootCmd.PersistentFlags().BoolVarP(&Verbose, "verbose", "", false, "verbose output")
	rootCmd.PersistentFlags().BoolVarP(&Version, "version", "v", false, "print version")
}
