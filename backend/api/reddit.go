package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"spamtube/backend/domain"
	"spamtube/backend/keywords"

	"github.com/gin-gonic/gin"
	"github.com/patrickmn/go-cache"
)

func GetTopPosts(c *cache.Cache) gin.HandlerFunc {
	fn := func(con *gin.Context) {
		subredditName := con.Param("subreddit_name")
		url := fmt.Sprintf("https://reddit.com/r/%s/top.json?t=day", subredditName)
		req, err := http.NewRequest("GET", url, nil)

		if err != nil {
			con.JSON(http.StatusInternalServerError, gin.H{
				"message": fmt.Sprintf("Error: %s", err),
			})
			return
		}

		if err != nil {
			con.JSON(http.StatusInternalServerError, gin.H{
				"message": fmt.Sprintf("Error: %s", err),
			})
			return
		}

		// Send req using http Client
		client := &http.Client{}
		resp, err := client.Do(req)

		if err != nil {
			con.JSON(http.StatusInternalServerError, gin.H{
				"message": fmt.Sprintf("Error: %s", err),
			})
		}

		defer resp.Body.Close()

		res := &domain.RedditResponseTop{}
		json.NewDecoder(resp.Body).Decode(&res)

		if resp.StatusCode != 200 {
			con.JSON(http.StatusInternalServerError, gin.H{
				"message": "Could not retrieve top posts",
			})
			return
		}

		for i, post := range res.Data.Children {
			k, _ := keywords.Extract(post.Data.Selftext)
			res.Data.Children[i].Data.Keywords = k
		}

		con.JSON(http.StatusOK, res.Data.Children)
	}

	return gin.HandlerFunc(fn)
}

func SearchSubreddits(c *cache.Cache) gin.HandlerFunc {
	fn := func(con *gin.Context) {
		searchQuery := con.Param("search_query")
		url := fmt.Sprintf("https://reddit.com/search.json?q=%s&sort=top&t=all&type=sr", searchQuery)
		req, err := http.NewRequest("GET", url, nil)

		if err != nil {
			con.JSON(http.StatusInternalServerError, gin.H{
				"message": fmt.Sprintf("Error: %s", err),
			})
			return
		}

		if err != nil {
			con.JSON(http.StatusInternalServerError, gin.H{
				"message": fmt.Sprintf("Error: %s", err),
			})
			return
		}

		// Send req using http Client
		client := &http.Client{}
		resp, err := client.Do(req)

		if err != nil {
			con.JSON(http.StatusInternalServerError, gin.H{
				"message": fmt.Sprintf("Error: %s", err),
			})
		}

		defer resp.Body.Close()

		res := &domain.SearchResult{}
		json.NewDecoder(resp.Body).Decode(&res)

		if resp.StatusCode != 200 {
			con.JSON(http.StatusInternalServerError, gin.H{
				"message": "Could not search for subreddits",
			})
			return
		}

		subreddit_names := []string{}

		for _, n := range res.Data.Children {
			subreddit_names = append(subreddit_names, n.Data.Title)
		}

		con.JSON(http.StatusOK, subreddit_names)
	}

	return gin.HandlerFunc(fn)
}
