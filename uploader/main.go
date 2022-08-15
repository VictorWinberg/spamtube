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
	description = flag.String("description", "This is video was made by the SpamTube Team", "Description of the video")
	category    = flag.String("category", "22", "Category")
	keywords    = flag.String("keywords", "spamtube,ai,fyp", "Keyword for search")
	privacy     = "public"
)

func main() {
	flag.Parse()

	if *apikey == "" {
		panic("You must provide a apikey")
	}

	if *filename == "" {
		panic("You must provide a filename of a video file to upload")
	}

	if *title == "" {
		panic("You must provide a title")
	}

	fmt.Println("Upload file " + *filename)
	id, err := yt.UploadToYoutube(*apikey, *filename, *title+"#shorts", *description, *category, privacy, *keywords)
	if err != nil {
		log.Fatalln(err)
		return
	}
	fmt.Println("Youtube Id: " + id)
	// Push id to db

}
