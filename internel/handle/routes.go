package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const Version = "1.0.9"

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello World from CelerBuild!",
		})
	})

	r.GET("/version", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"version": Version,
		})
	})

	return r
}
