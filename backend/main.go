package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"math/rand"
	"net/http"
	internalApi "spamtube/backend/api"
	"spamtube/backend/autoupload"
	"spamtube/backend/database"
	"spamtube/backend/domain"
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
	rand.Seed(time.Now().UnixNano())

	port := flag.String("port", helpers.GetEnv("PORT", "3000"), "Server port")
	flag.Parse()

	subredditCronjobs := make(map[string]cron.EntryID)
	err = database.Connect()
	if err != nil {
		panic(err)
	}

	kron := cron.New()
	kron.Start()

	startCronJobs(kron, subredditCronjobs)

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

		api.GET("/subreddits", func(con *gin.Context) {
			items, err := database.GetAllSubreddits()

			if err != nil {
				con.JSON(http.StatusInternalServerError, gin.H{
					"message": fmt.Sprintf("Error: %s", err),
				})
				return
			}
			con.JSON(http.StatusOK, items)
		})

		api.DELETE("/subreddits/:id", func(con *gin.Context) {
			id := con.Param("id")
			err := database.DeleteSubreddit(id)

			if err != nil {
				con.JSON(http.StatusInternalServerError, gin.H{
					"message": fmt.Sprintf("Error: %s", err),
				})
				return
			}

			cronJobId := subredditCronjobs[id]
			fmt.Printf("CRON: remove subreddit %s", id)
			kron.Remove(cronJobId)
			con.JSON(http.StatusOK, id)
		})

		api.PUT("/subreddits", func(con *gin.Context) {
			subreddit := domain.UpsertSubreddit{}
			json.NewDecoder(con.Request.Body).Decode(&subreddit)
			item, err := database.UpsertSubreddit(subreddit)

			if err != nil {
				con.JSON(http.StatusInternalServerError, gin.H{
					"message": fmt.Sprintf("Error: %s", err),
				})
				return
			}

			cronJobId := subredditCronjobs[*subreddit.Id]
			kron.Remove(cronJobId)

			if subreddit.Cron != nil {
				cronJobId, err = addCronJob(kron, *subreddit.Name, *subreddit.Cron)
				subredditCronjobs[*subreddit.Id] = cronJobId

				if err != nil {
					con.JSON(http.StatusInternalServerError, gin.H{
						"message": fmt.Sprintf("Error: %s", err),
					})
					return
				}
			}

			con.JSON(http.StatusOK, item)
		})

		api.POST("/img", func(con *gin.Context) {
			body := &domain.ImgBody{}
			json.NewDecoder(con.Request.Body).Decode(&body)
			if err != nil {
				con.JSON(http.StatusInternalServerError, gin.H{
					"message": fmt.Sprintf("Error: %s", err),
				})
				return
			}

			err = database.InsertImgUrl(*body.url)

			if err != nil {
				con.JSON(http.StatusInternalServerError, gin.H{
					"message": fmt.Sprintf("Error: %s", err),
				})
				return
			}
			con.JSON(http.StatusOK)
		})
		api.GET("/img", func(con *gin.Context) {
			items, err := database.SelectAllImages()

			if err != nil {
				con.JSON(http.StatusInternalServerError, gin.H{
					"message": fmt.Sprintf("Error: %s", err),
				})
				return
			}
			con.JSON(http.StatusOK,items)
		})
	}

	router.NoRoute(func(c *gin.Context) {
		c.File("./dist/index.html")
	})

	// Start and run the server
	router.Run(fmt.Sprintf(":%s", *port))
}

func startCronJobs(kron *cron.Cron, subredditCronjobs map[string]cron.EntryID) {
	fmt.Println("CRON: Starting cron jobs")
	subreddits, err := database.GetSubredditsWithCron()
	if err != nil {
		fmt.Printf("CRON: Error when fetching subreddits for cron: %v", err)
		return
	}
	fmt.Printf("CRON: Found %d subreddits to start cron job for", len(subreddits))
	for _, subreddit := range subreddits {
		cronJobId, err := addCronJob(kron, subreddit.Name, *subreddit.Cron)
		if err == nil {
			subredditCronjobs[subreddit.Id] = cronJobId
		}
	}

}

func addCronJob(kron *cron.Cron, subreddit string, cron_string string) (cron.EntryID, error) {
	fmt.Printf("CRON: add subreddit %s, with cron: %s", subreddit, cron_string)
	id, err := kron.AddFunc(cron_string, func() {
		err := autoupload.AutoUploadVideo(subreddit)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("Autouploading video successful!")
	})
	fmt.Printf("DEBUG: %d crons running", len(kron.Entries()))
	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	return id, nil
}
