package api

import (
	"fmt"
	"net/http"
	"time"

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
		var youtubeItems []*youtube.SearchResult

		toDateTime := time.Unix(time.Now().Unix(), 0).Format(time.RFC3339)
		fromDateTime := time.Unix(0, 0).Format(time.RFC3339) //carbon.Parse("1970-01-01").ToRfc3339String()
		for true {
			channelResponse, err := service.Search.List([]string{"snippet"}).Type("video").MaxResults(50).PublishedAfter(fromDateTime).PublishedBefore(toDateTime).ForMine(true).Do()
			if err != nil {
				con.JSON(http.StatusInternalServerError, err)
			}
			youtubeItems = append(youtubeItems, channelResponse.Items...)

			fromDateTime = *&channelResponse.Items[0].Snippet.PublishedAt
			if fromDateTime == "" {
				break
			}
			break
		}

		con.JSON(http.StatusOK, youtubeItems)
	}

	return gin.HandlerFunc(fn)
}
