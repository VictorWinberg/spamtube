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
	filename    = flag.String("filename", getEnv("FILENAME", ""), "Name of video file to upload")
	title       = flag.String("title", getEnv("TITLE", "SpamTube Default Title"), "Video title")
	description = flag.String("description", getEnv("DESCRIPTION", "SpamTube Default Description"), "Video description")
	category    = flag.String("category", getEnv("CATEGORY", "22"), "Video category")
	keywords    = flag.String("keywords", getEnv("KEYWORDS", "spamtube,news,ai"), "Comma separated list of video keywords")
	privacy     = flag.String("privacy", getEnv("PRIVACY", "public"), "Video privacy status")
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

	file, err := os.Open(*filename)
	if err != nil {
		log.Fatalf("Error opening %v: %v", *filename, err)
	}
	defer file.Close()

	videoResponse, err := service.Videos.Insert([]string{"snippet", "status"}, upload).Media(file).Do()
	if err != nil {
		panic(fmt.Sprintf("Error calling YouTube client: %v", err))
	}

	fmt.Printf("Upload successful! Video URL: https://www.youtube.com/watch?v=%v\n", videoResponse.Id)
}

func getEnv(key, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		value = fallback
	}
	return value
}
