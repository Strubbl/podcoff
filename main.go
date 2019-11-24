package main

import (
	"fmt"
	"podcoff/cmd"
)

func main() {
	fmt.Println("Hello World!")
	handleFlags()
	c, err := loadConfig()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(c.DownloadHandler)

	podcasts, err := loadPodcasts(c)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(podcasts)

	p1, err := getPodcast("gronkh", "https://www.youtube.com/feeds/videos.xml?channel_id=UCYJ61XIK64sp6ZFFS8sctxw")
	podcasts, err = addPostcast(podcasts, p1)
	if err != nil {
		fmt.Println(err)
	} else {
		savePodcasts(podcasts, c)
	}

	cmd.Execute()

}
