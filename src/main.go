package main

import (
	"log"
	"os"

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
		AllowedOrigins:   []string{os.Getenv("TrustedURL")},
		AllowCredentials: true,
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Access-Control-Allow-Origin", "Content-Type"},
	}))

	router.POST("/send", controller.UploadText)
	router.POST("/upload", controller.UploadFile)
	// router.GET("/allRecords", controller.GetAllRecords)
	router.GET("/history", controller.UploadMockHistory)
	router.Run(":" + os.Getenv("PORT"))

}
