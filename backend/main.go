package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	internalApi "spamtube/backend/api"
	"spamtube/backend/autoupload"
	"spamtube/backend/helpers"
	"time"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/robfig/cron/v3"
)

func main() {
	err := godotenv.Load(".env.local")
	if err != nil {
		fmt.Printf("[WARN]: %s", err)
	}
	godotenv.Load()

	port := flag.String("port", helpers.GetEnv("PORT", "3000"), "Server port")
	flag.Parse()

	loc, err := time.LoadLocation("Europe/Stockholm")
	if err != nil {
		log.Fatal(err)
	}

	kron := cron.New(cron.WithLocation(loc))
	kron.AddFunc("0 9 * * *", func() {
		err := autoupload.AutoUploadVideo("AmItheAsshole")
		if err != nil {
			log.Println(err)
			return
		}
		log.Println("Autouploading video successful!")
	})
	kron.Start()

	// Set the router as the default one shipped with Gin
	router := gin.Default()

	// Serve frontend static files
	router.Use(static.Serve("/", static.LocalFile("./dist", true)))

	// Setup route group for the API
	api := router.Group("/api")
	{
		api.GET("/top/:subreddit_name", func(con *gin.Context) {
			subredditName := con.Param("subreddit_name")
			posts, err := internalApi.GetTopPosts(subredditName)
			if err != nil {
				con.JSON(http.StatusInternalServerError, gin.H{
					"message": fmt.Sprintf("Error: %s", err),
				})
				return
			}
			con.JSON(http.StatusOK, posts)
		})

		api.GET("/search/:search_query", func(con *gin.Context) {
			searchQuery := con.Param("search_query")
			subreddits, err := internalApi.SearchSubreddit(searchQuery)
			if err != nil {
				con.JSON(http.StatusInternalServerError, gin.H{
					"message": fmt.Sprintf("Error: %s", err),
				})
				return
			}
			con.JSON(http.StatusOK, subreddits)
		})

		api.POST("/generate", func(con *gin.Context) {
			body := &internalApi.WorkflowInputBody{}
			json.NewDecoder(con.Request.Body).Decode(&body)
			res, err := internalApi.TriggerGithubAction(body)
			if err != nil {
				con.JSON(http.StatusInternalServerError, gin.H{
					"message": fmt.Sprintf("Error: %s", err),
				})
				return
			}
			con.JSON(http.StatusOK, res)
		})

		api.GET("/videos", func(con *gin.Context) {
			videos, err := internalApi.GetUploadedYoutubeVideos()
			if err != nil {
				con.JSON(http.StatusInternalServerError, gin.H{
					"message": fmt.Sprintf("Error: %s", err),
				})
				return
			}
			con.JSON(http.StatusOK, videos)
		})
	}

	router.NoRoute(func(c *gin.Context) {
		c.File("./dist/index.html")
	})

	// Start and run the server
	router.Run(fmt.Sprintf(":%s", *port))
}
