package controller

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	obj "github.com/TealWater/clear-scribe/src/Model"

	"github.com/gin-gonic/gin"
)

var oldMessage string
var mp map[string]string = make(map[string]string)

func init() {
	mp["i"] = "hi"
	mp["house"] = "dwelling"
	mp["pleased"] = "happy"
}

func UploadText(c *gin.Context) {
	msg := &obj.IncomingText{}
	var hi = "hi mom"
	fmt.Println(hi)
	if err := c.ShouldBindJSON(msg); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fmt.Printf("%+v\n", msg.Message)
	oldMessage = msg.Message
	words := strings.Split(msg.Message, " ")
	for _, v := range words {
		fmt.Println(v)
		//search map for synonym
		//words[k] = map[v]
	}

	newMessage := parse(oldMessage)
	insertMessages(oldMessage, newMessage)
	c.JSON(http.StatusOK, gin.H{"messageNew": newMessage})
}

func UploadFile(c *gin.Context) {
	fileUpload := &obj.FileUpload{}
	// Bind the file from the request to the struct
	if err := c.ShouldBind(fileUpload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Save the uploaded file
	err := c.SaveUploadedFile(fileUpload.File, fileUpload.File.Filename)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	fmt.Println("***: ", fileUpload.File.Filename)

	/*

		File parsing below
	*/

	src, err := fileUpload.File.Open()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	content, err := io.ReadAll(src)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	oldMessage = string(content)
	fmt.Println(oldMessage)
	newMessage := parse(oldMessage)
	insertMessages(oldMessage, newMessage)
	c.JSON(http.StatusOK, gin.H{"messageNew": newMessage})
}

func DeleteRecord(c *gin.Context) {
	id := c.Query("id")
	if err := deleteMessage(id); err != nil {
		log.Println(err)
		c.JSON(http.StatusNotFound, gin.H{"status": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "success"})
}

func GetAllRecords(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	c.Request.Method = "POST"
	allRecords := getAllMessages()
	c.JSON(http.StatusOK, allRecords)
}

/*

2 methods  --> one focus on open the file and getting the text data
			- focus on getting the data out of the response body

method #3 --> parse the text and add the changes


*/

// swap out complicated words for simple ones
func parse(msgOld string) string {
	words := strings.Split(msgOld, " ")
	for k, v := range words {
		v = strings.ToLower(v)
		if val, ok := mp[v]; ok {
			words[k] = val
		}
	}
	return strings.Join(words, " ")
}
