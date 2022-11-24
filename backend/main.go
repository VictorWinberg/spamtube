package main

import (
	"flag"
	"fmt"
	internalApi "spamtube/backend/api"
	"spamtube/backend/helpers"
	"time"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	cache "github.com/patrickmn/go-cache"
)

func main() {
	err := godotenv.Load(".env.local")
	if err != nil {
		fmt.Printf("[WARN]: %s", err)
	}
	godotenv.Load()

	port := flag.String("port", helpers.GetEnv("PORT", "3000"), "Server port")
	flag.Parse()
	c := cache.New(23*time.Hour, 10*time.Minute)

	// Set the router as the default one shipped with Gin
	router := gin.Default()

	// Serve frontend static files
	router.Use(static.Serve("/", static.LocalFile("./dist", true)))

	// Setup route group for the API
	api := router.Group("/api")
	{
		api.GET("/top/:subreddit_name", internalApi.GetTopPosts(c))
		api.GET("/search/:search_query", internalApi.SearchSubreddits(c))
		api.POST("/generate", internalApi.TriggerGithubAction())
		api.GET("/videos", internalApi.GetUploadedYoutubeVideos())
	}

	router.NoRoute(func(c *gin.Context) {
		c.File("./dist/index.html")
	})

	// Start and run the server
	router.Run(fmt.Sprintf(":%s", *port))
}
