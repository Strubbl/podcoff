package podcoff

import (
	"errors"
	"log"
	"strings"

	"github.com/mmcdole/gofeed"
)

// CheckPodcasts loads all feeds and checks for new items
func (p *Podcoff) CheckPodcasts() error {
	podcasts := (*p).Podcasts
	if len(podcasts) <= 0 {
		return errors.New("You haven't any podcasts added. Nothing to check for")
	}
	for i := 0; i < len(podcasts); i++ {
		err := (*p).checkFeed(podcasts[i], p.Verbose)
		if err != nil {
			return err
		}
	}
	return nil
}

func (p *Podcoff) checkFeed(pc Podcast, verbose bool) error {
	if p.Verbose {
		log.Println("Checking", pc.Name, pc.FeedURL)
	}
	pis, err := p.loadPodcastItems(pc)
	if err != nil {
		return errors.New(strings.Join([]string{"error loading podcast items for feed", pc.Name, "error is:", err.Error()}, " "))
	}

	fp := gofeed.NewParser()
	feed, _ := fp.ParseURL(pc.FeedURL)
	for i := 0; i < len(feed.Items); i++ {
		if p.Verbose {
			log.Println(pc.Name, "- found link", feed.Items[i].Link)
		}
		isLinkKnown := false
		for k := 0; k < len(pis); k++ {
			if pis[k].Link == feed.Items[i].Link {
				isLinkKnown = true
				break
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
	err = p.savePodcastItems(pis, pc)
	if err != nil {
		return errors.New(strings.Join([]string{"error loading podcast items for feed", pc.Name, "error is:", err.Error()}, " "))
	}
	return nil
}
