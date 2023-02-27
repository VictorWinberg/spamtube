package autoupload

import (
	"fmt"
	"log"
	"math/rand"
	api "spamtube/backend/api"
	"strings"
	"time"
)

func AutoUploadVideo(subreddit_name string) error {
	fmt.Printf("Autouploading video from %s\n", subreddit_name)
	posts, err := api.GetTopPosts(subreddit_name)
	if err != nil {
		return err
	}
	rand.Seed(time.Now().UnixNano())
	post := posts[rand.Intn(len(posts))]

	data := &api.WorkflowInputBody{
		Ref: "master",
		Inputs: &api.WorkflowInputs{
			Title:       post.Data.Title,
			Description: post.Data.URL,
			Image:       strings.Join(post.Data.Keywords, " "),
			Voice:       post.Data.Selftext,
			Service:     "image-finder",
		},
	}

	resp, err := api.TriggerGithubAction(data)
	if err != nil {
		return err
	}
	log.Println(resp)
	return nil
}
