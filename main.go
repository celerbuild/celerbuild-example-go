package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Version of the application
const Version = "1.0.0"

func main() {
	// Create a default Gin router
	r := gin.Default()

	// Define a GET route for the root path "/"
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello World from CelerBuild!",
		})
	})

	// Define a GET route for "/version"
	r.GET("/version", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"version": Version,
		})
	})

	// Run the server on port 8080
	r.Run(":8080")
}
