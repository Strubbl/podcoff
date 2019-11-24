package main

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
)

// Podcast represents one podcast with its metadata, name, url, download handler etc.
type Podcast struct {
	Name            string
	FeedURL         string
	DownloadHandler string
	Filter          []string
}

func loadPodcasts(c Configuration) ([]Podcast, error) {
	var p []Podcast

	if _, err := os.Stat(c.DatabasePath); os.IsNotExist(err) {
		// return nil as error cause it's okay to have no podcast database, so
		// we just start with an empty one
		return p, nil
	}
	file, err := os.Open(c.DatabasePath)
	if err != nil {
		return p, err
	}
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&p)
	if err != nil {
		return p, err
	}
	return p, nil
}

func savePodcasts(p []Podcast, c Configuration) error {
	b, err := json.MarshalIndent(p, "", "	")
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(c.DatabasePath, b, 0644)
	if err != nil {
		return err
	}
	return nil
}

func getPodcast(name string, url string) (Podcast, error) {
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
		return errors.New("getPodcast: name and url shall not be empty")
	}
	return nil
}

func addPostcast(podcasts []Podcast, newPodcast Podcast) ([]Podcast, error) {
	err := checkPodcastNameAndFeedNotEmpty(newPodcast.Name, newPodcast.FeedURL)
	if err != nil {
		return podcasts, err
	}
	for i := 0; i < len(podcasts); i++ {
		if podcasts[i].Name == newPodcast.Name || podcasts[i].FeedURL == newPodcast.FeedURL {
			return podcasts, errors.New("addPostcast: a podcast with that name or feed url is already in the database")
		}
	}
	podcasts = append(podcasts, newPodcast)
	return podcasts, nil
}
