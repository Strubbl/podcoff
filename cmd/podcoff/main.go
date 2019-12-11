package main

import (
	"fmt"
	"log"
	"os"

	"github.com/strubbl/podcoff"
	"github.com/strubbl/podcoff/cmd"
)

func main() {
	log.SetOutput(os.Stdout)
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	cmd.Execute()

	// root command flags
	// check for add command used
	if cmd.Version {
		fmt.Println(podcoff.Version)
		os.Exit(0)
	}

	p := &podcoff.Podcoff{}
	// subcommands
	// for all following command blocks we need the initialized podcoff instance
	if cmd.Debug {
		log.Println("Detected debug flag, activating it in Podcoff")
		(*p).Debug = true
		(*p).Verbose = true
	}
	if cmd.Verbose {
		log.Println("Verbose output")
		(*p).Verbose = true
	}
	err := p.Init(cmd.ConfigJSON)
	if err != nil {
		log.Fatal(err)
	}

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
		err = p.AddPostcast(cmd.AddName, cmd.AddFeedURL)
		if err != nil {
			log.Fatal(err)
		} else {
			err = p.SavePodcasts()
			if err != nil {
				log.Fatal(err)
			}
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
		err := p.CheckPodcasts()
		if err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
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
		err := p.DownloadPodcasts()
		if err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}

	//           _     _    __ _ _ _
	//  __ _  __| | __| |  / _(_) | |_ ___ _ __
	// / _` |/ _` |/ _` | | |_| | | __/ _ \ '__|
	//| (_| | (_| | (_| | |  _| | | ||  __/ |
	// \__,_|\__,_|\__,_| |_| |_|_|\__\___|_|
	//
	if cmd.AddFilterKeyword != "" && cmd.AddFilterPodcastName != "" {
		if cmd.Debug {
			log.Println("found addFilter flag")
		}
		err := p.AddFilter(cmd.AddFilterCondition, cmd.AddFilterField, cmd.AddFilterKeyword, cmd.AddFilterPodcastName)
		if err != nil {
			log.Fatal(err)
		}
		err = p.SavePodcasts()
		if err != nil {
			fmt.Println("Error while saving podcasts config file:", err)
			os.Exit(1)
		}
		fmt.Println("Added filter to podcast", cmd.AddFilterPodcastName)
		os.Exit(0)

	}

	if cmd.Debug {
		log.Println("Finish program")
	}
}
