package main

import (
	"fmt"
	"os"
	"podcoff/cmd"
)

func main() {
	cmd.Execute()
	c, err := loadConfig(cmd.ConfigJSON)
	if err != nil {
		fmt.Println(err)
	}

	// check for add command used
	if cmd.AddFeedURL != "" && cmd.AddName != "" {
		fmt.Println("found add params")
		podcasts, err := loadPodcasts(c)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(podcasts)

		newPodcast, err := getPodcast(cmd.AddName, cmd.AddFeedURL)
		podcasts, err = addPostcast(podcasts, newPodcast)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		} else {
			savePodcasts(podcasts, c)
		}
		os.Exit(0)
	}

	// check for add command used
	if cmd.Version {
		fmt.Println(version)
		os.Exit(0)
	}
}
