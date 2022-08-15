/*docs:
 https://developers.google.com/oauthplayground/?code=4/0AdQt8qh9L5I7lRECK03RdtetenpQ2xKfpB_lQ4U0KJbTtvFeYe5YUxIvWvd841eAAVp-kg&scope=https://www.googleapis.com/auth/youtube%20https://www.googleapis.com/auth/youtube.channel-memberships.creator%20https://www.googleapis.com/auth/youtube.force-ssl%20https://www.googleapis.com/auth/youtube.readonly%20https://www.googleapis.com/auth/youtube.upload%20https://www.googleapis.com/auth/youtubepartner%20https://www.googleapis.com/auth/youtubepartner-channel-audit
https://github.com/googleapis/google-api-go-client/blob/main/GettingStarted.md

*/
package youtube

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/pkg/errors"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/youtube/v3"
)

func UploadToYoutube(apikey string, filename string, title string, description string, category string, privacy string, keywords string) (string, error) {

	ctx := context.WithValue(context.Background(), oauth2.HTTPClient, nil)
	var oauthConfig = &oauth2.Config{
		ClientID:     "<client-id>", // from https://console.developers.google.com/project/<your-project-id>/apiui/credential
		ClientSecret: "<secret>",    // from https://console.developers.google.com/project/<your-project-id>/apiui/credential
		Endpoint:     google.Endpoint,
		Scopes:       []string{youtube.YoutubeUploadScope},
	}
	httpClient := oauthConfig.Client(ctx, &oauth2.Token{AccessToken: "<access token>",
		TokenType:    "Bearer",
		RefreshToken: "<refresh token>",
	})
	service, err := youtube.New(httpClient)
	if err != nil {
		return "", errors.Wrap(err, "Could not create Youtube client")
	}

	upload := &youtube.Video{
		Snippet: &youtube.VideoSnippet{
			Title:       title,
			Description: description,
			CategoryId:  category,
			ChannelId:   "UCfiK-MFgww_xENZqgnaH9_A",
		},
		Status: &youtube.VideoStatus{PrivacyStatus: privacy},
	}

	// The API returns a 400 Bad Request response if tags is an empty string.
	if strings.Trim(keywords, "") != "" {
		upload.Snippet.Tags = strings.Split(keywords, ",")
	}

	//Fix this part, feels odd
	var part = make([]string, 2)

	part[0] = "snippet"
	part[1] = "status"

	fmt.Println(part)
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
