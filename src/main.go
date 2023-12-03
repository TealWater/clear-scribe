package main

import (
	controller "github.com/TealWater/clear-scribe/src/Controller"
	"github.com/gin-gonic/gin"
	cors "github.com/rs/cors/wrapper/gin"
)

func main() {
	router := gin.Default()

	router.GET("/time", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello world",
		})
	})

	router.Use(cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowCredentials: true,
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Access-Control-Allow-Origin", "Content-Type"},
	}))

	router.POST("/send", controller.Parse)
	router.POST("/upload", controller.Upload)
	router.Run(":8080")
}
