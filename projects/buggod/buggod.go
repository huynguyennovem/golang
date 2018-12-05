package main

import (
	"flag"
	"fmt"
	"os"
)

type File struct {
	name     string
	path     string
	format   string
	date     string
	capacity float32
}

func downloadFile(downloadChan chan File) {
	var file File
	downloadChan <- file
}

func downloadCompleted(file File) {

}

func main() {

	url := flag.String("", "", "Desinate url (Required)")
	flag.Parse()

	if flag.Parsed() == false {
		flag.PrintDefaults()
		os.Exit(1)
	}

	if *url == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	fmt.Println("Downloading ", *url, "...")

	downloadChannel := make(chan File)
	go downloadFile(downloadChannel)
	file := <- downloadChannel
	defer downloadCompleted(file)
}
