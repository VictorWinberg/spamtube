package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"time"
	"encoding/json"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/patrickmn/go-cache"
)

type RedditResponseTop struct {
	Kind string `json:"kind"`
	Data struct {
		After     string `json:"after"`
		Children  []struct {
			Data struct {
				Subreddit                  string        `json:"subreddit"`
				Selftext                   string        `json:"selftext"`
				AuthorFullname             string        `json:"author_fullname"`
				Title                      string        `json:"title"`
				UpvoteRatio                float64       `json:"upvote_ratio"`
				Ups                        int           `json:"ups"`
				IsOriginalContent     bool          `json:"is_original_content"`
				Score               int           `json:"score"`
				Thumbnail           string        `json:"thumbnail"`
				Created             float64     `json:"created"`
				Over18              bool        `json:"over_18"`
				SubredditID              string        `json:"subreddit_id"`
				ID                       string        `json:"id"`
				Author                   string        `json:"author"`
				Permalink                string        `json:"permalink"`
				URL                      string        `json:"url"`
				CreatedUtc               float64       `json:"created_utc"`
			} `json:"data"`
		} `json:"children"`
		Before interface{} `json:"before"`
	} `json:"data"`
}

func main() {
	port := flag.String("port", getEnv("PORT", "3000"), "Server port")
	flag.Parse()
	c := cache.New(23*time.Hour, 10*time.Minute)

	// Set the router as the default one shipped with Gin
	router := gin.Default()

	// Serve frontend static files
	router.Use(static.Serve("/", static.LocalFile("../frontend/dist", true)))

	// Setup route group for the API
	api := router.Group("/api")
	{
		api.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "pong",
			})
		})

		api.GET("/top/:subreddit_name", func(con *gin.Context) {
			subredditName := con.Param("subreddit_name")
			url := fmt.Sprintf("https://oauth.reddit.com/r/%s/top/?t=day.json", subredditName)
			req, err := http.NewRequest("GET", url, nil)
			
			if err != nil {
				fmt.Print(err.Error())
			}

			token, found := c.Get("token")

			if !found {
				// do request to reddit and get token
				// Winbergs awesome metod
			}

			// add authorization header to the req
			req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token.(string)))

			// Send req using http Client
			client := &http.Client{}
			resp, err := client.Do(req)
			
			if err != nil {
				con.JSON(http.StatusInternalServerError, gin.H{
					"message" : fmt.Sprintf("Error: %s", err),
				})
			}

			defer resp.Body.Close()

			res := &RedditResponseTop{}
			json.NewDecoder(resp.Body).Decode(&res)

			if resp.StatusCode != 200 {
				con.JSON(http.StatusInternalServerError, gin.H{
					"message" : "Could not retrieve top posts",
				})
				return
			}

			con.JSON(http.StatusOK, res)
			return
			})
	}

	// Start and run the server
	router.Run(fmt.Sprintf(":%s", *port))
}

func getEnv(key, fallback string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		value = fallback
	}
	return value
}
