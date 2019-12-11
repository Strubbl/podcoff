package podcoff

import (
	"errors"
	"fmt"
	"os/exec"
	"strings"
	"sync"

	"github.com/strubbl/podcoff/cmd"
)

func downloadItems(p Podcast, c Configuration) error {
	pis, err := loadPodcastItems(p, c)
	if err != nil {
		return err
	}
	for i := 0; i < len(pis); i++ {
		if pis[i].Status == FRESH {
			filterMatched := doesFilterMatch(pis[i], p.Filter)
			if filterMatched {
				err = downloadPodcastItem(pis[i], p, c)
				if err != nil {
					fmt.Println("Error downloading in podcast", p.Name, "the item", pis[i].Link, ":", err)
					pis[i].Status = FAIL
				} else {
					pis[i].Status = SUCCESS
				}
			} else {
				fmt.Println("Filter prevents downloading", p.Name, pis[i].Title, pis[i].Link, "--> skipped")
				pis[i].Status = SKIPPED
			}
			err = savePodcastItems(pis, p, c)
			if err != nil {
				return err
			}
		}
	}
	return err
}

func downloadPodcastItem(item PodcastItem, p Podcast, c Configuration) error {
	fmt.Println("Downloading", p.Name, item.Title, item.Link)
	downloadFolder := c.DownloadsPath + "/" + p.Name
	createDirIfNotExists(downloadFolder)
	var downloadHandler string
	if p.DownloadHandler == "" {
		downloadHandler = c.DownloadHandler
	} else {
		downloadHandler = p.DownloadHandler
	}
	if downloadHandler == "" {
		return errors.New("No download handler defined to the podcast or in the config")
	}
	command := downloadHandler + " " + item.Link
	var wg sync.WaitGroup
	wg.Add(1)
	out, err := exe_cmd(downloadFolder, command, &wg)
	if cmd.Debug {
		if out != "" {
			fmt.Println("downloadPodcastItem: command output:\n", out)
		}
	}
	return err
}

// based on https://stackoverflow.com/a/20438245/709697
func exe_cmd(directory string, command string, wg *sync.WaitGroup) (string, error) {
	if cmd.Debug {
		fmt.Println("exe_cmd: command is", command)
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
