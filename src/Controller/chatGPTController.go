package controller

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"

	model "github.com/TealWater/clear-scribe/src/Model"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("unable to laod environment variables")
	}
}

/*
TODO:
  - figure out it you have to build the model first then send the prompt or
    you can build the model and then the prompt all in one request.
*/
func sendPrompt(c *gin.Context) {

	chat := &model.ChatGPT{}
	gptResponse := &model.GPTResponse{}
	chat.Model = "gpt-3.5-turbo"
	chat.Messages[0].Role = "assistant"
	chat.Messages[0].Content = "You are a helpful assistant who speaks multiple languages fluently and can simplify complex words given in a text."

	// Create a http client to build a request
	client := &http.Client{}
	url := "https://api.openai.com/v1/chat/completions"
	gptJSONBytes, err := json.Marshal(chat)
	if err != nil {
		log.Println("unable to bind JSON from Chat GPT struct")
	}

	// Build the request and add in chat GPT Struct/Object in JSON Form to the request
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(gptJSONBytes))
	if err != nil {
		log.Println("unable to create request")
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	req.Header = http.Header{
		"content-type":   {"application/json; charset=UTF-8"},
		"Authorization:": {"Bearer " + os.Getenv("GPT_KEY")},
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Println("There is an error with processing your request")
		log.Println("err: ", err)
	}
	//close the response body after the func finished executing
	defer resp.Body.Close()

	//reading the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	//place response JSON in a struct/object
	err = json.Unmarshal(body, gptResponse)
	if err != nil {
		log.Println()
		log.Println("can't parse json")
	}

	c.JSON(http.StatusOK, gptResponse)

}
