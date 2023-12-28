package model

import (
	"mime/multipart"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type IncomingText struct {
	Message string `json:"message"`
}

type FileUpload struct {
	File *multipart.FileHeader `form:"file" binding:"required"`
}

type EditedEssay struct {
	ID         primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	CreatedAt  primitive.DateTime `json:"createdAt,omitempty" bson:"createdAt,omitempty"`
	DateString string             `json:"dateString,omitempty" bson:"dateString,omitempty"`
	MessageOld string             `json:"messageOld,omitempty" bson:"messageOld,omitempty"`
	MessageNew string             `json:"messageNew,omitempty" bson:"messageNew,omitempty"`
}

type ChatGPT struct {
	Model    string    `json:"model"`
	Messages []GPTSpec `json:"messages"`
}

type GPTSpec struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type GPTResponse struct {
	Choices []struct {
		FinishReason string  `json:"finish_reason"`
		Index        int     `json:"index"`
		Message      GPTSpec `json:"message"`
		Logprobs     any     `json:"logprobs"`
	} `json:"choices"`
	Created int    `json:"created"`
	ID      string `json:"id"`
	Model   string `json:"model"`
	Object  string `json:"object"`
	Usage   struct {
		CompletionTokens int `json:"completion_tokens"`
		PromptTokens     int `json:"prompt_tokens"`
		TotalTokens      int `json:"total_tokens"`
	} `json:"usage"`
}
