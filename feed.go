package main

import (
	"fmt"
	"podcoff/cmd"

	"github.com/mmcdole/gofeed"
)

func checkFeed(p Podcast, c Configuration) {
	if cmd.Verbose {
		fmt.Println("Checking", p.Name)
	}
	pis, err := loadPodcastItems(p, c)
	if err != nil {
		fmt.Println("error loading podcast items for feed", p.Name)
	}

	fp := gofeed.NewParser()
	feed, _ := fp.ParseURL(p.FeedURL)
	for i := 0; i < len(feed.Items); i++ {
		fmt.Println(feed.Items[i].Link)
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
			pis = append(pis, item)
		}
	}
	err = savePodcastItems(pis, p, c)
	if err != nil {
		fmt.Println("error loading podcast items for feed", p.Name)
	}
}
