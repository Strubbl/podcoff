package podcoff

import (
	"errors"
	"fmt"
	"log"
	"os/exec"
	"strings"
	"sync"

	"github.com/strubbl/podcoff/cmd"
)

func (p *Podcoff) DownloadPodcasts() error {
	podcasts := (*p).Podcasts
	if len(podcasts) <= 0 {
		log.Fatal("You haven't any podcasts added. Nothing to download")
	}
	for i := 0; i < len(podcasts); i++ {
		p.downloadItems(podcasts[i])
	}
	return nil
}

func (p *Podcoff) MarkPodcastsAsSkipped(podcastName string) error {
	var pc Podcast
	for i := 0; i < len(p.Podcasts); i++ {
		if p.Podcasts[i].Name == podcastName {
			pc = p.Podcasts[i]
			break
		}
	}
	pis, err := p.loadPodcastItems(pc)
	if err != nil {
		return err
	}
	// mark all items, which are fresh, as skipped now
	for i := 0; i < len(pis); i++ {
		if pis[i].Status == FRESH {
			pis[i].Status = SKIPPED
		}
	}
	err = p.savePodcastItems(pis, pc)
	if err != nil {
		return err
	}
	return err
}

func (p *Podcoff) downloadItems(pc Podcast) error {
	pis, err := p.loadPodcastItems(pc)
	if err != nil {
		return err
	}
	for i := 0; i < len(pis); i++ {
		if pis[i].Status == FRESH {
			filterMatched := doesFilterMatch(pis[i], pc.Filter)
			if filterMatched {
				err = p.downloadPodcastItem(pis[i], pc)
				if err != nil {
					log.Println("Error downloading in podcast", pc.Name, "the item", pis[i].Link, ":", err)
					pis[i].Status = FAIL
				} else {
					pis[i].Status = SUCCESS
				}
			} else {
				if p.Verbose {
					log.Println("Filter prevents downloading", pc.Name, pis[i].Title, pis[i].Link, "--> skipped")
				}
				pis[i].Status = SKIPPED
			}
			err = p.savePodcastItems(pis, pc)
			if err != nil {
				return err
			}
		}
	}
	return err
}

func (p *Podcoff) downloadPodcastItem(item PodcastItem, pc Podcast) error {
	if p.Verbose {
		log.Println("Downloading", pc.Name, item.Title, item.Link)
	}
	downloadFolder := p.Config.DownloadsPath + "/" + pc.Name
	createDirIfNotExists(downloadFolder)
	var downloadHandler string
	if pc.DownloadHandler == "" {
		downloadHandler = p.Config.DownloadHandler
	} else {
		downloadHandler = pc.DownloadHandler
	}
	if downloadHandler == "" {
		return errors.New("No download handler defined to the podcast or in the config")
	}
	command := downloadHandler + " " + item.Link
	var wg sync.WaitGroup
	wg.Add(1)
	out, err := exe_cmd(downloadFolder, command, &wg)
	if p.Debug && out != "" {
		fmt.Println("downloadPodcastItem: command output:\n", out)
	}
	return err
}

// based on https://stackoverflow.com/a/20438245/709697
func exe_cmd(directory string, command string, wg *sync.WaitGroup) (string, error) {
	if cmd.Debug {
		fmt.Println("exec command is:", command)
	}
	// splitting head => g++ parts => rest of the command
	parts := strings.Fields(command)
	head := parts[0]
	parts = parts[1:len(parts)]

	execCmd := exec.Command(head, parts...)
	execCmd.Dir = directory
	out, err := execCmd.Output()
	wg.Done()
	return string(out), err
}
