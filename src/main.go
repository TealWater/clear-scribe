package main

import (
	"log"

	controller "github.com/TealWater/clear-scribe/src/Controller"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	cors "github.com/rs/cors/wrapper/gin"
)

func init() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("unable to laod environment variables")
	}

}

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

	router.POST("/send", controller.UploadText)
	router.POST("/upload", controller.UploadFile)
	router.GET("/allRecords", controller.GetAllRecords)
	router.Run(":8080")
}
