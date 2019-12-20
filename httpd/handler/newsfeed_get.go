package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"go-gin-training/platform/newsfeed"
	"fmt"
)


func NewsFeeds(feed newsfeed.Getter) gin.HandlerFunc {
	fmt.Println(feed)
	return func(c *gin.Context){
		results := feed.GetAll()
		c.JSON(http.StatusOK, results) 
	}
}