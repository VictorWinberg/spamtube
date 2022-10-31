package api

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gocolly/colly"
)

func GetUploadedYoutubeVideos() gin.HandlerFunc {
	fn := func(con *gin.Context) {
		var array []*string
		c := colly.NewCollector(
			colly.AllowedDomains("books.toscrape.com"),
		)
		// https://oxylabs.io/blog/golang-web-scraper
		c.OnHTML("title", func(e *colly.HTMLElement) {
			fmt.Println(e.Text)
			array = append(array, &e.Text)
		})

		c.OnResponse(func(r *colly.Response) {
			fmt.Println(r.StatusCode)
		})

		c.OnRequest(func(r *colly.Request) {
			fmt.Println("Visiting", r.URL)
		})

		c.Visit("https://books.toscrape.com/")

		con.JSON(http.StatusOK, array)
	}

	return gin.HandlerFunc(fn)
}
