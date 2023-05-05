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

	req.Header.Set("User-agent", "spamtube 0.1")

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
		return nil, fmt.Errorf("STATUS CODE: %d", resp.StatusCode)
	}

	for i, post := range res.Data.Children {
		postKeywords, _ := keywords.Extract(post.Data.Selftext)
		res.Data.Children[i].Data.Keywords = postKeywords
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
