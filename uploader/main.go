package main

import (
	"flag"
	"fmt"
	"log"
	yt "spamtube/uploader/youtube"
)

var (
	apikey      = flag.String("apikey", "", "Youtube API Key")
	filename    = flag.String("filename", "", "Name of video file to upload")
	title       = flag.String("title", "", "Video title")
	description = "Test Description"
	category    = "22"
	keywords    = "spamtube"
	privacy     = "listed"
)

func main() {
	flag.Parse()

	if *apikey == "" {
		panic("You must provide a apikey")
	}

	if *filename == "" {
		panic("You must provide a filename of a video file to upload")
	}

	fmt.Println("Upload file " + *filename)
	id, err := yt.UploadToYoutube(*apikey, *filename, *title, description, category, privacy, keywords)
	if err != nil {
		log.Fatalln(err)
		return
	}
	fmt.Println("Youtube Id: " + id)
	// Push id to db

}
