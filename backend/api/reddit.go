package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"spamtube/backend/domain"
	"spamtube/backend/helpers"

	"github.com/gin-gonic/gin"
	"github.com/patrickmn/go-cache"
)

func GetTopPosts(c *cache.Cache) gin.HandlerFunc {
	fn := func(con *gin.Context) {
		subredditName := con.Param("subreddit_name")
		url := fmt.Sprintf("https://oauth.reddit.com/r/%s/top/?t=day.json", subredditName)
		req, err := http.NewRequest("GET", url, nil)

		if err != nil {
			con.JSON(http.StatusInternalServerError, gin.H{
				"message": fmt.Sprintf("Error: %s", err),
			})
			return
		}

		token, err := helpers.HandleTokenLogic(c, con)

		if err != nil {
			con.JSON(http.StatusInternalServerError, gin.H{
				"message": fmt.Sprintf("Error: %s", err),
			})
			return
		}

		// add authorization header to the req
		req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))

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

		con.JSON(http.StatusOK, res.Data.Children)
	}

	return gin.HandlerFunc(fn)
}

func SearchSubreddits(c *cache.Cache) gin.HandlerFunc {
	fn := func(con *gin.Context) {
		searchQuery := con.Param("search_query")
		url := fmt.Sprintf("https://oauth.reddit.com/api/search_reddit_names?query=%s&include_over_18=true", searchQuery)
		req, err := http.NewRequest("GET", url, nil)

		if err != nil {
			con.JSON(http.StatusInternalServerError, gin.H{
				"message": fmt.Sprintf("Error: %s", err),
			})
			return
		}

		token, err := helpers.HandleTokenLogic(c, con)

		if err != nil {
			con.JSON(http.StatusInternalServerError, gin.H{
				"message": fmt.Sprintf("Error: %s", err),
			})
			return
		}

		// add authorization header to the req
		req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))

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

		con.JSON(http.StatusOK, res.Names)
	}

	return gin.HandlerFunc(fn)
}
