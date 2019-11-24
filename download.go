package main

import (
	"log"
	"os"
	"podcoff/cmd"
	"strings"
	"syscall"
)

func downloadItems(p Podcast, c Configuration) error {
	pis, err := loadPodcastItems(p, c)
	if err != nil {
		return err
	}
	for i := 0; i < len(pis); i++ {
		if pis[i].Status == FRESH {
			downloadPodcastItem(pis[i].Link, p, c)
		}
	}
	return nil
}

func downloadPodcastItem(url string, p Podcast, c Configuration) error {
	var binary string
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
		binary = downloadHandlerParams[0]
	} else {
		args = append(args, url)
		binary = downloadHandler
	}
	if cmd.Debug {
		log.Println("downloadPodcastItem: binary is", binary)
		log.Println("downloadPodcastItem: downloadHandlerParams are", downloadHandlerParams)
		log.Println("downloadPodcastItem: args are", args)
	}
	env := os.Environ()
	err := syscall.Exec(binary, args, env)
	if err != nil {
		return err
	}
	return nil
}
