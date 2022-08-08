package youtube

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/pkg/errors"
	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
)

func UploadToYoutube(apikey string, filename string, title string, description string, category string, privacy string, keywords string) (string, error) {
	ctx := context.Background()
	service, err := youtube.NewService(ctx, option.WithAPIKey(apikey))
	if err != nil {
		return "", errors.Wrap(err, "Could not create Youtube client")
	}

	upload := &youtube.Video{
		Snippet: &youtube.VideoSnippet{
			Title:       title,
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

	file, err := os.Open(filename)
	defer file.Close()
	if err != nil {
		log.Fatalf("Error opening %v: %v", filename, err)
	}

	response, err := call.Media(file).Do()
	if err != nil {
		return "", errors.Wrap(err, "Could not upload media file to Youtube")
	}
	fmt.Printf("Upload successful! Video ID: %v\n", response.Id)
	return response.Id, nil
}
