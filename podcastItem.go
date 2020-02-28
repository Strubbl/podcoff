package podcoff

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

const jsonFileEnding = ".json"

// PodcastItemDownloadStatus is used as an enum to represent the download status of a podcast item
type PodcastItemDownloadStatus int

// FRESH and all following consts are used as enum for PodcastItemDownloadStatus
const (
	FRESH = iota
	SUCCESS
	FAIL
	SKIPPED
)

// PodcastItem represents one podcast item with its metadata
type PodcastItem struct {
	Link   string
	Status int
	Title  string
}

func (p *Podcoff) loadPodcastItems(pc Podcast) ([]PodcastItem, error) {
	var pis []PodcastItem

	podcastItemDataPath := p.Config.MetadataPath + "/" + pc.Name + jsonFileEnding
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

func (p *Podcoff) savePodcastItems(pis []PodcastItem, pc Podcast) error {
	b, err := json.MarshalIndent(pis, "", "	")
	if err != nil {
		return err
	}
	podcastItemDataPath := p.Config.MetadataPath + "/" + pc.Name + jsonFileEnding
	err = ioutil.WriteFile(podcastItemDataPath, b, 0644)
	if err != nil {
		return err
	}
	return nil
}
