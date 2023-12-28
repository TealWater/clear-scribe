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
		AllowedOrigins:   []string{os.Getenv("TRUSTED_URL")},
		AllowCredentials: true,
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Access-Control-Allow-Origin", "Content-Type"},
	}))

	router.POST("/send", controller.UploadText)
	router.POST("/upload", controller.UploadFile)
	router.GET("/history", controller.GetAllRecords)
	router.GET("/gpt", controller.SendPrompt)
	router.DELETE("/history", controller.DeleteRecord)
	router.Run(":" + os.Getenv("PORT"))

	defer controller.CloseDB()
}
