package main

import (
	"fmt"
	"podcoff/cmd"

	"github.com/mmcdole/gofeed"
)

func checkFeed(p Podcast, c Configuration) {
	fmt.Println("Checking", p.Name, p.FeedURL)
	pis, err := loadPodcastItems(p, c)
	if err != nil {
		fmt.Println("error loading podcast items for feed", p.Name, "error is:", err)
	}

	fp := gofeed.NewParser()
	feed, _ := fp.ParseURL(p.FeedURL)
	for i := 0; i < len(feed.Items); i++ {
		if cmd.Verbose {
			fmt.Println(p.Name, "- found link", feed.Items[i].Link)
		}
		isLinkKnown := false
		for k := 0; k < len(pis); k++ {
			if pis[k].Link == feed.Items[i].Link {
				isLinkKnown = true
				continue
			}
		}
		if !isLinkKnown {
			var item PodcastItem
			item.Link = feed.Items[i].Link
			item.Status = FRESH
			item.Title = feed.Items[i].Title
			pis = append(pis, item)
		}
	}
	err = savePodcastItems(pis, p, c)
	if err != nil {
		fmt.Println("error loading podcast items for feed", p.Name, "error is:", err)
	}
}
