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

	var hashTags strings.Builder
	for _, keyword := range post.Data.Keywords {
		hashTags.WriteString(strings.Replace(keyword, "'", "", -1))
	}

	var keywords []string
	for _, words := range chunkBy(post.Data.Keywords, 3) {
		keywords = append(keywords, strings.Join(words, " "))
	}

	data := &api.WorkflowInputBody{
		Ref: "master",
		Inputs: &api.WorkflowInputs{
			Title:         post.Data.Title,
			Description:   fmt.Sprintf("Reddit url: %s \n %s", post.Data.URL, hashTags.String()),
			ImageKeywords: strings.Join(keywords, ","),
			TextContent:   post.Data.Selftext,
		},
	}

	resp, err := api.TriggerGithubAction(data)
	if err != nil {
		return err
	}
	fmt.Println(resp)
	return nil
}

func chunkBy[T any](items []T, chunkSize int) (chunks [][]T) {
	for chunkSize < len(items) {
		items, chunks = items[chunkSize:], append(chunks, items[0:chunkSize:chunkSize])
	}
	return append(chunks, items)
}
