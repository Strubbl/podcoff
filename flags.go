package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

const defaultConfigJSON = "config.json"

var debug = flag.Bool("d", false, "get debug output (implies verbose mode)")
var h = flag.Bool("h", false, "print help")
var v = flag.Bool("v", false, "print version")
var verbose = flag.Bool("verbose", false, "verbose mode")

var configJSON = flag.String("settings", defaultConfigJSON, "file name of config JSON file")

func usage() {
	fmt.Fprintf(os.Stderr, "Usage: %s [OPTION]... URL\n", os.Args[0])
	fmt.Fprintf(os.Stderr, "Try '%s -help' for more information.\n", os.Args[0])
}

func helptext() {
	fmt.Fprintf(os.Stderr, "Usage: %s [OPTION]\n", os.Args[0])
	flag.PrintDefaults()
	fmt.Println("<none yet>")
}

func handleFlags() {
	flag.Parse()
	if *debug && len(flag.Args()) > 0 {
	}
	if *h {
		helptext()
		os.Exit(0)
	}

	// version first, because it directly exits here
	if *v {
		fmt.Printf("version %v\n", version)
		os.Exit(0)
	}
	// test verbose before debug because debug implies verbose
	if *verbose && !*debug {
		log.Printf("verbose mode")
	}
	if *debug {
		log.Printf("handleFlags: debug mode")
		// debug implies verbose
		*verbose = true
	}
}
