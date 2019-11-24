package main

import (
	"fmt"
	"log"
	"os"
	"podcoff/cmd"
)

func main() {
	log.SetOutput(os.Stdout)
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	cmd.Execute()
	c, err := loadConfig(cmd.ConfigJSON)
	if err != nil {
		log.Fatal(err)
	}

	// check for add command used
	if cmd.AddFeedURL != "" && cmd.AddName != "" {
		if cmd.Debug {
			log.Println("found add params", cmd.AddName, cmd.AddFeedURL)
		}
		podcasts, err := loadPodcasts(c)
		if err != nil {
			log.Println(err)
		}
		if cmd.Debug {
			log.Println(podcasts)
		}
		newPodcast, err := getPodcast(cmd.AddName, cmd.AddFeedURL)
		podcasts, err = addPostcast(podcasts, newPodcast)
		if err != nil {
			log.Fatal(err)
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

	// check for check command used
	if cmd.Check {
		if cmd.Debug {
			log.Println("found check flag")
		}
		podcasts, err := loadPodcasts(c)
		if err != nil {
			log.Println(err)
		}
		if len(podcasts) <= 0 {
			log.Fatal("You have no podcasts added. Nothing to check for")
		}

	}
	log.Println("Finish program")
}
