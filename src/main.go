package main

import (
	"context"
	"log"
	"os"

	controller "github.com/TealWater/clear-scribe/src/Controller"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	cors "github.com/rs/cors/wrapper/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func init() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("unable to laod environment variables")
	}

}

func main() {
	router := gin.Default()
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	mongoURI := "mongodb+srv://" + os.Getenv("MONGO_DB_USERNAME") + ":" + os.Getenv("MONGO_DB_PASSWORD") + "@cluster0.lx82yxi.mongodb.net/?retryWrites=true&w=majority"
	opts := options.Client().ApplyURI(mongoURI).SetServerAPIOptions(serverAPI)

	mongoClient, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		log.Println("unable to connect to DB")
		log.Panic(err)
	}

	defer func() {
		if err = mongoClient.Disconnect(context.TODO()); err != nil {
			log.Println("error trying to disconnect from the DB")
			log.Panic(err)
		}
	}()

	// Send a ping to confirm a successful connection
	if err := mongoClient.Database("admin").RunCommand(context.TODO(), bson.D{{"ping", 1}}).Err(); err != nil {
		panic(err)
	}
	log.Println("Pinged your deployment. You successfully connected to MongoDB!")

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
