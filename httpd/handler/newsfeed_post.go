package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"go-gin-training/platform/newsfeed"
)

type newsFeedPostRequest struct{
	Title string `json:"title"`
	Post string `json:"post"`
}

func NewsFeedPost(feed *newsfeed.Repo) gin.HandlerFunc {
	return func(c *gin.Context){
		requestBody := newsFeedPostRequest{}
		c.Bind(&requestBody)

		item := newsfeed.Item{
			Title: requestBody.Title,
			Post: requestBody.Post,
		}

		feed.Add(item)

		c.Status(http.StatusNoContent,) 
	}
}

// curl -X POST -H "Content-Type: application/json" -d '{"title":"johndoe", "post":"abcd1234"}' localhost:8080/newsfeed