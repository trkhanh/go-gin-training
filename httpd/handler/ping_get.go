package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)


func PingGet(depenencies string) gin.HandlerFunc {
	return func(c *gin.Context){
		c.JSON(http.StatusOK, map[string]string{"hello": `found  me` + depenencies}) 
	}
}