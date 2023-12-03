package controller

import (
	"context"
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

const dbName = "notes"

// const colNameOld = "messageOld"
// const colNameNew = "messageNew"
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

	collection = mongoClient.Database(dbName).Collection(colName)
}

func insertMessages(messageOld, messageNew string) {
	entry := model.EditedEssay{
		ID:         primitive.NewObjectID(),
		CreatedAt:  primitive.NewDateTimeFromTime(time.Time{}),
		MessageOld: messageOld,
		MessageNew: messageNew,
	}

	inserted, err := collection.InsertOne(context.Background(), entry)
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

	insertMessages(mmg.MessageOld, mmg.MessageNew)

}
