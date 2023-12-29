package controller

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	model "github.com/TealWater/clear-scribe/src/Model"
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("unable to laod environment variables")
	}
}

func sendPrompt(oldMessage string) (string, error) {
	dir, _ := os.Getwd()
	chat := &model.ChatGPT{}
	gptResponse := &model.GPTResponse{}
	chat.Model = "gpt-3.5-turbo"
	//Can't append values to an empty slice
	chat.Messages = make([]model.GPTSpec, 2)
	chat.Messages[0].Role = "system"
	chat.Messages[0].Content = "You are a helpful assistant who speaks multiple languages fluently and can simplify complex words given in a text."
	chat.Messages[1].Role = "user"
	chat.Messages[1].Content = fmt.Sprintf("Simplify this message for me %s", oldMessage)

	// Create a http client to build a request
	client := &http.Client{}
	url := os.Getenv("OPEN_AI_URL")
	gptJSONBytes, err := json.Marshal(chat)
	if err != nil {
		log.Println("unable to bind JSON from Chat GPT struct. \n Err: ", err)
		return "", errors.New("Error in " + dir + "\n On line 42")
	}

	// Build the request and add in chat GPT Struct/Object in JSON Form to the request
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(gptJSONBytes))
	if err != nil {
		log.Println("unable to create request. \n Err: ", err)
		log.Println("error at line 51")
		return "", errors.New("Error in " + dir + "\n On line 49")
	}

	req.Header = http.Header{
		"content-type":  {"application/json; charset=UTF-8"},
		"Authorization": {"Bearer " + os.Getenv("GPT_KEY")},
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Println("There is an error with processing your request. \n Err: ", err)
		return "", errors.New("Error in " + dir + "\n On line 61")
	}
	//close the response body after the func finished executing
	defer resp.Body.Close()

	//reading the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("Err:", err)
		return "", errors.New("Error in " + dir + "\n On line 70")
	}

	//place response JSON in a struct/object
	err = json.Unmarshal(body, gptResponse)
	if err != nil {
		log.Println("\n can't parse json. Err: ", err)
		return "", errors.New("Error in " + dir + "\n On line 77")
	}
	return gptResponse.Choices[0].Message.Content, nil
}
