package controller

import (
	"fmt"
	"net/http"
	"strings"

	obj "github.com/TealWater/clear-scribe/src/Model"

	"github.com/gin-gonic/gin"
)

func Parse(c *gin.Context) {
	msg := &obj.IncomingText{}
	var hi = "hi mom"
	fmt.Println(hi)
	if err := c.ShouldBindJSON(msg); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fmt.Printf("%+v\n", *msg)

	//fmt.Println(msg.Message)
	words := strings.Split(msg.Message, " ")
	//fmt.Println(words[0])

	for _, v := range words {
		fmt.Println(v)
		//search map for synonym
		//words[k] = map[v]
	}

	c.JSON(http.StatusOK, gin.H{"message": "POST request recieved"})

}

// func Upload(c *gin.Context) {
// 	//save the income file to an actual file
// 	file, err := c.FormFile("file")
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	src, err := file.Open()
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}

// 	content, err := io.ReadAll(src)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}
// 	c.JSON(http.StatusOK, gin.H{"message": "File uploaded successfully"})
// }

/*

2 methods  --> one focus on open the file and getting the text data
			- focus on getting the data out of the response body

method #3 --> parse the text and add the changes


*/
