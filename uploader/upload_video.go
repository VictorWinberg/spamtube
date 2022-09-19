package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"google.golang.org/api/youtube/v3"
)

var (
	filename    = flag.String("filename", "", "Name of video file to upload")
	title       = flag.String("title", "Test Title", "Video title")
	description = flag.String("description", "Test Description", "Video description")
	category    = flag.String("category", "22", "Video category")
	keywords    = flag.String("keywords", "", "Comma separated list of video keywords")
	privacy     = flag.String("privacy", "unlisted", "Video privacy status")
)

func main() {
	flag.Parse()

	if *filename == "" {
		panic("You must provide a filename of a video file to upload")
	}

	client := getClient(youtube.YoutubeReadonlyScope, youtube.YoutubeUploadScope)

	service, err := youtube.New(client)
	if err != nil {
		panic(fmt.Sprintf("Error creating YouTube client: %v", err))
	}

	upload := &youtube.Video{
		Snippet: &youtube.VideoSnippet{
			Title:       *title,
			Description: *description,
			CategoryId:  *category,
			Tags:        []string{"spamtube"},
		},
		Status: &youtube.VideoStatus{PrivacyStatus: *privacy},
	}

	// The API returns a 400 Bad Request response if tags is an empty string.
	if strings.Trim(*keywords, "") != "" {
		upload.Snippet.Tags = strings.Split(*keywords, ",")
	}

	channelResponse, err := service.Channels.List([]string{"snippet", "contentDetails", "statistics"}).Mine(true).Do()
	if err != nil {
		panic(fmt.Sprintf("Error calling YouTube client: %v", err))
	}
	fmt.Printf("Channel id: %v\n", channelResponse.Items[0].Id)
	fmt.Printf("Channel name: %v\n", channelResponse.Items[0].Snippet.Title)

	// TODO: Remove this when you want to upload
	if true {
		os.Exit(0)
	}

	file, err := os.Open(*filename)
	if err != nil {
		log.Fatalf("Error opening %v: %v", *filename, err)
	}
	defer file.Close()

	videoResponse, err := service.Videos.Insert([]string{"snippet", "status"}, upload).Media(file).Do()
	if err != nil {
		panic(fmt.Sprintf("Error calling YouTube client: %v", err))
	}

	fmt.Printf("Upload successful! Video ID: %v\n", videoResponse.Id)
}
