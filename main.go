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

	// root command flags
	// check for add command used
	if cmd.Version {
		fmt.Println(version)
		os.Exit(0)
	}

	// subcommands
	//           _     _    __               _
	//  __ _  __| | __| |  / _| ___  ___  __| |
	// / _` |/ _` |/ _` | | |_ / _ \/ _ \/ _` |
	//| (_| | (_| | (_| | |  _|  __/  __/ (_| |
	// \__,_|\__,_|\__,_| |_|  \___|\___|\__,_|
	//
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

	//      _               _
	//  ___| |__   ___  ___| | __
	// / __| '_ \ / _ \/ __| |/ /
	//| (__| | | |  __/ (__|   <
	// \___|_| |_|\___|\___|_|\_\
	//
	if cmd.Check {
		if cmd.Debug {
			log.Println("found check flag")
		}
		podcasts, err := loadPodcasts(c)
		if err != nil {
			log.Println(err)
		}
		if len(podcasts) <= 0 {
			log.Fatal("You haven't any podcasts added. Nothing to check for")
		}
		for i := 0; i < len(podcasts); i++ {
			checkFeed(podcasts[i], c)
		}
	}

	//     _                     _                 _
	//  __| | _____      ___ __ | | ___   __ _  __| |
	// / _` |/ _ \ \ /\ / / '_ \| |/ _ \ / _` |/ _` |
	//| (_| | (_) \ V  V /| | | | | (_) | (_| | (_| |
	// \__,_|\___/ \_/\_/ |_| |_|_|\___/ \__,_|\__,_|
	//
	if cmd.Download {
		if cmd.Debug {
			log.Println("found download flag")
		}
		podcasts, err := loadPodcasts(c)
		if err != nil {
			log.Println(err)
		}
		if len(podcasts) <= 0 {
			log.Fatal("You haven't any podcasts added. Nothing to download")
		}
		for i := 0; i < len(podcasts); i++ {
			downloadItems(podcasts[i], c)
		}
	}

	if cmd.Debug {
		log.Println("Finish program")
	}
}
