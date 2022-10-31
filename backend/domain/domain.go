package domain

type AccessToken struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
	Scope       string `json:"scope"`
}

type SearchResult struct {
	Names []string `json:"names"`
}

type RedditResponseTop struct {
	Data struct {
		Children []struct {
			Data struct {
				Subreddit         string   `json:"subreddit"`
				Selftext          string   `json:"selftext"`
				AuthorFullname    string   `json:"author_fullname"`
				Title             string   `json:"title"`
				UpvoteRatio       float64  `json:"upvote_ratio"`
				Ups               int      `json:"ups"`
				IsOriginalContent bool     `json:"is_original_content"`
				Score             int      `json:"score"`
				Thumbnail         string   `json:"thumbnail"`
				Created           float64  `json:"created"`
				Over18            bool     `json:"over_18"`
				SubredditID       string   `json:"subreddit_id"`
				ID                string   `json:"id"`
				Author            string   `json:"author"`
				Permalink         string   `json:"permalink"`
				URL               string   `json:"url"`
				CreatedUtc        float64  `json:"created_utc"`
				Keywords          []string `json:"keywords"`
			} `json:"data"`
		} `json:"children"`
	} `json:"data"`
}
