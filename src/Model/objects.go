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
