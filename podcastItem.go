package podcoff

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

const jsonFileEnding = ".json"

type PodcastItemDownloadStatus int

const (
	FRESH = iota
	SUCCESS
	FAIL
	SKIPPED
)

// Podcast represents one podcast item with its metadata
type PodcastItem struct {
	Link   string
	Status int
	Title  string
}

func loadPodcastItems(p Podcast, c Configuration) ([]PodcastItem, error) {
	var pis []PodcastItem

	podcastItemDataPath := c.MetadataPath + "/" + p.Name + jsonFileEnding
	if _, err := os.Stat(podcastItemDataPath); os.IsNotExist(err) {
		// return nil as error cause it's okay to have no podcast database, so
		// we just start with an empty one
		return pis, nil
	}
	file, err := os.Open(podcastItemDataPath)
	if err != nil {
		return pis, err
	}
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&pis)
	if err != nil {
		return pis, err
	}
	return pis, nil
}

func savePodcastItems(pis []PodcastItem, p Podcast, c Configuration) error {
	b, err := json.MarshalIndent(pis, "", "	")
	if err != nil {
		return err
	}
	podcastItemDataPath := c.MetadataPath + "/" + p.Name + jsonFileEnding
	err = ioutil.WriteFile(podcastItemDataPath, b, 0644)
	if err != nil {
		return err
	}
	return nil
}
