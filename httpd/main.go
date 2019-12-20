package main

import (
	"go-gin-training/platform/newsfeed"
	"github.com/gin-gonic/gin"
	"go-gin-training/httpd/handler"

)

func main() {
	feed := newsfeed.New()
	r := gin.Default()
	
	r.GET("/ping", handler.PingGet("ASUMM_DEPDENCE"))
	r.GET("/newsfeed", handler.NewsFeeds(feed))
	r.POST("/newsfeed", handler.NewsFeedPost(feed))

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

	// feed := newsfeed.New()
	// fmt.Println(feed)

	// feed.Add(newsfeed.Item{"Hello", "I'm new Feed"})

	// fmt.Println(feed)
}