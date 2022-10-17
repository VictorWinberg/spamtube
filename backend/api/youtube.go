package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"spamtube/backend/domain"

	"github.com/gin-gonic/gin"
	"google.golang.org/api/youtube/v3"
)

func GetUploadedYoutubeVideos() gin.HandlerFunc {
	fn := func(con *gin.Context) {
		client := GetClient(youtube.YoutubeReadonlyScope, youtube.YoutubeUploadScope)

		service, err := youtube.New(client)
		if err != nil {
			panic(fmt.Sprintf("Error creating YouTube client: %v", err))
		}

		url := fmt.Sprintf("https://oauth.reddit.com/api/search_reddit_names?query=%s&include_over_18=true", searchQuery)
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

		// add authorization header to the req
		// req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))

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
