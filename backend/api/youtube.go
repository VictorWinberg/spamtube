package api

import (
	"encoding/json"
	"net/http"
	"os"
	"unicode/utf8"

	"strings"

	"github.com/anaskhan96/soup"
	"github.com/gin-gonic/gin"
)

func GetUploadedYoutubeVideos() gin.HandlerFunc {
	fn := func(con *gin.Context) {
		resp, err := soup.Get("https://www.youtube.com/channel/UCTIp7LYLKOA6zq_PT21_NgA/shorts")
		if err != nil {
			os.Exit(1)
		}
		doc := soup.HTMLParse(resp)
		links := doc.FindAll("script")
		for _, link := range links {

			var isFound = strings.Contains(link.HTML(), "GFEEDBACK")
			if isFound {
				var s = strings.ReplaceAll(link.HTML(), "<script nonce=", "") // >var ytInitialData =
				var s1 = strings.ReplaceAll(s, ">var ytInitialData = ", "")   // ;</script>"""
				var s2 = strings.ReplaceAll(s1, ";</script>", "")
				var s3 = s2[24:len(s2)]
				var s4 = trimLastChar(s3)
				var s5 = trimLastChar(s4)
				var s6 = trimLastChar(s5)
				var s7 = strings.Trim(s6, "\"")
				var jsonMap interface{}
				err := json.Unmarshal([]byte(s7), &jsonMap)
				if err != nil {
					con.JSON(http.StatusOK, err)
				}
				//var s4 = s3[1 : len(s3)-2]
				con.JSON(http.StatusOK, jsonMap)
			}

		}
		con.JSON(http.StatusOK, "")
	}

	return gin.HandlerFunc(fn)
}
func trimLastChar(s string) string {
	r, size := utf8.DecodeLastRuneInString(s)
	if r == utf8.RuneError && (size == 0 || size == 1) {
		size = 0
	}
	return s[:len(s)-size]
}
