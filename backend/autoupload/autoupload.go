package autoupload

import (
	"log"
	"math/rand"
	api "spamtube/backend/api"
	"strings"
)

func AutoUploadVideo(subreddit_name string) {
	posts, err := api.GetTopPosts(subreddit_name)
	if err != nil {
		log.Fatal(err)
	}
	post := posts[rand.Intn(len(posts))]

	data := &api.WorkflowInputBody{
		Ref: "master",
		Inputs: &api.WorkflowInputs{
			Title:       post.Data.Title,
			Description: "Hello spamtubers",
			Image:       strings.Join(post.Data.Keywords, " "),
			Voice:       post.Data.Selftext,
			Service:     "image-finder",
		},
	}

	resp, err := api.TriggerGithubAction(data)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(resp)
}
