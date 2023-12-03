package controller

import (
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	obj "github.com/TealWater/clear-scribe/src/Model"

	"github.com/gin-gonic/gin"
)

var oldMessage string
var mp map[string]string = make(map[string]string)

func init() {
	//mp = make(map[string]string)
	mp["i"] = "hi"
	mp["house"] = "dwelling"

}

func UploadText(c *gin.Context) {
	msg := &obj.IncomingText{}
	var hi = "hi mom"
	fmt.Println(hi)
	if err := c.ShouldBindJSON(msg); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fmt.Printf("%+v\n", *&msg.Message)
	oldMessage = *&msg.Message

	//fmt.Println(msg.Message)
	words := strings.Split(msg.Message, " ")
	//fmt.Println(words[0])

	for _, v := range words {
		fmt.Println(v)
		//search map for synonym
		//words[k] = map[v]
	}

	// c.JSON(http.StatusOK, gin.H{"message": "POST request recieved"})

	newMessage := parse(oldMessage)
	result := obj.MockEditedEssay{
		ID:         0,
		CreatedAt:  time.Now().String(),
		MessageOld: oldMessage,
		MessageNew: newMessage,
	}

	c.JSON(http.StatusOK, result)

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
	// c.JSON(http.StatusOK, gin.H{"message": "File uploaded successfully"})
	newMessage := parse(oldMessage)
	result := obj.MockEditedEssay{
		ID:         0,
		CreatedAt:  time.Now().String(),
		MessageOld: oldMessage,
		MessageNew: newMessage,
	}
	c.JSON(http.StatusOK, result)
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
