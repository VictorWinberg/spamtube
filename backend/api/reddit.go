package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"spamtube/backend/domain"
	"spamtube/backend/keywords"
)

func GetTopPosts(subredditName string) (domain.RedditItems, error) {
	url := fmt.Sprintf("https://reddit.com/r/%s/top.json?t=day", subredditName)
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return nil, err
	}

	// Send req using http Client
	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	res := &domain.RedditResponseTop{}
	json.NewDecoder(resp.Body).Decode(&res)

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("COULD NOT RETRIEVE TOP POSTS")
	}

	for i, post := range res.Data.Children {
		k, _ := keywords.Extract(post.Data.Selftext)
		res.Data.Children[i].Data.Keywords = k
	}

	return res.Data.Children, nil
}

func SearchSubreddit(searchQuery string) ([]string, error) {
	url := fmt.Sprintf("https://reddit.com/search.json?q=%s&sort=top&t=all&type=sr", searchQuery)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	// Send req using http Client
	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	res := &domain.SearchResult{}
	json.NewDecoder(resp.Body).Decode(&res)

	if resp.StatusCode != 200 {
		return nil, err
	}

	subreddits := []string{}

	for _, n := range res.Data.Children {
		subreddits = append(subreddits, n.Data.Title)
	}

	return subreddits, nil
}
