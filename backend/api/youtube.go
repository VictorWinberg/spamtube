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
		channelResponse, err := service.Videos.List([]string{"contentDetails",""}).Do()
		if err != nil {
			con.JSON(http.StatusInternalServerError, err)
		}
		con.JSON(http.StatusOK, channelResponse.Items)
	}

	return gin.HandlerFunc(fn)
}
