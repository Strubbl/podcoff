package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World!")
	handleFlags()
	c, err := loadConfig()
	if err != nil {
		fmt.Println(err)
	}
	podcasts, err := loadPodcasts(c)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(podcasts)
}
