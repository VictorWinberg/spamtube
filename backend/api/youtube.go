package api

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"spamtube/backend/domain"

	"github.com/gin-gonic/gin"
)

const CHANNEL_ID = "UCTIp7LYLKOA6zq_PT21_NgA"

func getXML(url string) ([]byte, error) {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("Content-Type", "text/xml")
	resp, err := client.Do(req)

	if err != nil {
		return nil, fmt.Errorf("get error: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status error: %v", resp.StatusCode)
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read body: %v", err)
	}

	return data, nil
}

func GetUploadedYoutubeVideos() gin.HandlerFunc {
	fn := func(con *gin.Context) {
		url := fmt.Sprintf("https://www.youtube.com/feeds/videos.xml?channel_id=%s", CHANNEL_ID)
		resp, err := getXML(url)

		if err != nil {
			con.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			return
		}

		res := &domain.Feed{}
		xml.Unmarshal(resp, &res)

		con.JSON(http.StatusOK, res.Videos)
	}

	return gin.HandlerFunc(fn)
}
