package api

import (
	"fmt"
	"net/http"

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
		channelResponse, err := service.Search.List([]string{"snippet"}).Type("video").MaxResults(100).ForMine(true).Do()
		if err != nil {
			con.JSON(http.StatusInternalServerError, err)
		}
		con.JSON(http.StatusOK, channelResponse.Items)
	}

	return gin.HandlerFunc(fn)
}
