package autoupload

import (
	"fmt"
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
	if len(posts) == 0 {
		fmt.Printf("Had no posts for subreddit %s\n", subreddit_name)
		return nil
	}
	rand.Seed(time.Now().UnixNano())
	post := posts[rand.Intn(len(posts))]
	hashTaggedKeywords := ""

	for _, keyword := range post.Data.Keywords {
		// ignore keywords containing '
		if !strings.Contains(keyword, "'") {
			hashTaggedKeywords = hashTaggedKeywords + fmt.Sprintf("#%s ", keyword)
		}
	}

	data := &api.WorkflowInputBody{
		Ref: "master",
		Inputs: &api.WorkflowInputs{
			Title:       post.Data.Title,
			Description: fmt.Sprintf("See the reddit post that generated this video here: %s \n %s", post.Data.URL, hashTaggedKeywords),
			Image:       strings.Join(post.Data.Keywords, " "),
			TextContent: post.Data.Selftext,
		},
	}

	resp, err := api.TriggerGithubAction(data)
	if err != nil {
		return err
	}
	fmt.Println(resp)
	return nil
}
