package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	cache "github.com/patrickmn/go-cache"
	"github.com/pkg/errors"
)

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
	Kind string `json:"kind"`
	Data struct {
		Children []struct {
			Data struct {
				Subreddit         string  `json:"subreddit"`
				Selftext          string  `json:"selftext"`
				AuthorFullname    string  `json:"author_fullname"`
				Title             string  `json:"title"`
				UpvoteRatio       float64 `json:"upvote_ratio"`
				Ups               int     `json:"ups"`
				IsOriginalContent bool    `json:"is_original_content"`
				Score             int     `json:"score"`
				Thumbnail         string  `json:"thumbnail"`
				Created           float64 `json:"created"`
				Over18            bool    `json:"over_18"`
				SubredditID       string  `json:"subreddit_id"`
				ID                string  `json:"id"`
				Author            string  `json:"author"`
				Permalink         string  `json:"permalink"`
				URL               string  `json:"url"`
				CreatedUtc        float64 `json:"created_utc"`
			} `json:"data"`
		} `json:"children"`
	} `json:"data"`
}

func main() {
	err := godotenv.Load(".env.local")
	if err != nil {
		fmt.Printf("[WARN]: %s", err)
	}
	godotenv.Load()

	port := flag.String("port", getEnv("PORT", "3000"), "Server port")
	flag.Parse()
	c := cache.New(23*time.Hour, 10*time.Minute)

	// Set the router as the default one shipped with Gin
	router := gin.Default()

	// Serve frontend static files
	router.Use(static.Serve("/", static.LocalFile("./dist", true)))

	// Setup route group for the API
	api := router.Group("/api")
	{
		api.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "pong",
			})
		})

		api.GET("/reddit", func(con *gin.Context) {
			token, err := getAccessToken(c)

			if err != nil {
				con.JSON(http.StatusInternalServerError, gin.H{
					"message": fmt.Sprintf("Error: %s", err),
				})
				return
			}

			con.JSON(http.StatusOK, token)
		})

		api.GET("/top/:subreddit_name", func(con *gin.Context) {
			subredditName := con.Param("subreddit_name")
			url := fmt.Sprintf("https://oauth.reddit.com/r/%s/top/?t=day.json", subredditName)
			req, err := http.NewRequest("GET", url, nil)

			if err != nil {
				con.JSON(http.StatusInternalServerError, gin.H{
					"message": fmt.Sprintf("Error: %s", err),
				})
				return
			}

			token, err := handleTokenLogic(c, con)

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

			res := &RedditResponseTop{}
			json.NewDecoder(resp.Body).Decode(&res)

			if resp.StatusCode != 200 {
				con.JSON(http.StatusInternalServerError, gin.H{
					"message": "Could not retrieve top posts",
				})
				return
			}

			con.JSON(http.StatusOK, res.Data.Children)
		})

		api.GET("/search/:search_query", func(con *gin.Context) {
			searchQuery := con.Param("search_query")
			url := fmt.Sprintf("https://oauth.reddit.com/api/search_reddit_names?query=%s&include_over_18=true", searchQuery)
			req, err := http.NewRequest("GET", url, nil)

			if err != nil {
				con.JSON(http.StatusInternalServerError, gin.H{
					"message": fmt.Sprintf("Error: %s", err),
				})
				return
			}

			token, err := handleTokenLogic(c, con)

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

			res := &SearchResult{}
			json.NewDecoder(resp.Body).Decode(&res)

			if resp.StatusCode != 200 {
				con.JSON(http.StatusInternalServerError, gin.H{
					"message": "Could not search for subreddits",
				})
				return
			}

			con.JSON(http.StatusOK, res.Names)
		})
	}

	router.NoRoute(func(c *gin.Context) {
		c.File("./dist/index.html")
	})

	// Start and run the server
	router.Run(fmt.Sprintf(":%s", *port))
}

func getEnv(key, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		value = fallback
	}
	return value
}

func getAccessToken(c *cache.Cache) (*AccessToken, error) {
	client := &http.Client{}
	URL := "https://www.reddit.com/api/v1/access_token"
	v := url.Values{
		"grant_type": {"client_credentials"},
		"username":   {os.Getenv("REDDIT_USERNAME")},
		"password":   {os.Getenv("REDDIT_PASSWORD")},
	}
	req, err := http.NewRequest("POST", URL, strings.NewReader(v.Encode()))

	if err != nil {
		return nil, err
	}

	req.SetBasicAuth(os.Getenv("REDDIT_APP_USERNAME"), os.Getenv("REDDIT_APP_PRIVATE_KEY"))
	res, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	if res.StatusCode != 200 {
		return nil, errors.Wrap(err, "Could not retrieve access token from reddit")
	}

	token := &AccessToken{}
	defer res.Body.Close()
	json.NewDecoder(res.Body).Decode(&token)

	c.Set("token", token.AccessToken, cache.DefaultExpiration)

	return token, nil
}

func handleTokenLogic(c *cache.Cache, con *gin.Context) (string, error) {
	var token string
	if val, found := c.Get("token"); found {
		token = val.(string)
	} else {
		t, err := getAccessToken(c)
		if err != nil {
			return "", err
		}
		token = t.AccessToken
	}
	return token, nil
}
