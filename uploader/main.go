package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
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
	ctx := context.Background()
	service, err := youtube.NewService(ctx, option.WithAPIKey("API_KEY"))
	if err != nil {
		log.Fatalf("Error creating YouTube client: %v", err)
	}

	upload := &youtube.Video{
		Snippet: &youtube.VideoSnippet{
			Title:       *title,
			Description: description,
			CategoryId:  category,
		},
		Status: &youtube.VideoStatus{PrivacyStatus: privacy},
	}

	// The API returns a 400 Bad Request response if tags is an empty string.
	if strings.Trim(keywords, "") != "" {
		upload.Snippet.Tags = strings.Split(keywords, ",")
	}

	//Fix this part, feels odd
	var part []string
	part[0] = "snippet"
	part[1] = "status"
	call := service.Videos.Insert(part, upload)

	file, err := os.Open(*filename)
	defer file.Close()
	if err != nil {
		log.Fatalf("Error opening %v: %v", *filename, err)
	}

	response, err := call.Media(file).Do()
	if err != nil {
		log.Fatalf("Could not upload media to Youtube. Error: %s", err)
		return
	}
	fmt.Printf("Upload successful! Video ID: %v\n", response.Id)

}
