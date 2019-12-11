package podcoff

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
)

// podcast represents one podcast with its metadata, name, url, download handler etc.
type Podcast struct {
	Name            string
	FeedURL         string
	DownloadHandler string
	Filter          Filter
}

func (p *Podcoff) loadPodcasts() error {
	if _, err := os.Stat((*p).Config.DatabasePath); os.IsNotExist(err) {
		// return nil as error cause it's okay to have no podcast database, so
		// we just start with an empty one
		return nil
	}
	file, err := os.Open((*p).Config.DatabasePath)
	if err != nil {
		return err
	}
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&p.Podcasts)
	if err != nil {
		return err
	}
	return nil
}

func (p *Podcoff) SavePodcasts() error {
	b, err := json.MarshalIndent((*p).Podcasts, "", "	")
	if err != nil {
		return err
	}
	err = ioutil.WriteFile((*p).Config.DatabasePath, b, 0644)
	if err != nil {
		return err
	}
	return nil
}

func (p *Podcoff) AddPostcast(name string, url string) error {
	newPodcast, err := createPodcast(name, url)
	if err != nil {
		return err
	}
	err = checkPodcastNameAndFeedNotEmpty(newPodcast.Name, newPodcast.FeedURL)
	if err != nil {
		return err
	}
	podcasts := (*p).Podcasts
	for i := 0; i < len(podcasts); i++ {
		if podcasts[i].Name == newPodcast.Name || podcasts[i].FeedURL == newPodcast.FeedURL {
			return errors.New("addPostcast: a podcast with that name or feed url is already in the database")
		}
	}
	podcasts = append(podcasts, newPodcast)
	return nil
}

func createPodcast(name string, url string) (Podcast, error) {
	var p Podcast
	err := checkPodcastNameAndFeedNotEmpty(name, url)
	if err != nil {
		return p, err
	}
	p.Name = name
	p.FeedURL = url
	return p, nil
}

func checkPodcastNameAndFeedNotEmpty(name string, url string) error {
	if name == "" || url == "" {
		return errors.New("checkPodcastNameAndFeedNotEmpty: name and url shall not be empty")
	}
	return nil
}
