package domain

type AccessToken struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
	Scope       string `json:"scope"`
}

type RedditResponseTop struct {
	Data struct {
		Children RedditItems `json:"children"`
	} `json:"data"`
}

type RedditItems []struct {
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
}

type SearchResult struct {
	Data struct {
		Children []struct {
			Data struct {
				Title string `json:"title"`
			} `json:"data,omitempty"`
		} `json:"children"`
	} `json:"data"`
}

type Feed struct {
	Videos []Video `xml:"entry"`
}

type Video struct {
	Title string `xml:"title"`
	Link  struct {
		Href string `xml:"href,attr"`
	} `xml:"link"`
	Published string `xml:"published"`
	Updated   string `xml:"updated"`
	Group     struct {
		Thumbnail struct {
			URL string `xml:"url,attr"`
		} `xml:"thumbnail"`
		Description string `xml:"description"`
		Community   struct {
			StarRating struct {
				Count   string `xml:"count,attr"`
				Average string `xml:"average,attr"`
			} `xml:"starRating"`
			Statistics struct {
				Views string `xml:"views,attr"`
			} `xml:"statistics"`
		} `xml:"community"`
	} `xml:"group"`
	VideoId string `xml:"videoId"`
}
