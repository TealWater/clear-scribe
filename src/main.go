package main

import (
	controller "github.com/TealWater/clear-scribe/src/Controller"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/time", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello world",
		})
	})

	router.POST("/send", controller.Parse)
	// router.POST("/upload", controller.Upload)
	router.Run(":8080")
}
