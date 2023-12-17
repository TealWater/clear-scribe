package controller

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	model "github.com/TealWater/clear-scribe/src/Model"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var collection *mongo.Collection
var database *mongo.Database
var mongoClient *mongo.Client

const dbName = "notes"
const colName = "messages"

func init() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("unable to laod environment variables")
	}

	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	mongoURI := "mongodb+srv://" + os.Getenv("MONGO_DB_USERNAME") + ":" + os.Getenv("MONGO_DB_PASSWORD") + "@cluster0.lx82yxi.mongodb.net/?retryWrites=true&w=majority"
	opts := options.Client().ApplyURI(mongoURI).SetServerAPIOptions(serverAPI)

	mongoClient, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		log.Println("unable to connect to DB")
		log.Panic(err)
	}

	//running this will close the database, thus throwing an error
	// defer func() {
	// 	if err = mongoClient.Disconnect(context.TODO()); err != nil {
	// 		log.Println("error trying to disconnect from the DB")
	// 		log.Panic(err)
	// 	}
	// }()

	// Send a ping to confirm a successful connection
	if err := mongoClient.Database("admin").RunCommand(context.TODO(), bson.D{{"ping", 1}}).Err(); err != nil {
		panic(err)
	}
	log.Println("Pinged your deployment. You successfully connected to MongoDB!")

	databases, err := mongoClient.ListDatabaseNames(context.TODO(), bson.M{})
	if err != nil {
		log.Fatal("line 57!: ", err)
	}
	log.Println(databases)

	database = mongoClient.Database(dbName)
	collection = database.Collection(colName)
	fmt.Println("collection name: ", collection.Name())

	/* Testing */
	// insertResult, err := collection.InsertOne(context.Background(), bson.D{
	// 	{Key: "createdAt", Value: time.Now()},
	// 	{Key: "messageOld", Value: "Taking a stroll in the park"},
	// 	{Key: "messageNew", Value: "Taking a walk in the park"},
	// })
	// if err != nil {
	// 	log.Fatal("failed at line 75!", err)
	// }
	// fmt.Println(insertResult.InsertedID)

	// tempEssay := model.EditedEssay{
	// 	CreatedAt:  primitive.NewDateTimeFromTime(time.Now()),
	// 	MessageOld: "This book is a very cathardic read",
	// 	MessageNew: "This book is a very relaxing read",
	// }
	// insertResult, err := collection.InsertOne(context.TODO(), tempEssay)
	// if err != nil {
	// 	log.Fatal("failed at line 75!", err)
	// }
	// fmt.Printf("Inserted %v!", insertResult.InsertedID)

}

func InsertMessages(messageOld, messageNew string) {
	entry := model.EditedEssay{
		CreatedAt:  primitive.NewDateTimeFromTime(time.Now()),
		MessageOld: messageOld,
		MessageNew: messageNew,
	}

	inserted, err := collection.InsertOne(context.TODO(), entry)
	if err != nil {
		log.Println("unable to insert data")
		log.Fatal(err)
	}

	log.Println("Insereted new row entry with the id of: ", inserted.InsertedID)
}

func deleteMessages(messageId string) {
	id, _ := primitive.ObjectIDFromHex(messageId)
	filter := bson.M{"_id": id}
	deleteCount, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Println("unable to delete record from the database")
		log.Fatal(err)
	}

	log.Println("Message with id ", messageId, " was deleted, with count of: ", deleteCount.DeletedCount)
}

func deleteAllMessages() {
	deleteCount, err := collection.DeleteMany(context.Background(), bson.D{{}}, nil)
	if err != nil {
		log.Println("Unable to delete all records")
		log.Fatal(err)
	}

	log.Println("we deleted ", deleteCount.DeletedCount, " records")
}

func getAllMessages() ([]primitive.M, []primitive.M) {
	cur, err := collection.Find(context.Background(), bson.D{{}})
	if err != nil {
		log.Println("Unable to retreive all of the records")
		log.Fatalln(err)
	}

	var messagesOld []primitive.M
	var messagesNew []primitive.M

	defer cur.Close(context.Background())
	for cur.Next(context.Background()) {
		var messageOld bson.M
		var messageNew bson.M

		if err := cur.Decode(&messageOld); err != nil {
			log.Println("unable to decode old message")
			log.Fatalln(err)
		}
		if err := cur.Decode(&messageNew); err != nil {
			log.Println("unable to decode new message")
			log.Fatalln(err)
		}

		messagesOld = append(messagesOld, messageOld)
		messagesNew = append(messagesNew, messageNew)
	}

	return messagesOld, messagesNew
}

func GetAllRecords(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	c.Request.Method = "POST"
	allRecordsOld, allRecordsNew := getAllMessages()
	c.JSON(http.StatusOK, allRecordsOld)
	c.JSON(http.StatusOK, allRecordsNew)
}

func AddARecord(c *gin.Context, msgOld, msgNew string) {
	var mmg model.EditedEssay

	mmg.MessageOld = msgOld
	mmg.MessageNew = msgNew

	InsertMessages(mmg.MessageOld, mmg.MessageNew)

}

func CloseDB() {
	if err := mongoClient.Disconnect(context.TODO()); err != nil {
		log.Println("error trying to disconnect from the DB")
		log.Panic(err)
	}
}
