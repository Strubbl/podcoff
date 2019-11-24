package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"podcoff/cmd"
	"strings"
	"sync"
	"syscall"
)

func downloadItems(p Podcast, c Configuration) error {
	pis, err := loadPodcastItems(p, c)
	if err != nil {
		return err
	}
	for i := 0; i < len(pis); i++ {
		if pis[i].Status == FRESH {
			err = downloadPodcastItem(pis[i].Link, p, c)
			if err != nil {
				fmt.Println("Error downloading in podcast", p.Name, "the item", pis[i].Link, ":", err)
			}
		}
	}
	return nil
}

func downloadPodcastItem(url string, p Podcast, c Configuration) error {
	var command string
	var downloadHandler string
	var args []string

	downloadFolder := c.DownloadsPath + "/" + p.Name
	createDirIfNotExists(downloadFolder)
	if p.DownloadHandler == "" {
		downloadHandler = defaultDownloadHandler
	} else {
		downloadHandler = p.DownloadHandler
	}
	downloadHandlerParams := strings.Fields(downloadHandler)
	if len(downloadHandlerParams) > 1 {
		args = append(downloadHandlerParams[1:], url)
		command = downloadHandlerParams[0]
	} else {
		args = append(args, url)
		command = downloadHandler
	}
	binary, err := exec.LookPath(command)
	if err != nil {
		return err
	}
	env := os.Environ()
	if cmd.Debug {
		log.Println("downloadPodcastItem: binary is", binary)
		log.Println("downloadPodcastItem: args are", args)
	}
	err = syscall.Exec(binary, args, env)
	if err != nil {
		return err
	}
	return nil
}

// https://stackoverflow.com/a/20438245/709697
func exe_cmd(cmd string, wg *sync.WaitGroup) {
	if cmd.Debug {
		fmt.Println("command is ", cmd)
	}
	// splitting head => g++ parts => rest of the command
	parts := strings.Fields(cmd)
	head := parts[0]
	parts = parts[1:len(parts)]

	out, err := exec.Command(head, parts...).Output()
	if err != nil {
		fmt.Printf("%s", err)
	}
	fmt.Printf("%s", out)
	wg.Done() // Need to signal to waitgroup that this goroutine is done
}
